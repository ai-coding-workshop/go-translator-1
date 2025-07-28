# Performance Optimization Guide

This document provides guidance on optimizing the performance of the Translation Service.

## Table of Contents
- [Benchmarking](#benchmarking)
- [Profiling](#profiling)
- [Optimization Strategies](#optimization-strategies)
- [Caching](#caching)
- [Connection Pooling](#connection-pooling)
- [Resource Management](#resource-management)

## Benchmarking

The project includes benchmarking tools to measure performance:

1. **Go benchmarks** - Run with `go test -bench=.`:
   ```bash
   go test -bench=.
   ```

2. **HTTP benchmark script** - Located in `scripts/benchmark.sh`:
   ```bash
   ./scripts/benchmark.sh
   ```

## Profiling

Use Go's built-in profiling tools to identify performance bottlenecks:

```bash
# CPU profiling
go test -cpuprofile=cpu.prof -bench=.

# Memory profiling
go test -memprofile=mem.prof -bench=.

# View profiles
go tool pprof cpu.prof
go tool pprof mem.prof
```

## Optimization Strategies

### 1. Concurrent Request Handling
The service already uses Go's concurrent features, but you can adjust the server configuration:

```go
server := &http.Server{
    Addr:        ":" + cfg.ServerPort,
    Handler:     router,
    ReadTimeout: 30 * time.Second,
    // Increase these values for higher throughput
    WriteTimeout: 30 * time.Second,
    IdleTimeout:  60 * time.Second,
}
```

### 2. HTTP Client Optimization
The HTTP clients used for LLM APIs can be optimized with connection pooling:

```go
client := &http.Client{
    Timeout: 30 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     90 * time.Second,
    },
}
```

### 3. Request Batching
For high-volume scenarios, consider batching multiple translation requests:

```go
type BatchTranslationRequest struct {
    Requests []TranslationRequest `json:"requests"`
}
```

## Caching

Implement caching for frequently requested translations:

### In-Memory Caching
```go
import "github.com/patrickmn/go-cache"

var translationCache = cache.New(5*time.Minute, 10*time.Minute)

func getCachedTranslation(text, model string) (*TranslationResponse, bool) {
    key := text + ":" + model
    if cached, found := translationCache.Get(key); found {
        return cached.(*TranslationResponse), true
    }
    return nil, false
}
```

### Redis Caching
For distributed deployments, use Redis:

```go
import "github.com/go-redis/redis/v8"

var redisClient *redis.Client

func init() {
    redisClient = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
}
```

## Connection Pooling

Configure connection pooling for external services:

### Database Connection Pooling
If adding database storage:
```go
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(25)
db.SetConnMaxLifetime(5 * time.Minute)
```

### HTTP Client Connection Pooling
Already implemented in LLM clients with Transport settings.

## Resource Management

### Memory Management
- Use object pools for frequently allocated objects
- Implement proper garbage collection monitoring
- Use streaming for large responses when possible

### Goroutine Management
- Limit concurrent requests with semaphores
- Implement proper timeout handling
- Use context cancellation for long-running operations

```go
// Limit concurrent translations
var semaphore = make(chan struct{}, 100) // Max 100 concurrent translations

func (ts *TranslatorService) Translate(ctx context.Context, req *TranslationRequest) (*TranslationResponse, error) {
    select {
    case semaphore <- struct{}{}:
        defer func() { <-semaphore }()
    case <-ctx.Done():
        return nil, ctx.Err()
    }

    // Translation logic here
}
```

## Monitoring

### Metrics Collection
Add metrics for performance monitoring:
- Request latency
- Error rates
- Throughput
- Resource utilization

### Health Checks
The service includes a `/health` endpoint for monitoring:
```bash
curl http://localhost:8080/health
```

## Scaling Strategies

### Horizontal Scaling
- Run multiple instances behind a load balancer
- Use shared caching (Redis) for consistency
- Implement sticky sessions if needed

### Vertical Scaling
- Increase server resources (CPU, memory)
- Optimize container resource limits (Docker)
- Tune OS-level network settings

## Best Practices

1. **Monitor Performance Continuously**
   - Set up alerts for performance degradation
   - Track key metrics over time
   - Regular performance testing

2. **Optimize for Common Cases**
   - Profile real-world usage patterns
   - Optimize the most frequently used paths
   - Cache common translations

3. **Plan for Growth**
   - Design for horizontal scaling from the start
   - Implement circuit breakers for external services
   - Use asynchronous processing for non-critical tasks
