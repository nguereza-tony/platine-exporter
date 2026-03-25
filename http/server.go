package http

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Start(addr, path string) error {

	mux := http.NewServeMux()
	mux.Handle(path, promhttp.Handler())

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("Exporter started on %s", addr)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Printf("Server error: %v", err)
		return err
	}

	return nil
}
