package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIp := ctx.ClientIP()

		log.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode, latencyTime, clientIp, reqMethod, reqUri,
		)
	}
}
