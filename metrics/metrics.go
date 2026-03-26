package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	RequestTotal    *prometheus.CounterVec
	SlowRequests    *prometheus.CounterVec
	RequestDuration *prometheus.HistogramVec
	DBDuration      *prometheus.HistogramVec

	LastProcess prometheus.Gauge
	LastTs      prometheus.Gauge
	RPS         prometheus.Gauge
}

var PromRegistry = prometheus.NewRegistry()

func New(prefix string) *Metrics {

	m := &Metrics{}

	// Important: status = "2xx", "4xx", "5xx"
	m.RequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: prefix + "requests_total",
			Help: "Total API requests",
		},
		[]string{"path", "method", "status"},
	)

	m.RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: prefix + "request_duration",
			Help: "Request duration in ms",
			Buckets: []float64{
				5, 10, 25, 50, 100, 250, 500, 1000, 2000, 5000,
			},
		},
		[]string{"path", "method", "status"},
	)

	m.DBDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: prefix + "db_duration",
			Help: "Database duration in ms",
			Buckets: []float64{
				1, 5, 10, 25, 50, 100, 250, 500,
			},
		},
		[]string{"path", "method", "status"},
	)

	m.SlowRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: prefix + "slow_requests_total",
		},
		[]string{"path", "method", "status"},
	)

	m.LastTs = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: prefix + "last_log_timestamp",
	})

	m.LastProcess = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: prefix + "last_processing_time",
	})

	PromRegistry.MustRegister(
		m.RequestTotal,
		m.RequestDuration,
		m.DBDuration,
		m.LastTs,
		m.LastProcess,
		m.SlowRequests,
	)

	return m
}
