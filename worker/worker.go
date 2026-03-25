package worker

import (
	"time"

	"platine-exporter/dedup"
	"platine-exporter/metrics"
	"platine-exporter/parser"
)

var pool = make(chan *parser.LogEntry, 10000)

func getEntry() *parser.LogEntry {
	select {
	case e := <-pool:
		return e
	default:
		return &parser.LogEntry{}
	}
}

func putEntry(e *parser.LogEntry) {
	select {
	case pool <- e:
	default:
	}
}

func statusClass(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "2xx"
	case code >= 400 && code < 500:
		return "4xx"
	case code >= 500:
		return "5xx"
	default:
		return "other"
	}
}

func Start(jobs <-chan []byte, m *metrics.Metrics, d *dedup.Dedup, slowDuration int) {

	for line := range jobs {

		if d.Exists(line) {
			continue
		}

		e := getEntry()

		parser.Parse(line, e)

		m.RequestTotal.WithLabelValues(string(e.Path), string(e.Method), statusClass(e.Status)).Inc()

		m.RequestDuration.WithLabelValues(string(e.Path), string(e.Method)).Observe(e.Time)
		m.DBDuration.WithLabelValues(string(e.Path), string(e.Method)).Observe(e.DBTime)

		if e.Time > float64(slowDuration) {
			m.SlowRequests.WithLabelValues(string(e.Path), string(e.Method)).Inc()
		}

		m.LastTs.Set(float64(time.Now().Unix()))

		putEntry(e)
	}
}
