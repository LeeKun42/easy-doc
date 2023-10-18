package http

import (
	"context"
	"easy-doc/app/http/controller"
	"easy-doc/app/http/middleware"
	"easy-doc/app/lib/log"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/cors"
	recover2 "github.com/kataras/iris/v12/middleware/recover"
	"github.com/spf13/viper"
	"time"
)

const ServiceName = "user"

// StartWebServer 开启web服务
func StartWebServer() {
	app := iris.New()
	app.Use(recover2.New())
	//跨域中间件
	app.UseRouter(cors.New().Handler())
	app.Use(log.HttpLogHandler)

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		fmt.Println("Shutdown server")
		// 关闭所有主机
		app.Shutdown(ctx)
	})

	/** restful风格路由 */
	app.PartyFunc("/api", func(api iris.Party) {
		userController := controller.NewUserController()
		api.PartyFunc("/users", func(user iris.Party) {
			user.Post("", userController.Register)
			user.Get("/token", userController.Login)
			user.Get("/new-token", userController.RefreshToken).Use(middleware.JwtAuthCheck)
			user.Delete("/token", userController.Logout).Use(middleware.JwtAuthCheck)
			user.Get("/info", userController.Info).Use(middleware.JwtAuthCheck)
			user.Patch("/current", userController.Update).Use(middleware.JwtAuthCheck)
			user.Patch("/password", userController.UpdatePassword).Use(middleware.JwtAuthCheck)
		})
		projectController := controller.NewProjectController()
		api.PartyFunc("/projects", func(project iris.Party) {
			project.Use(middleware.JwtAuthCheck)
			project.Get("", projectController.GetProjects)
			project.Post("", projectController.Create)
			project.Get("/{project_id:int}", projectController.Get)
			project.Patch("/{project_id:int}", projectController.Update)
			project.Delete("/{project_id:int}", projectController.Delete)

			project.Get("/{project_id:int}/users", projectController.ListProjectUser)
			project.Post("/{project_id:int}/users", projectController.AddProjectUser)
			project.Delete("/{project_id:int}/users/{user_id:int}", projectController.DeleteProjectUser)

			project.Get("/{project_id:int}/directories", projectController.ListDirectory)
			project.Post("/{project_id:int}/directories", projectController.CreateDirectory)
			project.Get("/{project_id:int}/directories/{directory_id:int}", projectController.GetDirectory)
			project.Patch("/{project_id:int}/directories/{directory_id:int}", projectController.UpdateDirectory)
			project.Delete("/{project_id:int}/directories/{directory_id:int}", projectController.DeleteDirectory)

			project.Get("/{project_id:int}/apis", projectController.ListApi)
			project.Post("/{project_id:int}/apis", projectController.CreateApi)
			project.Put("/{project_id:int}/apis", projectController.ImportApis)
			project.Patch("/{project_id:int}/apis/{api_id:int}", projectController.UpdateApi)
			project.Post("/{project_id:int}/apis/{api_id:int}", projectController.CopyApi)
			project.Delete("/{project_id:int}/apis/{api_id:int}", projectController.DeleteApi)
			project.Get("/{project_id:int}/apis/{api_id:int}", projectController.GetApi)
		})
	})

	port := viper.GetInt("server.http")

	/**
	开启web服务
	参数1：监听地址和端口
	参数2：允许body多次消费
	*/
	app.Run(iris.Addr(fmt.Sprintf(":%d", port)), iris.WithoutBodyConsumptionOnUnmarshal)
}
