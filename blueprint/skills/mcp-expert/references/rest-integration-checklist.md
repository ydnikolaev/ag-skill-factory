# REST API Integration Checklist

Use this checklist when building API integrations, ETL workflows, or system-to-system sync.

## 1. Contract Definition
- [ ] Define inputs (format, required fields, validation)
- [ ] Define outputs (shape, destination)
- [ ] Define success criteria
- [ ] Define non-goals (what NOT to do)

## 2. Authentication
- [ ] Identify auth type: API key, OAuth2, JWT, mTLS
- [ ] Never hardcode secrets (use env vars / secret manager)
- [ ] Plan token refresh (OAuth2)

## 3. Idempotency & Deduplication
Pick one:
- [ ] Use provider idempotency keys
- [ ] Use stable external IDs for upserts
- [ ] Keep local mapping (source ID → target ID)
- [ ] Use deterministic hashes for dedupe

**Document strategy explicitly!**

## 4. Pagination & Incremental Sync
- [ ] Detect pagination style: `next` link, cursor, page+limit
- [ ] Set max pages / max time limits
- [ ] Prefer incremental sync (`updated_since`, ETag)
- [ ] Handle out-of-order updates

## 5. Retries & Timeouts
- [ ] Set connect/read timeouts
- [ ] Retry on: network failures, 429, 5xx
- [ ] Exponential backoff with jitter
- [ ] Do NOT retry on 4xx (except 408/409/429)
- [ ] Cap retries, surface failures

## 6. Rate Limits
- [ ] Respect `Retry-After` header
- [ ] Adaptive backoff on 429
- [ ] Consider batch endpoints
- [ ] Avoid bursty concurrency

## 7. Data Mapping & Validation
- [ ] Explicit mapping layer (source → normalized → target)
- [ ] Validate required fields and types
- [ ] Normalize formats (dates, enums, currency)
- [ ] Handle nullability and partial payloads
- [ ] Record rejected records (don't silently drop)

## 8. Error Handling Strategy
Per error class, choose:
- [ ] **Skip with log** (non-critical)
- [ ] **Retry** (transient)
- [ ] **Quarantine** (store for later)
- [ ] **Fail the run** (systemic)

Report clear summary at end.

## 9. Observability
Minimum:
- [ ] Run ID / correlation ID
- [ ] Per-request logs: method, path, status, latency
- [ ] Counters: processed, created, updated, skipped, failed
- [ ] Prefer structured logs (JSON)

## 10. Webhooks (if applicable)
- [ ] Verify signature
- [ ] Handle replay (event ID idempotency)
- [ ] Respond quickly, process async
- [ ] Store raw event payloads

## 11. Safety Controls
- [ ] Dry-run mode (no writes)
- [ ] Limit scope (max records per run)
- [ ] Kill switch config flag
- [ ] Backup/rollback plan for destructive ops
