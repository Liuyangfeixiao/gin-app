package bootstrap

import (
	"context"
	"gin-demo/global"
	"gin-demo/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// 前端静态资源
	r.StaticFile("/", "./static/dist/index.html")
	r.Static("/assets", "./static/dist/assets")
	r.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
	// 其他静态资源
	r.Static("/public", "./static")
	r.Static("/storage", "./storage/app/public")

	// 注册 api 分组路由
	apiGroup := r.Group("/api")
	routes.SetApiGroupRoutes(apiGroup)

	return r
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()

	srv := &http.Server{
		Addr:    ":" + global.App.Config.App.Port,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待信号终端以优雅的关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞直到受到信号
	<-quit
	log.Println("Shutdown Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server Exiting")
}
