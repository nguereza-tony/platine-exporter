package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	TotalRequest *prometheus.CounterVec
	Duration     *prometheus.HistogramVec

	LastTs      prometheus.Gauge
	LastProcess prometheus.Gauge
	RPS         prometheus.Gauge

	SlowRequests prometheus.Counter
}

func New(prefix string) *Metrics {

	m := &Metrics{}

	m.TotalRequest = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: prefix + "total_request",
		},
		[]string{"path", "method"},
	)

	m.Duration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: prefix + "request_duration",
		},
		[]string{"path", "method"},
	)

	m.LastTs = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: prefix + "last_log_timestamp",
	})

	m.LastProcess = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: prefix + "last_processing_time",
	})

	m.RPS = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: prefix + "rps",
	})

	m.SlowRequests = prometheus.NewCounter(prometheus.CounterOpts{
		Name: prefix + "api_slow_requests_total",
	})

	prometheus.MustRegister(
		m.TotalRequest,
		m.Duration,
		m.LastTs,
		m.LastProcess,
		m.RPS,
		m.SlowRequests,
	)

	return m
}
