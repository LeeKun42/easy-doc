package controller

import (
	"api-doc/app/http/response"
	_const "api-doc/app/lib/const"
	"api-doc/app/model/request"
	response2 "api-doc/app/model/resp"
	"api-doc/app/service/jwt"
	"api-doc/app/service/user"
	"github.com/kataras/iris/v12"
	"strings"
)

type UserController struct {
	UserService *user.Service
	JwtService  *jwt.Service
}

func NewUserController() *UserController {
	return &UserController{
		UserService: user.NewService(),
		JwtService:  jwt.NewService(),
	}
}

func (uc *UserController) Register(context iris.Context) {
	var input request.RegisterUser
	context.ReadJSON(&input)
	if strings.Count(input.Passwd, "") < 6 {
		response.Fail(context, "密码至少6位字符")
		return
	}
	id, err := uc.UserService.Register(input)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{"user_id": id})
	}
}

func (uc *UserController) Login(context iris.Context) {
	var input request.UserLogin
	context.ReadQuery(&input)
	token, err := uc.UserService.Login(input)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{"token": token})
	}
}

func (uc *UserController) RefreshToken(context iris.Context) {
	token, err := uc.JwtService.Refresh(context.Values().GetString("jwt_token"))
	if err != nil {
		response.Error(context, 401, err.Error())
	} else {
		response.Success(context, iris.Map{"token": token})
	}
}

func (uc *UserController) Logout(context iris.Context) {
	uc.JwtService.Invalidate(context.Values().GetString("jwt_token"))
	response.Success(context, iris.Map{})
}

func (uc *UserController) Update(context iris.Context) {
	var input request.UserUpdate
	context.ReadJSON(&input)
	loginUserId, _ := context.Values().GetInt("user_id")
	input.ID = loginUserId
	err := uc.UserService.Update(input)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{})
	}
}

func (uc *UserController) UpdatePassword(context iris.Context) {
	var input request.UserUpdatePassword
	context.ReadJSON(&input)
	loginUserId, _ := context.Values().GetInt("user_id")
	input.ID = loginUserId
	err := uc.UserService.UpdatePassword(input)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		response.Success(context, iris.Map{})
	}
}

func (uc *UserController) Info(context iris.Context) {
	loginUserId, _ := context.Values().GetInt("user_id")
	userInfo, err := uc.UserService.Get(loginUserId)
	if err != nil {
		response.Fail(context, err.Error())
	} else {

		result := response2.UserInfo{
			ID:        userInfo.ID,
			NickName:  userInfo.NickName,
			Account:   userInfo.Account,
			Status:    userInfo.Status,
			CreatedAt: userInfo.CreatedAt.Format(_const.DateTimeFormat),
		}

		response.Success(context, result)
	}
}

func (uc *UserController) GetList(context iris.Context) {
	var input request.UserSearch
	context.ReadQuery(&input)
	if input.PageIndex == 0 {
		input.PageIndex = 1
	}
	if input.PageSize == 0 {
		input.PageSize = 10
	}
	users, err := uc.UserService.GetList(input)
	if err != nil {
		response.Fail(context, err.Error())
	} else {
		var res []response2.UserInfo
		for _, userInfo := range users {
			result := response2.UserInfo{
				ID:       userInfo.ID,
				NickName: userInfo.NickName,
				Account:  userInfo.Account,
				Status:   userInfo.Status,
			}
			res = append(res, result)
		}

		response.Success(context, res)
	}
}
