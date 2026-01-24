# TDD Quick Reference

Shared guide for all skills. Red → Green → Refactor.

## The Process

| Step | Description | Output |
|------|-------------|--------|
| 1. Red | Write failing test that defines expected behavior | Test fails ❌ |
| 2. Green | Implement minimal code to pass | Test passes ✅ |
| 3. Refactor | Clean up code, keep tests passing | All green ✅ |
| 4. Verify | Run full test suite | Suite passes |

## When to Apply

| Situation | TDD Required? |
|-----------|---------------|
| New feature | ✅ Yes — test defines behavior |
| Bug fix | ✅ Yes — regression test |
| Refactor | ✅ Yes — protect behavior first |
| Tiny change | ⚠️ At least document verification |

## Test Naming

Name by **behavior**, not implementation:
- ✅ `TestUserCannotLoginWithWrongPassword`
- ❌ `TestPasswordCheck`

## Reporting (Required)

When changing code, always include:
```markdown
### Tests
- **Added**: `TestUserLogin_FailsOnInvalidPassword`
- **Run**: `go test ./internal/auth/...`
- **Proves**: Invalid passwords are rejected with 401
```

## Go Commands
```bash
go test ./...                    # All tests
go test ./internal/domain/...    # Specific package
go test -v -run TestName         # Single test
go test -cover ./...             # Coverage
```

## Frontend Commands
```bash
npm test                         # All tests
npm test -- --filter UserCard    # Specific
npm run test:coverage            # Coverage
```

## Quick Decision Tree

```
Is this a code change?
  │
  ├─ Yes → Does it change behavior?
  │         ├─ Yes → Write test first (TDD)
  │         └─ No  → Ensure existing tests pass
  │
  └─ No  → No test needed
```
