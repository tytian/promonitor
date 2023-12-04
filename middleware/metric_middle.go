package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"promonitor/monitor"
	"strconv"
	"time"
)

func MetricMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		now := time.Now()
		method := ctx.Request.Method
		path := ctx.Request.URL.Path
		monitor.MetricMonitor.RecvRequestTotal(TypeHTTP, method, path)
		// ioutil.ReadAll 读取到的是字节流[]byte，读完body就没有了
		bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			log.WithFields(log.Fields{
				"path":    path,
				"errcode": "Get request body fail: " + err.Error(),
			})
		}
		// 读出来之后需要重新写回去，不然在接口处理函数中无法获取到数据
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		ctx.Next()
		status := strconv.Itoa(ctx.Writer.Status())
		monitor.MetricMonitor.RecvServerHandlerSeconds(TypeHTTP, method, path, status, time.Now().Sub(now).Seconds())
		monitor.MetricMonitor.RecvServerHandlerStatus(TypeHTTP, method, path, status)
		if ctx.Writer.Status() != http.StatusOK {
			log.WithFields(log.Fields{
				"request": string(bodyBytes),
				"code":    status,
			}).Warn("request fail ")
		}
	}
}
