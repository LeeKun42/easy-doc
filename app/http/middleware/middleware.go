package middleware

import (
	"easy-doc/app/service/jwt"
	"errors"
	"github.com/kataras/iris/v12"
	"strings"
)

func JwtAuthCheck(ctx iris.Context) {
	authorization := ctx.GetHeader("Authorization")
	if authorization == "" {
		ctx.StopWithError(401, errors.New("token无效"))
	}
	authArr := strings.Split(authorization, " ")
	if len(authArr) != 2 {
		ctx.StopWithError(401, errors.New("token无效"))
	}
	tokenString := authArr[1]
	claims, err := jwt.NewService().Check(tokenString)
	if err != nil {
		ctx.StopWithError(401, err)
	}
	ctx.Values().Set("user_id", claims.UserId)
	ctx.Values().Set("jwt_token", tokenString)
	ctx.Next()
}
