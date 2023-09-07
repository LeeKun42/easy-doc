package controller

import (
	"api-doc/app/http/response"
	"api-doc/app/model/request"
	"api-doc/app/service/project"
	"github.com/kataras/iris/v12"
)

type ProjectController struct {
	ProjectService *project.Service
}

func NewProjectController() *ProjectController {
	return &ProjectController{
		ProjectService: project.NewService(),
	}
}

func (pc *ProjectController) GetProjects(context iris.Context) {
	loginUserId, _ := context.Values().GetInt("user_id")
	projects, err := pc.ProjectService.GetList(loginUserId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, projects)
	}
}

func (pc *ProjectController) Get(context iris.Context) {
	projectId, _ := context.Params().GetInt("project_id")
	project := pc.ProjectService.Get(projectId)
	response.Success(context, iris.Map{"project": project})
}

func (pc *ProjectController) Create(context iris.Context) {
	var inputParams request.ProjectParams
	context.ReadJSON(&inputParams)
	loginUserId, _ := context.Values().GetInt("user_id")
	_, err := pc.ProjectService.Create(inputParams, loginUserId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{})
	}
}

func (pc *ProjectController) Update(context iris.Context) {
	var inputParams request.ProjectParams
	context.ReadJSON(&inputParams)
	projectId, _ := context.Params().GetInt("project_id")
	inputParams.ID = projectId
	loginUserId, _ := context.Values().GetInt("user_id")
	err := pc.ProjectService.Update(inputParams, loginUserId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{})
	}
}

func (pc *ProjectController) Delete(context iris.Context) {
	projectId, _ := context.Params().GetInt("project_id")
	loginUserId, _ := context.Values().GetInt("user_id")
	err := pc.ProjectService.Delete(projectId, loginUserId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{})
	}
}

func (pc *ProjectController) ListProjectUser(context iris.Context) {
	projectId, _ := context.Params().GetInt("project_id")
	users := pc.ProjectService.ListProjectUser(projectId)
	response.Success(context, iris.Map{"users": users})
}

func (pc *ProjectController) AddProjectUser(context iris.Context) {
	var inputParams request.ProjectUserParams
	context.ReadJSON(&inputParams)
	projectId, _ := context.Params().GetInt("project_id")
	loginUserId, _ := context.Values().GetInt("user_id")
	inputParams.ProjectID = projectId
	err := pc.ProjectService.AddProjectUser(inputParams, loginUserId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{})
	}
}

func (pc *ProjectController) DeleteProjectUser(context iris.Context) {
	projectId, _ := context.Params().GetInt("project_id")
	userId, _ := context.Params().GetInt("user_id")
	loginUserId, _ := context.Values().GetInt("user_id")
	if userId == loginUserId {
		response.Fail(context, "不能删除自己")
		return
	}
	err := pc.ProjectService.DeleteProjectUser(projectId, userId, loginUserId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{})
	}
}

func (pc *ProjectController) ListDirectory(context iris.Context) {
	var params request.ProjectApiParams
	context.ReadJSON(&params)
	projectId, _ := context.Params().GetInt("project_id")
	params.ProjectID = projectId
	loginUserId, _ := context.Values().GetInt("user_id")
	directories, err := pc.ProjectService.GetDirectories(projectId, loginUserId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{"directories": directories})
	}
}

func (pc *ProjectController) GetDirectory(context iris.Context) {
	var params request.ProjectApiParams
	context.ReadJSON(&params)
	directoryId, _ := context.Params().GetInt("directory_id")
	loginUserId, _ := context.Values().GetInt("user_id")
	directory, err := pc.ProjectService.GetDirectory(directoryId, loginUserId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, directory)
	}
}

func (pc *ProjectController) CreateDirectory(context iris.Context) {
	var params request.ProjectDirectoryParams
	context.ReadJSON(&params)
	projectId, _ := context.Params().GetInt("project_id")
	params.ProjectID = projectId
	loginUserId, _ := context.Values().GetInt("user_id")
	err := pc.ProjectService.CreateDirectory(params, loginUserId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{})
	}
}

func (pc *ProjectController) UpdateDirectory(context iris.Context) {
	var params request.ProjectDirectoryParams
	context.ReadJSON(&params)
	projectId, _ := context.Params().GetInt("project_id")
	directoryId, _ := context.Params().GetInt("directory_id")
	params.ID = directoryId
	params.ProjectID = projectId
	loginUserId, _ := context.Values().GetInt("user_id")
	err := pc.ProjectService.UpdateDirectory(params, loginUserId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{})
	}
}

func (pc *ProjectController) DeleteDirectory(context iris.Context) {
	projectId, _ := context.Params().GetInt("project_id")
	directoryId, _ := context.Params().GetInt("directory_id")
	loginUserId, _ := context.Values().GetInt("user_id")
	err := pc.ProjectService.DeleteDirectory(projectId, directoryId, loginUserId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{})
	}
}

func (pc *ProjectController) ListApi(context iris.Context) {
	var params request.ProjectApiParams
	context.ReadJSON(&params)
	projectId, _ := context.Params().GetInt("project_id")
	params.ProjectID = projectId
	loginUserId, _ := context.Values().GetInt("user_id")
	directories, err := pc.ProjectService.GetApis(projectId, loginUserId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{"directories": directories})
	}
}

func (pc *ProjectController) CreateApi(context iris.Context) {
	var params request.ProjectApiParams
	context.ReadJSON(&params)
	projectId, _ := context.Params().GetInt("project_id")
	params.ProjectID = projectId
	loginUserId, _ := context.Values().GetInt("user_id")
	id, err := pc.ProjectService.CreateApi(params, loginUserId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{"id": id})
	}
}

func (pc *ProjectController) UpdateApi(context iris.Context) {
	var params request.ProjectApiParams
	context.ReadJSON(&params)
	projectId, _ := context.Params().GetInt("project_id")
	apiId, _ := context.Params().GetInt("api_id")
	params.ID = apiId
	params.ProjectID = projectId
	loginUserId, _ := context.Values().GetInt("user_id")
	err := pc.ProjectService.UpdateApi(params, loginUserId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{})
	}
}

func (pc *ProjectController) CopyApi(context iris.Context) {
	apiId, _ := context.Params().GetInt("api_id")
	loginUserId, _ := context.Values().GetInt("user_id")
	id, err := pc.ProjectService.CopyApi(apiId, loginUserId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{"id": id})
	}
}

func (pc *ProjectController) DeleteApi(context iris.Context) {
	projectId, _ := context.Params().GetInt("project_id")
	apiId, _ := context.Params().GetInt("api_id")
	loginUserId, _ := context.Values().GetInt("user_id")
	err := pc.ProjectService.DeleteApi(projectId, apiId, loginUserId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{})
	}
}

func (pc *ProjectController) GetApi(context iris.Context) {
	apiId, _ := context.Params().GetInt("api_id")
	api, err := pc.ProjectService.GetApi(apiId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, api)
	}
}
