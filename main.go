package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"net/http"
	"promonitor/middleware"
)

func PromProxy(handler http.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func main() {
	engine := gin.Default()
	engine.Use(middleware.MetricMiddleware())
	engine.Any("/metrics", PromProxy(promhttp.Handler()))
	engine.POST("/hello", func(ctx *gin.Context) {
		bodyBytes, _ := ioutil.ReadAll(ctx.Request.Body)
		defer ctx.Request.Body.Close()
		fmt.Printf("Hello %s", bodyBytes)
	})
	engine.Run(":8081")
}
