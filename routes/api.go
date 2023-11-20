package routes

import (
	"github.com/gin-gonic/gin"
	"todo_list/app/controllers/app"
	"todo_list/app/middleware"
	"todo_list/app/services"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/register", app.Register)
	router.POST("/auth/login", app.Login)
	//	使用JWTAuth中间件，客户端需要使用正确的 Token 才能访问在authRouter 分组下的路由
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
	}
}
