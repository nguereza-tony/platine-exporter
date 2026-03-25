package main

import (
	"time"

	"platine-exporter/config"
	"platine-exporter/dedup"
	"platine-exporter/metrics"
	"platine-exporter/reader"
	"platine-exporter/state"
	"platine-exporter/worker"

	httpSrv "platine-exporter/http"
)

func main() {

	cfg := config.Load()

	jobs := make(chan []byte, 100000)

	m := metrics.New(cfg.Prefix)
	d := dedup.New(500000, 30*time.Second)

	st := state.Load(cfg.StateFile)

	var offset = st.Offset
	var inode = st.Inode

	for i := 0; i < cfg.Workers; i++ {
		go worker.Start(jobs, m, d)
	}

	go reader.Tail(cfg.LogFile, jobs, &offset, &inode)

	go func() {
		for {
			state.Save(cfg.StateFile, &state.State{
				Offset: offset,
				Inode:  inode,
			})
			time.Sleep(5 * time.Second)
		}
	}()

	httpSrv.Start(cfg.Addr, cfg.MetricsPath)
}
