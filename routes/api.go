package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/test", func(c *gin.Context) {
		time.Sleep(time.Second * 5)
		c.String(http.StatusOK, "success")
	})
}
