package main

import (
	"gin-demo/bootstrap"
	"gin-demo/global"
)

func main() {
	// 初始化配置
	bootstrap.InitConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitLog()
	global.App.Log.Info("log init success!")

	// 初始化数据库
	global.App.DB = bootstrap.InitDB()
	// 程序关闭前释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 启动服务器
	bootstrap.RunServer()
}
