package project

import (
	_const "api-doc/app/lib/const"
	"api-doc/app/model"
	"api-doc/app/model/dto"
	"api-doc/app/model/request"
	"api-doc/app/model/resp"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	DB *gorm.DB
}

func NewService() *Service {
	return &Service{
		DB: model.Instance(),
	}
}

func (service *Service) GetList(loginUserId int) ([]dto.Project, error) {
	var projects []dto.Project
	var projectIds []int
	service.DB.Table(dto.TableNameProjectUser).Where("user_id", loginUserId).Pluck("project_id", &projectIds)
	if len(projectIds) <= 0 {
		return projects, nil
	}
	service.DB.Table(dto.TableNameProject).Where("id IN ?", projectIds).Find(&projects)

	return projects, nil
}

func (service *Service) Get(projectId int) resp.ProjectInfo {
	var project dto.Project
	var user dto.User
	service.DB.Table(project.TableName()).Where("id", projectId).First(&project)
	service.DB.Table(user.TableName()).Where("id", project.OwnerUserID).First(&user)
	var resp = resp.ProjectInfo{
		ID:        project.ID,
		Name:      project.Name,
		OwnerUser: user.NickName,
		CreatedAt: project.CreatedAt.Format(_const.DateTimeFormat),
	}
	return resp
}

func (service *Service) Create(params request.ProjectParams, loginUserId int) (int, error) {
	var projectDto dto.Project
	service.DB.Table(projectDto.TableName()).Where("name", params.Name).Where("owner_user_id", loginUserId).First(&projectDto)
	if projectDto.ID != 0 {
		return 0, errors.New("已创建过同名的项目，请修改名称后重试")
	}
	projectDto = dto.Project{
		Name:        params.Name,
		OwnerUserID: loginUserId,
	}
	tx := service.DB.Begin()
	result := tx.Create(&projectDto)
	if result.Error != nil {
		tx.Rollback()
		return 0, errors.New("创建失败")
	}
	pu := dto.ProjectUser{
		ProjectID: projectDto.ID,
		UserID:    projectDto.OwnerUserID,
	}
	tx.Create(&pu)
	tx.Commit()
	return projectDto.ID, nil
}

func (service *Service) Update(params request.ProjectParams, loginUserId int) error {
	var projectDto dto.Project
	service.DB.Table(projectDto.TableName()).Where("id", params.ID).First(&projectDto)
	if projectDto.ID == 0 {
		return errors.New("项目不存在")
	}
	if projectDto.OwnerUserID != loginUserId {
		return errors.New("你没有权限修改")
	}
	var count int64
	service.DB.Table(projectDto.TableName()).Where("name", params.Name).Where("owner_user_id", loginUserId).Where("id", "!=", params.ID).Count(&count)
	if count > 0 {
		return errors.New("已创建过同名的项目，请修改名称后重试")
	}
	result := service.DB.Table(projectDto.TableName()).Where("id", params.ID).Update("name", params.Name)
	if result.Error != nil {
		return errors.New("修改失败")
	}
	return nil
}

func (service *Service) Delete(projectId int, loginUserId int) error {
	var projectDto dto.Project
	service.DB.Table(projectDto.TableName()).Where("id", projectId).First(&projectDto)
	if projectDto.ID == 0 {
		return errors.New("项目不存在")
	}
	if projectDto.OwnerUserID != loginUserId {
		return errors.New("你没有权限删除")
	}
	tx := service.DB.Begin()
	//删除项目
	tx.Delete(&projectDto)
	//删除项目用户
	tx.Where("project_id = ?", projectId).Delete(&dto.ProjectUser{})
	//删除项目目录
	tx.Where("project_id = ?", projectId).Delete(&dto.ProjectDirectory{})
	//删除项目接口
	tx.Where("project_id = ?", projectId).Delete(&dto.ProjectAPI{})
	tx.Commit()
	return nil
}

func (service *Service) CheckProjectPermissions(projectId int, userId int, action string) bool {
	var projectDto dto.Project
	service.DB.Table(projectDto.TableName()).Where("id", projectId).First(&projectDto)
	if projectDto.ID == 0 {
		return false
	}

	switch action {
	case "add-project-user":
		if projectDto.OwnerUserID != userId {
			return false
		} else {
			return true
		}
		break
	case "delete-project-user":
		if projectDto.OwnerUserID != userId {
			return false
		} else {
			return true
		}
		break
	}

	return false
}

func (service *Service) ListProjectUser(projectId int) []dto.User {
	var userIds []int
	var users []dto.User
	service.DB.Table(dto.TableNameProjectUser).Where("project_id", projectId).Pluck("user_id", &userIds)
	service.DB.Table(dto.TableNameUser).Where("id IN ?", userIds).Find(&users)
	return users
}

