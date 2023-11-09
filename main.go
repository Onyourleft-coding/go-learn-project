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
	//初始化校验器
	bootstrap.InitializeValidator()

	//启动服务器
	bootstrap.RunServer()

}
