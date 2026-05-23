# sys-monitor

A Linux system monitoring daemon built in Go that reads kernel metrics directly from /proc and exposes them in Prometheus format.

## Tech Stack
- Go — reads /proc/stat, /proc/meminfo, syscall.Statfs
- Prometheus — industry-standard metrics format
- Docker — containerized for portable deployment
- GitHub Actions — CI pipeline on every push

## Features
- Real-time CPU, memory, and disk usage monitoring
- Prometheus metrics endpoint at /metrics
- Configurable alert threshold and check interval
- Reads directly from Linux kernel /proc filesystem

## Run Locally
go run . --interval 5 --threshold 80

## Flags
--interval   Check interval in seconds (default 5)
--threshold  CPU alert threshold percent (default 80)

## Run with Docker
docker build -t sys-monitor:latest .
docker run -p 9090:9090 sys-monitor:latest

## Metrics Endpoint
curl http://localhost:9090/metrics

## Sample Output
sys_cpu_usage_percent 12.50
sys_memory_usage_percent 34.20
sys_disk_usage_percent 31.16
