package worker

import (
	"time"

	"platine-exporter/dedup"
	"platine-exporter/metrics"
	"platine-exporter/parser"
)

var pool = make(chan *parser.Entry, 10000)

func getEntry() *parser.Entry {
	select {
	case e := <-pool:
		return e
	default:
		return &parser.Entry{}
	}
}

func putEntry(e *parser.Entry) {
	select {
	case pool <- e:
	default:
	}
}

func Start(jobs <-chan []byte, m *metrics.Metrics, d *dedup.Dedup) {

	for line := range jobs {

		if d.Seen(line) {
			continue
		}

		e := getEntry()

		parser.Parse(line, e)

		m.TotalRequest.WithLabelValues(string(e.Path), string(e.Method)).Inc()

		m.Duration.WithLabelValues(string(e.Path), string(e.Method)).Observe(e.Time)

		if e.Time >= 0.5 {
			m.SlowRequests.Inc()
		}

		m.LastTs.Set(float64(time.Now().Unix()))

		putEntry(e)
	}
}
