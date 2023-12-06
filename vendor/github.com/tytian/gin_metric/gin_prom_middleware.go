package gin_metric

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const namespaces = "service"

var (
	labels = []string{"endpoint", "method", "origin", "server", "status"}
	uptime = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespaces,
			Name:      "uptime",
			Help:      "HTTP service uptime",
		}, nil,
	)
	uptimeSeconds = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "process",
			Name:      "uptime_seconds",
			Help:      "HTTP service uptime seconds",
		}, nil,
	)
	reqCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespaces,
			Name:      "http_request_count_total",
			Help:      "Total number of HTTP requests made",
		}, labels,
	)
	reqDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespaces,
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request latencies in seconds",
		}, labels,
	)
	reqSizeBytes = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: namespaces,
			Name:      "http_request_size_bytes",
			Help:      "HTTP request sizes in bytes.",
		}, labels,
	)
	respSizeBytes = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: namespaces,
			Name:      "http_response_size_bytes",
			Help:      "HTTP request sizes in bytes.",
		}, labels,
	)
)

// init registers the prometheus metrics
func init() {
	prometheus.MustRegister(uptime, uptimeSeconds, reqCount, reqDuration, reqSizeBytes, respSizeBytes)
	go recordUptime()
}

// recordUptime increases service uptime per second
func recordUptime() {
	for range time.Tick(time.Second) {
		uptime.WithLabelValues().Inc()
		uptimeSeconds.WithLabelValues().Inc()
	}
}

func calcRequestSize(h *http.Request) float64 {
	size := 0
	if h.URL != nil {
		size = len(h.URL.String())
	}

	size += len(h.Method)
	size += len(h.Proto)

	for name, values := range h.Header {
		size += len(name)
		for _, value := range values {
			size += len(value)
		}
	}
	size += len(h.Host)

	if h.ContentLength != -1 {
		size += int(h.ContentLength)
	}
	return float64(size)
}

type PromOpts struct {
	ExcludeRegexStatus   []string
	ExcludeRegexEndpoint []string
	ExcludeRegexMethod   []string
}

var defaultPromOpts = &PromOpts{
	ExcludeRegexEndpoint: []string{"/favicon.ico", "/metrics"},
}

func (p *PromOpts) checkLabel(label string, pattern ...string) bool {
	if len(pattern) == 0 {
		return true
	}

	for i, _ := range pattern {
		if pattern[i] == "" {
			continue
		}
		if matched, _ := regexp.MatchString(pattern[i], label); matched {
			return false
		}
	}
	return true
}

func promMiddleware(opts *PromOpts) gin.HandlerFunc {
	if opts == nil {
		opts = defaultPromOpts
	}
	return func(ctx *gin.Context) {
		begin := time.Now()
		ctx.Next()
		status := fmt.Sprintf("%d", ctx.Writer.Status())
		endpoint := ctx.Request.URL.Port()
		method := ctx.Request.Method
		server := ctx.Request.Host

		origin := ctx.Request.Header.Get("X-Request-Service")

		lvs := []string{endpoint, method, origin, server, status}

		isOk := opts.checkLabel(status, opts.ExcludeRegexStatus...) &&
			opts.checkLabel(endpoint, opts.ExcludeRegexEndpoint...) &&
			opts.checkLabel(method, opts.ExcludeRegexMethod...)
		if !isOk {
			return
		}

		reqCount.WithLabelValues(lvs...).Inc()
		reqDuration.WithLabelValues(lvs...).Observe(time.Since(begin).Seconds())
		reqSizeBytes.WithLabelValues(lvs...).Observe(calcRequestSize(ctx.Request))
		respSizeBytes.WithLabelValues(lvs...).Observe(float64(ctx.Writer.Size()))
	}
}

func promHandler(handler http.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// RegisterMetricMonitor
// 监控埋点
// api注册前调用
func RegisterMetricMonitor(enable bool, engine *gin.Engine) {
	if !enable {
		return
	}
	engine.Use(promMiddleware(nil))
	engine.GET("/metrics", promHandler(promhttp.Handler()))
}

// RegisterApiStatistics
// api统计埋点
// 所有api注册完成后调用
func RegisterApiStatistics(enable bool, engine *gin.Engine) {
	if !enable {
		return
	}
	engine.GET("/getApiStatistics", func(ctx *gin.Context) {
		data := make([]map[string]string, 0)
		for _, r := range engine.Routes() {
			if r.Path == "/getApiStatistics" || r.Path == "/metrics" {
				continue
			}
			api := make(map[string]string)
			api["method"] = strings.ToLower(r.Method)
			api["path"] = r.Path
			data = append(data, api)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
}
