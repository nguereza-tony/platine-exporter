package config

import "flag"

type Config struct {
	LogFile      string
	Addr         string
	MetricsPath  string
	StateFile    string
	Prefix       string
	Workers      int
	SlowDuration int
}

func Load() *Config {
	cfg := &Config{}

	flag.StringVar(
		&cfg.LogFile,
		"log",
		"/var/log/api.log",
		"Path to the API log file to monitor. The exporter continuously reads this file and parses new entries in real time. Supports log rotation (copytruncate and rename).",
	)

	flag.StringVar(
		&cfg.Addr,
		"addr",
		":9100",
		"HTTP server listen address in the format [host]:port. Example: ':9100' (all interfaces) or '127.0.0.1:9100'.",
	)

	flag.StringVar(
		&cfg.MetricsPath,
		"metrics",
		"/metrics",
		"HTTP endpoint path where Prometheus metrics are exposed. Default is '/metrics'.",
	)

	flag.StringVar(
		&cfg.StateFile,
		"state",
		"/tmp/state.json",
		"Path to the state file used to persist the current read offset and file tracking information. This allows the exporter to resume from the last position after restart and avoid reprocessing logs.",
	)

	flag.StringVar(
		&cfg.Prefix,
		"prefix",
		"platine_",
		"Prefix added to all exported Prometheus metrics. Useful to namespace metrics when running multiple exporters.",
	)

	flag.IntVar(
		&cfg.Workers,
		"workers",
		4,
		"Number of concurrent worker goroutines used to process log entries. Increasing this value improves throughput but may increase CPU usage.",
	)

	flag.IntVar(
		&cfg.SlowDuration,
		"slow-duration",
		500,
		"Slow duration in millisecond",
	)

	flag.Parse()
	return cfg
}