func (service *Service) AddProjectUser(params request.ProjectUserParams, loginUserId int) error {
	if !service.CheckProjectPermissions(params.ProjectID, loginUserId, "add-project-user") {
		return errors.New("你没有权限操作")
	}
	var user dto.User
	result := service.DB.Table(user.TableName()).Where("account", params.Account).First(&user)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("输入的账号不存在")
	}
	var count int64
	service.DB.Table(dto.TableNameProjectUser).Where("project_id", params.ProjectID).Where("user_id", user.ID).Count(&count)
	if count > 0 {
		return errors.New("账号已添加")
	}
	service.DB.Table(dto.TableNameProjectUser).Create(dto.ProjectUser{ProjectID: params.ProjectID, UserID: user.ID})
	return nil
}

func (service *Service) DeleteProjectUser(projectId int, userId int, loginUserId int) error {
	if !service.CheckProjectPermissions(projectId, loginUserId, "delete-project-user") {
		return errors.New("你没有权限操作")
	}
	service.DB.Where("project_id", projectId).Where("user_id", userId).Delete(&dto.ProjectUser{})
	return nil
}

func (service *Service) GetDirectory(directoryId int, userId int) (dto.ProjectDirectory, error) {
	var directory dto.ProjectDirectory
	//获取所有目录
	result := service.DB.Table(dto.TableNameProjectDirectory).Where("id", directoryId).First(&directory)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf(result.Error.Error())
	}
	return directory, nil
}

func (service *Service) GetDirectories(projectId int, userId int) ([]resp.Directory, error) {
	var directories []dto.ProjectDirectory
	//获取所有目录
	result := service.DB.Table(dto.TableNameProjectDirectory).Where("project_id", projectId).Select("id", "name", "parent_id", "seq").Order("seq").Find(&directories)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf(result.Error.Error())
	}
	//处理目录层级
	var rootDir = resp.Directory{ID: 0}
	treeDirs(&rootDir, directories)
	return rootDir.Children, nil
}

func treeDirs(currentDir *resp.Directory, directories []dto.ProjectDirectory) {
	for _, directory := range directories {
		if directory.ParentID == currentDir.ID {
			dir := resp.Directory{ID: directory.ID, Name: directory.Name, Children: []resp.Directory{}, Apis: []resp.Api{}}
			treeApis(&dir, directories)
			currentDir.Children = append(currentDir.Children, dir)
		}
	}
}

func (service *Service) CreateDirectory(params request.ProjectDirectoryParams, userId int) error {
	var pdDto dto.ProjectDirectory
	service.DB.Table(pdDto.TableName()).Where("name", params.Name).Where("project_id", params.ProjectID).First(&pdDto)
	if pdDto.ID != 0 {
		return errors.New("已创建过同名的目录，请修改名称后重试")
	}
	pdDto = dto.ProjectDirectory{
		ProjectID: params.ProjectID,
		ParentID:  params.ParentID,
		Name:      params.Name,
		Seq:       params.Seq,
		Desc:      params.Desc,
	}
	result := service.DB.Create(&pdDto)
	if result.Error != nil {
		return errors.New("创建失败")
	}
	return nil
}

func (service *Service) UpdateDirectory(params request.ProjectDirectoryParams, userId int) error {
	var pdDto dto.ProjectDirectory
	var count int64
	service.DB.Table(pdDto.TableName()).Where("name", params.Name).Where("project_id", params.ProjectID).Where("id", "!=", params.ID).Count(&count)
	if count > 0 {
		return errors.New("已创建过同名的目录，请修改名称后重试")
	}
	result := service.DB.Table(pdDto.TableName()).Where("id", params.ID).
		Select("Name", "Desc", "ParentID", "Seq").
		Updates(dto.ProjectDirectory{ParentID: params.ParentID, Name: params.Name, Desc: params.Desc, Seq: params.Seq})
	if result.Error != nil {
		return errors.New("修改失败")
	}
	return nil
}

// DeleteDirectory 删除目录需要递归删除子目录
func (service *Service) DeleteDirectory(projectId int, directoryId int, userId int) error {
	var pdDto dto.ProjectDirectory
	var ids []int
	service.DB.Table(pdDto.TableName()).Where("id", directoryId).First(&pdDto)
	if pdDto.ID == 0 {
		return errors.New("目录不存在")
	}
	if pdDto.ParentID == 0 {
		service.DB.Table(pdDto.TableName()).Where("parent_id", directoryId).Pluck("id", &ids)
	}

	ids = append(ids, directoryId)
	tx := service.DB.Begin()
	tx.Where("id IN ?", ids).Delete(&dto.ProjectDirectory{})
	tx.Where("directory_id IN ?", ids).Delete(&dto.ProjectAPI{})
	tx.Commit()
	return nil
}

