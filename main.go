package main

import (
	"gin-demo/bootstrap"
	"gin-demo/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 初始化配置
	bootstrap.InitConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitLog()
	global.App.Log.Info("log init success!")

	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 启动服务器
	r.Run(":" + global.App.Config.App.Port)
}
