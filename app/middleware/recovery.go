package middleware

import (
	"gin-demo/app/common/response"
	"gin-demo/global"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func CustomRecovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(&lumberjack.Logger{
		Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Log.Filename,
		MaxSize:    global.App.Config.Log.MaxSize,
		MaxAge:     global.App.Config.Log.MaxAge,
		MaxBackups: global.App.Config.Log.MaxBackups,
		LocalTime:  false,
		Compress:   global.App.Config.Log.Compress,
	}, response.ServerError)
}