func (service *Service) GetApis(projectId int, userId int) ([]resp.Directory, error) {
	var directories []dto.ProjectDirectory
	//获取所有目录
	result := service.DB.Table(dto.TableNameProjectDirectory).Where("project_id", projectId).Select("id", "name", "parent_id", "seq").Preload("Apis", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "directory_id", "name", "seq").Order("project_apis.seq")
	}).Order("seq").Find(&directories)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf(result.Error.Error())
	}
	//处理目录层级
	var rootDir = resp.Directory{ID: 0}
	treeApis(&rootDir, directories)
	return rootDir.Children, nil
}

func treeApis(currentDir *resp.Directory, directories []dto.ProjectDirectory) {
	for _, directory := range directories {
		if directory.ParentID == currentDir.ID {
			dir := resp.Directory{ID: directory.ID, Name: directory.Name, Children: []resp.Directory{}, Apis: []resp.Api{}}
			for _, api := range directory.Apis {
				dir.Apis = append(dir.Apis, resp.Api{ID: api.ID, Name: api.Name})
			}
			treeApis(&dir, directories)
			currentDir.Children = append(currentDir.Children, dir)
		}
	}
}

func (service *Service) CreateApi(params request.ProjectApiParams, userId int) (int, error) {
	var apiDto dto.ProjectAPI
	service.DB.Table(apiDto.TableName()).Where("name", params.Name).Where("project_id", params.ProjectID).First(&apiDto)
	if apiDto.ID != 0 {
		return 0, errors.New("已创建过同名的接口，请修改名称后重试")
	}
	apiDto = dto.ProjectAPI{
		ProjectID:      params.ProjectID,
		DirectoryID:    params.DirectoryID,
		Name:           params.Name,
		Path:           params.Path,
		Method:         params.Method,
		RequestHeaders: params.RequestHeaders,
		RequestQuery:   params.RequestQuery,
		RequestPath:    params.RequestPath,
		RequestBody:    params.RequestBody,
		ResponseBody:   params.ResponseBody,
		Desc:           params.Desc,
	}
	result := service.DB.Create(&apiDto)
	if result.Error != nil {
		return 0, errors.New("创建失败")
	}
	return apiDto.ID, nil
}

func (service *Service) UpdateApi(params request.ProjectApiParams, userId int) error {
	var apiDto dto.ProjectAPI
	var count int64
	service.DB.Table(apiDto.TableName()).Where("name", params.Name).Where("project_id", params.ProjectID).Where("id", "!=", params.ID).Count(&count)
	if count > 0 {
		return errors.New("已创建过同名的接口，请修改名称后重试")
	}
	apiDto.Name = params.Name
	apiDto.DirectoryID = params.DirectoryID
	apiDto.Name = params.Name
	apiDto.Path = params.Path
	apiDto.Method = params.Method
	apiDto.RequestHeaders = params.RequestHeaders
	apiDto.RequestQuery = params.RequestQuery
	apiDto.RequestPath = params.RequestPath
	apiDto.RequestBody = params.RequestBody
	apiDto.ResponseBody = params.ResponseBody
	apiDto.Desc = params.Desc
	result := service.DB.Table(apiDto.TableName()).Where("id", params.ID).Updates(apiDto)
	if result.Error != nil {
		return errors.New("修改失败")
	}
	return nil
}

func (service *Service) CopyApi(apiId int, userId int) (int, error) {
	var oldDto dto.ProjectAPI
	service.DB.Table(oldDto.TableName()).Where("id", apiId).First(&oldDto)
	if oldDto.ID == 0 {
		return 0, errors.New("接口不存在，无法复制")
	}
	oldDto.ID = 0
	oldDto.Name = oldDto.Name + "_copy"
	oldDto.Seq = oldDto.Seq + 1
	oldDto.CreatedAt = time.Now()
	result := service.DB.Create(&oldDto)
	if result.Error != nil {
		return 0, errors.New("创建失败")
	}
	return oldDto.ID, nil
}

func (service *Service) DeleteApi(projectId int, apiId int, userId int) error {
	var pdDto dto.ProjectAPI
	service.DB.Table(pdDto.TableName()).Where("id", apiId).First(&pdDto)
	if pdDto.ID == 0 {
		return errors.New("接口不存在")
	}
	service.DB.Delete(&pdDto)
	return nil
}

func (service *Service) GetApi(id int) (dto.ProjectAPI, error) {
	var apiDto dto.ProjectAPI
	service.DB.Table(apiDto.TableName()).Where("id", id).First(&apiDto)
	if apiDto.ID == 0 {
		return apiDto, errors.New("接口不存在")
	}
	return apiDto, nil
}
