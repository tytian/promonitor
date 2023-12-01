package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func PromProxy(handler http.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func main() {
	engine := gin.Default()
	engine.Any("/metrics", PromProxy(promhttp.Handler()))
	engine.Run(":8081")
}
