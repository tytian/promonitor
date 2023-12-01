package monitor

import "github.com/prometheus/client_golang/prometheus"

var (
	recvRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "receive_request_total",
		Help: "Number of server requests received in total",
	}, []string{"method", "path"})
	recvServerHandlerStatus = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "receive_server_status",
		Help: "Number of",
	}, []string{"method", "path", "status"})
	recvServerHandlerSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "receive_server_seconds",
	}, []string{"method", "path", "status"})
)

func init() {
	prometheus.MustRegister(recvRequestTotal, recvServerHandlerStatus, recvServerHandlerSeconds)
}

type metricMonitor struct{}

var MetricMonitor = &metricMonitor{}

func (m *metricMonitor) RecvRequestTotal(method, path string) {
	recvRequestTotal.WithLabelValues(method, path).Inc()
}

func (m *metricMonitor) RecvServerHandlerStatus(method, path, status string) {
	recvServerHandlerStatus.WithLabelValues(method, path, status).Inc()
}

func (m *metricMonitor) RecvServerHandlerSeconds(method, path, status string, seconds float64) {
	recvServerHandlerSeconds.With(prometheus.Labels{
		"method": method,
		"path":   path,
		"status": status,
	}).Observe(seconds)
}
