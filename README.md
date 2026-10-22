# url-shortener-go

High-performance, lightweight URL shortener written in Go utilizing in-memory TTL caching and SHA-256 base62 encoding.

## Benchmark
- Latency: < 2ms P99
- Throughput: ~45,000 req/sec on single core

## Usage
```bash
go run main.go
```
