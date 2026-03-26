# 🚀 Platine Exporter

Platine Exporter is a high-performance log-based Prometheus exporter written in Go.
It parses API traffic logs in real time and exposes useful metrics for monitoring and observability.

## ✨ Features

* ⚡ High-performance log processing (worker pool + channels)
* 📊 Native Prometheus metrics (`client_golang`)
* 🔁 Log rotation support (copytruncate & rename)
* 💾 Persistent state (offset + resume after restart)
* 🧠 Deduplication (LRU + TTL)
* 🧵 Concurrent processing (multi-workers)
* 📉 Request duration & slow request tracking
* 📈 Real-time metrics (RPS, lag, timestamps)
* 🔌 Configurable via CLI flags
* 🖥️ Cross-platform (Windows, Linux, MacOS)

## ⚙️ Requirements

* Go 1.25+

## 📦 Installation

Clone the repository:

```bash
git clone https://github.com/nguereza-tony/platine-exporter.git
cd platine-exporter
```

Install dependencies:

```bash
go mod tidy
```

## ▶️ Run

```bash
go run ./cmd
```

## 🏗️ Build

```bash
go build -o platine-exporter ./cmd

# Build all
make build
```

Run binary:

```bash
./platine-exporter
```

On Windows:

```bash
platine-exporter.exe
```

## ⚙️ Configuration

The exporter is configured using CLI flags:

| Flag       | Default            | Description                           |
| ---------- | ------------------ | --------------------------------------|
| `-log`     | `/var/log/api.log` | Log file path                         |
| `-addr`    | `:9100`            | HTTP listen address                   |
| `-metrics` | `/metrics`         | Metrics endpoint                      |
| `-state`   | `/tmp/state.json`  | State file (offset persistence)       |
| `-prefix`  | `platine_api_`     | Metrics prefix                        |
| `-workers` | `1`                | Number of workers                     |
| `-slow-duration` | `500`        | Response slow duration in millisecond |

Example:

```bash
./platine-exporter -log=app.log -addr=:8080 -workers=8
```

## 📊 Metrics

The exporter provides the following Prometheus metrics (Note suppose prefix is `platine_api_`):

### Request Counters

* `platine_api_requests_total{path,method,status}`

### Latency

* `platine_api_request_duration{path,method,status}`
* `platine_api_db_duration{path,method,status}`

### Performance

* `platine_api_last_log_timestamp`
* `platine_api_last_processing_time`

### Alerts

* `platine_api_slow_requests_total{path,method,status}`

## 🔍 Example Log Format

```
ts=2026-03-24 17:33:11.821791	req=8afc6afb	ip=127.0.0.1	method=GET	path=/v1/documents/5	status=200	user=1	time=116ms	db_time=6.59ms
```

## 📡 Prometheus Configuration

Add the exporter to your Prometheus config:

```yaml
scrape_configs:
  - job_name: 'platine-exporter'
    static_configs:
      - targets: ['localhost:9100']
```


## 🔄 Log Rotation Handling

Supported:

* copytruncate (logrotate)
* file replacement (rename)

The exporter automatically:

* detects file truncation
* resets offset if needed
* resumes reading safely

## 💾 State Persistence

The exporter stores its state (offset + file tracking) to avoid reprocessing logs after restart.

Example:

```json
{
  "offset": 987654
}
```


## ⚠️ Notes

* The exporter uses file size + offset for portability
* Avoid extremely high log rates without tuning workers

## 🚀 Roadmap

* mmap-based reader (ultra performance)
* zero-allocation parser (unsafe)
* lock-free ring buffer
* advanced metrics aggregation
* distributed exporter mode


## 🤝 Contributing

Contributions are welcome!

Feel free to open issues or submit pull requests.

## 📄 License

MIT License
