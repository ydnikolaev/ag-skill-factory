# Backend Performance Guide

Optimization patterns for Go services.

## 1. N+1 Query Prevention

❌ Bad:
```go
for _, user := range users {
    orders, _ := repo.GetOrdersByUserID(user.ID) // N queries!
}
```

✅ Good:
```go
orders, _ := repo.GetOrdersByUserIDs(userIDs) // 1 query
ordersByUser := groupByUserID(orders)
```

## 2. Connection Pooling

```go
// pgx pool config
config.MaxConns = 25
config.MinConns = 5
config.MaxConnLifetime = time.Hour
```

## 3. Caching Strategy

| Data | Cache TTL | Invalidation |
|------|-----------|--------------|
| User profile | 5 min | On update |
| Config | 1 hour | On deploy |
| Session | Request | - |

```go
// Check cache first
if cached, ok := cache.Get(key); ok {
    return cached
}
// Fetch and cache
data := fetchFromDB()
cache.Set(key, data, 5*time.Minute)
```

## 4. Batch Operations

```go
// Insert many in one query
_, err := tx.CopyFrom(ctx, pgx.Identifier{"users"}, columns, pgx.CopyFromSlice(...))
```

## 5. Query Optimization

- [ ] Use EXPLAIN ANALYZE on slow queries
- [ ] Add indexes for WHERE/JOIN columns
- [ ] Limit SELECT fields (no `SELECT *`)
- [ ] Use pagination

## Quick Reference

| Problem | Solution |
|---------|----------|
| Slow queries | Add index, EXPLAIN |
| N+1 | Batch fetch |
| Memory | Stream large results |
| Latency | Cache hot data |
