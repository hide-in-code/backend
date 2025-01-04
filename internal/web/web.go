package web

import (
	"backend/internal/config"
	"backend/internal/convert"
	"backend/middleware"
	"backend/routers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitWeb(config *config.Config) {
	gin.SetMode(gin.DebugMode) //调试模式
	app := gin.New()
	app.NoRoute(middleware.NoRouteHandler())
	// 崩溃恢复
	app.Use(middleware.RecoveryMiddleware())
	app.LoadHTMLGlob(config.Web.StaticPath + "dist/*.html")
	app.Static("/static", config.Web.StaticPath+"dist/static")
	app.Static("/resource", config.Web.StaticPath+"resource")
	app.StaticFile("/favicon.ico", config.Web.StaticPath+"dist/favicon.ico")
	// 注册路由
	routers.RegisterRouter(app)
	go initHTTPServer(config, app)
}

// InitHTTPServer 初始化http服务
func initHTTPServer(config *config.Config, handler http.Handler) {
	srv := &http.Server{
		Addr:         ":" + convert.ToString(config.Web.Port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
