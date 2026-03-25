package config

import "flag"

type Config struct {
	LogFile     string
	Addr        string
	MetricsPath string
	StateFile   string
	Prefix      string
	Workers     int
}

func Load() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.LogFile, "log", "/var/log/api.log", "log file")
	flag.StringVar(&cfg.Addr, "addr", ":9100", "listen address")
	flag.StringVar(&cfg.MetricsPath, "metrics", "/metrics", "metrics path")
	flag.StringVar(&cfg.StateFile, "state", "/tmp/state.json", "state file")
	flag.StringVar(&cfg.Prefix, "prefix", "platine_", "metrics prefix")
	flag.IntVar(&cfg.Workers, "workers", 4, "workers")

	flag.Parse()
	return cfg
}
