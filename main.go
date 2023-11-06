package main

import (
	"todo_list/bootstrap"
	"todo_list/global"
)

func main() {

	//初始化配置
	bootstrap.InitializeConfig()

	//初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log int success!")

	//初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	//程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	bootstrap.RunServer()
	//r := gin.Default()
	////	测试路由
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"data":    "pong",
	//		"code":    0,
	//		"message": "请求成功",
	//	})
	//})
	//r.Run(":" + global.App.Config.App.Port)
}
