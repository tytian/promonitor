package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tytian/gin_metric"
	"net/http"
	"promonitor/middleware"
)

func StartMonitor(addr string) {
	engine := gin.Default()

	engine.Use(middleware.Logger(), gin.Recovery(), middleware.Waf())

	gin_metric.RegisterMetricMonitor(true, engine)
	gin_metric.RegisterApiStatistics(true, engine)

	engine.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	engine.POST("/listuser", UserList)
	engine.POST("/getuser", UserGet)
	engine.POST("/createuser", UserCreate)

	engine.Run(addr)
}
