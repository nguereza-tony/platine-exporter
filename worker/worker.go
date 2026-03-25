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

func Start(jobs <-chan []byte, m *metrics.Metrics, d *dedup.Dedup) {

	for line := range jobs {

		if d.Exists(line) {
			continue
		}

		e := getEntry()

		parser.Parse(line, e)

		m.TotalRequest.WithLabelValues(string(e.Path), string(e.Method)).Inc()

		m.Duration.WithLabelValues(string(e.Path), string(e.Method)).Observe(e.Time)

		if e.Time >= 500 {
			m.SlowRequests.Inc()
		}

		m.LastTs.Set(float64(time.Now().Unix()))

		putEntry(e)
	}
}
