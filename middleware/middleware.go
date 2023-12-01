package middleware

import (
	"github.com/gin-gonic/gin"
	"promonitor/monitor"
	"strconv"
	"time"
)

func MetricMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		now := time.Now()
		method := ctx.Request.Method
		path := ctx.Request.URL.Path
		monitor.MetricMonitor.RecvRequestTotal(method, path)
		ctx.Next()
		status := strconv.Itoa(ctx.Writer.Status())
		monitor.MetricMonitor.RecvServerHandlerSeconds(method, path, status, time.Now().Sub(now).Seconds())
		monitor.MetricMonitor.RecvServerHandlerStatus(method, path, status)
	}
}
