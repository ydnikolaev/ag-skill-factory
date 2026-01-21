# Backend Go Expert Checklist

## Before Starting
- [ ] Read `specs/backend-api.yaml` from Architect
- [ ] Check Go 1.25 features via `mcp_context7` query: "Go 1.25 release notes"

## Code Quality
- [ ] Domain entities are pure structs (no framework tags)
- [ ] Use cases contain only business logic
- [ ] Repositories use `pgx` or `sqlx` correctly
- [ ] HTTP handlers are thin (delegate to use cases)

## Testing
- [ ] `go test ./...` passes
- [ ] Test coverage > 70%
- [ ] Integration tests for database operations

## API Contract
- [ ] All endpoints match `specs/backend-api.yaml`
- [ ] JSON response structure is consistent
- [ ] Error responses follow standard format

## Final
- [ ] Notified `@qa-lead` for testing
- [ ] No hardcoded secrets in code
