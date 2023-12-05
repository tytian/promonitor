package server

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"promonitor/middleware"
)

func PromProxy(handler http.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func StartMonitor(addr string) {
	engine := gin.Default()
	engine.Use(middleware.MetricMiddleware())
	engine.Any("/metrics", PromProxy(promhttp.Handler()))
	engine.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})
	engine.POST("/userlist", UserList)
	engine.POST("/getuser", UserGet)
	engine.Run(addr)
}
