package monitor

import "github.com/prometheus/client_golang/prometheus"

var (
	recvRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "receive_request_total",
		Help: "Number of server requests received in total",
	}, []string{"type", "method", "path"})
	recvServerHandlerStatus = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "receive_server_status",
	}, []string{"type", "method", "path", "status"})
	recvServerHandlerSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "receive_server_seconds",
	}, []string{"type", "method", "path", "status"})
)

func init() {
	prometheus.MustRegister(recvRequestTotal, recvServerHandlerStatus, recvServerHandlerSeconds)
}

type metricMonitor struct{}

var MetricMonitor = &metricMonitor{}

func (m *metricMonitor) RecvRequestTotal(metricType, method, path string) {
	recvRequestTotal.WithLabelValues(metricType, method, path).Inc()
}

func (m *metricMonitor) RecvServerHandlerStatus(metricType, method, path, status string) {
	recvServerHandlerStatus.WithLabelValues(metricType, method, path, status).Inc()
}

func (m *metricMonitor) RecvServerHandlerSeconds(metricType, method, path, status string, seconds float64) {
	recvServerHandlerSeconds.With(prometheus.Labels{
		"type":   metricType,
		"method": method,
		"path":   path,
		"status": status,
	}).Observe(seconds)
}
