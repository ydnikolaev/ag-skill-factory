# Refactoring Overview: Example Project

> **Status**: Draft
> **Date**: 2026-01-XX
> **Scope**: Full project audit
> **Depth**: Deep

## Executive Summary

This refactoring plan addresses critical technical debt identified through static analysis and Context7 best practice comparison.

**Key Findings:**
- 3 god files (>500 LOC) requiring split
- 12 files missing test coverage
- 2 layering violations (domain → infra imports)
- Average cyclomatic complexity: 8.2 (target: <10)

## Priority Ranking

| Priority | Area | Effort | Executor |
|----------|------|--------|----------|
| 🔴 Critical | `handlers/user.go` (847 LOC) | 2 days | `@backend-go-expert` |
| 🔴 Critical | `domain/order.go` imports infra | 1 day | `@backend-go-expert` |
| 🟠 High | Missing tests in `services/` | 3 days | `@backend-go-expert` |
| 🟡 Medium | `components/Dashboard.vue` (412 LOC) | 1 day | `@frontend-nuxt` |

## Modules

### Module 1: User Handler Split
- **File**: `internal/handlers/user.go` (847 LOC)
- **Problem**: God file handling all user operations
- **Solution**: Split into `user_auth.go`, `user_profile.go`, `user_admin.go`
- **Executor**: `@backend-go-expert`
- **Enforcement**: Add `max-lines: 300` to golangci-lint

### Module 2: Layering Fix
- **File**: `internal/domain/order.go`
- **Problem**: Domain imports infra package directly
- **Solution**: Introduce repository interface in domain
- **Executor**: `@backend-go-expert`
- **Enforcement**: Add depguard rule blocking infra imports in domain

### Module 3: Test Coverage
- **Scope**: `internal/services/*.go`
- **Problem**: 12 files missing tests
- **Solution**: Add unit tests with 80% coverage target
- **Executor**: `@backend-go-expert`
- **Enforcement**: CI coverage threshold 80%

## Enforcement Summary

| Type | Rule | File |
|------|------|------|
| Lint | `max-lines: 300` | `.golangci.yaml` |
| Lint | `depguard` domain rules | `.golangci.yaml` |
| CI | Coverage threshold 80% | `.github/workflows/test.yaml` |
| ADR | Handler split decision | `enforcement/adrs/adr-001-handler-split.md` |

## Risk Assessment

| Risk | Mitigation |
|------|------------|
| Handler split breaks imports | Incremental migration with aliases |
| Test coverage delays | Prioritize critical paths first |
| Team unfamiliarity with new patterns | Document in ADRs |

---

> **Next Step**: User approval, then delegate modules to executors
