package bootstrap

import (
	"github.com/gin-gonic/gin"
	"todo_list/global"
	"todo_list/routes"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// 注册 api 分组路由
	apiGroup := router.Group("/api")
	routes.SetApiGroupRoutes(apiGroup)

	return router
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()
	r.Run(":" + global.App.Config.App.Port)
}
