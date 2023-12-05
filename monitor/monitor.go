package monitor

import "github.com/prometheus/client_golang/prometheus"

var (
	serverHandleRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "server_handle_request_total",
		Help: "Number of server requests received in total",
	}, []string{"type", "method", "path"})
	serverHandleRequestStatus = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "server_handle_request_status",
	}, []string{"type", "method", "path", "status"})
	serverHandleRequestSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "server_handle_request_seconds",
	}, []string{"type", "method", "path", "status"})
	clientHandleRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "client_handle_request_total",
	}, []string{"type", "name", "op", "peer"})
	clientHandleRequestSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "client_handle_seconds",
	}, []string{"type", "name", "op", "peer"})
	clientHandleRequestStatus = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "client_handle_request_status",
	}, []string{"type", "name", "op", "peer", "status"})
)

func init() {
	prometheus.MustRegister(
		serverHandleRequestTotal, serverHandleRequestStatus, serverHandleRequestSeconds,
		clientHandleRequestTotal, clientHandleRequestStatus, clientHandleRequestSeconds,
	)
}

type metricMonitor struct{}

var MetricMonitor = &metricMonitor{}

func (m *metricMonitor) ServerHandleRequestTotal(metricType, method, path string) {
	serverHandleRequestTotal.WithLabelValues(metricType, method, path).Inc()
}

func (m *metricMonitor) ServerHandleRequestStatus(metricType, method, path, status string) {
	serverHandleRequestStatus.WithLabelValues(metricType, method, path, status).Inc()
}

func (m *metricMonitor) ServerHandlerRequestSeconds(metricType, method, path, status string, seconds float64) {
	serverHandleRequestSeconds.With(prometheus.Labels{
		"type":   metricType,
		"method": method,
		"path":   path,
		"status": status,
	}).Observe(seconds)
}

func (m *metricMonitor) ClientHandleRequestTotal(metricType, name, op, peer string) {
	clientHandleRequestTotal.WithLabelValues(metricType, name, op, peer).Inc()
}

func (m *metricMonitor) ClientHandleRequestSeconds(metricType, name, op, peer string, seconds float64) {
	clientHandleRequestSeconds.WithLabelValues(metricType, name, op, peer).Observe(seconds)
}

func (m *metricMonitor) ClientHandleRequestStatus(metricType, name, op, peer string, status string) {
	clientHandleRequestStatus.WithLabelValues(metricType, name, op, peer, status).Inc()
}
