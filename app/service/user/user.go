package user

import (
	"easy-doc/app/lib/hash"
	"easy-doc/app/model"
	"easy-doc/app/model/dto"
	"easy-doc/app/model/request"
	"easy-doc/app/service/jwt"
	"errors"
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

const (
	// StatusDisabled 用户状态 禁用
	StatusDisabled int = 0
	// StatusEnable 用户状态 启用
	StatusEnable int = 1
)

func (service *Service) Register(input request.RegisterUser) (int, error) {
	var user dto.User
	service.DB.Table(user.TableName()).Where("account", input.Account).First(&user)
	if user.ID != 0 {
		return 0, errors.New("注册失败：账号已存在")
	}
	user = dto.User{
		Account:   input.Account,
		NickName:  input.NickName,
		Passwd:    hash.Make(input.Passwd),
		CreatedAt: time.Now(),
		Status:    StatusEnable,
	}
	result := service.DB.Create(&user)
	if result.Error != nil {
		return 0, errors.New("注册失败：账号已存在")
	}
	return user.ID, nil
}

func (service *Service) Login(input request.UserLogin) (string, error) {
	var user dto.User
	result := service.DB.Table(user.TableName()).Where("account", input.Account).First(&user)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", errors.New("账号不存在")
	}
	if user.Status == StatusDisabled {
		return "", errors.New("账号已被禁用，请联系管理员")
	}
	if !hash.Check(input.Passwd, user.Passwd) { //密码不正确
		return "", errors.New("密码不正确")
	}
	token := jwt.NewService().Create(user.ID, 0)
	return token, nil
}

func (service *Service) Disabled(userId int) error {
	var user dto.User
	tx := service.DB.Table(user.TableName()).Where("id", userId).Update("status", StatusDisabled)
	return tx.Error
}

func (service *Service) Enable(userId int) error {
	var user dto.User
	tx := service.DB.Table(user.TableName()).Where("id", userId).Update("status", StatusEnable)
	return tx.Error
}

func (service *Service) Update(uParams request.UserUpdate) error {
	var user dto.User
	tx := service.DB.Table(user.TableName()).Where("id", uParams.ID).Select("account", "nick_name").Updates(dto.User{Account: uParams.Account, NickName: uParams.NickName})
	return tx.Error
}

func (service *Service) UpdatePassword(uParams request.UserUpdatePassword) error {
	var user dto.User
	service.DB.Table(user.TableName()).Where("id", uParams.ID).First(&user)
	if user.ID == 0 {
		return errors.New("账号不存在")
	}
	if !hash.Check(uParams.OldPassword, user.Passwd) {
		return errors.New("旧密码输入错误")
	}
	tx := service.DB.Table(user.TableName()).Where("id", uParams.ID).Update("passwd", hash.Make(uParams.NewPassword))
	return tx.Error
}

func (service *Service) GetList(input request.UserSearch) ([]dto.User, error) {
	var users []dto.User
	offset := (input.PageIndex - 1) * input.PageSize
	result := service.DB.Offset(offset).Limit(input.PageSize).Find(&users)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return users, errors.New("用户不存在")
	}
	return users, nil
}

func (service *Service) Get(id int) (dto.User, error) {
	var user dto.User
	result := service.DB.Find(&user, id)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("用户不存在")
	}
	return user, nil
}
