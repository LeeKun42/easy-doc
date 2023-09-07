package response

import (
	"github.com/kataras/iris/v12"
)

// Success 正常响应
func Success(context iris.Context, v interface{}) {
	data := iris.Map{
		"code":    0,
		"message": "success",
		"data":    v,
	}
	context.StopWithJSON(200, data)
}

// Fail 业务错误响应
func Fail(context iris.Context, message string) {
	data := iris.Map{
		"code":    10000,
		"message": message,
		"data":    iris.Map{},
	}
	context.StopWithJSON(200, data)
}

// Error http异常响应
func Error(context iris.Context, code int, message string) {
	data := iris.Map{
		"code":    code,
		"message": message,
		"data":    iris.Map{},
	}
	context.StopWithJSON(code, data)
}
