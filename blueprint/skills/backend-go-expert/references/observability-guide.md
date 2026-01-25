# Observability Guide

Structured logging, correlation, health checks.

## 1. Structured Logging

Use JSON format for logs:
```go
slog.Info("user action",
    "user_id", userID,
    "action", "login",
    "correlation_id", correlationID,
    "duration_ms", duration.Milliseconds(),
)
```

**Required fields:**
- `level` — DEBUG, INFO, WARN, ERROR
- `msg` — human-readable message
- `correlation_id` — request trace
- `timestamp` — RFC3339

## 2. Correlation ID

Pass through entire request chain:
```go
// Middleware
func CorrelationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        corrID := r.Header.Get("X-Correlation-ID")
        if corrID == "" {
            corrID = uuid.NewString()
        }
        ctx := context.WithValue(r.Context(), CorrIDKey, corrID)
        w.Header().Set("X-Correlation-ID", corrID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

## 3. Log Levels

| Level | When to use |
|-------|-------------|
| DEBUG | Development details, verbose |
| INFO | Normal operations, audit trail |
| WARN | Recoverable issues, degraded state |
| ERROR | Failures requiring attention |

## 4. Health Endpoints

```go
// GET /health — basic liveness
// GET /health/ready — full readiness (DB, cache, etc.)
```

Response format:
```json
{
  "status": "healthy",
  "version": "1.2.3",
  "checks": {
    "database": "ok",
    "cache": "ok"
  }
}
```

## 5. Metrics (Counters)

Track at minimum:
- `http_requests_total{method, path, status}`
- `http_request_duration_seconds{method, path}`
- `db_query_duration_seconds`
- `errors_total{type}`

## 6. Request Logging

Log every request:
```
INFO | method=GET path=/api/users status=200 duration=45ms correlation_id=abc-123
```

## Quick Reference

| What | How |
|------|-----|
| Trace requests | Correlation ID header |
| Debug production | Structured JSON logs |
| Monitor health | /health endpoint |
| Measure performance | Duration logging |
