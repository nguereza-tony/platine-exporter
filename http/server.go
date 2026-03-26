package http

import (
	"log"
	"net/http"
	"platine-exporter/metrics"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Start(addr, path string) error {

	mux := http.NewServeMux()
	handler := promhttp.HandlerFor(
		metrics.PromRegistry,
		promhttp.HandlerOpts{
			EnableOpenMetrics: false,
		})
	mux.Handle(path, handler)

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("Platine Exporter Listened on %s", addr)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Printf("Server error: %v", err)
		return err
	}

	return nil
}
