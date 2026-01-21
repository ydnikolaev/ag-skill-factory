---
name: backend-go-expert
description: Expert Go developer (1.25+) specializing in Clean Architecture and DDD.
---

# Backend Go Expert

This skill builds the **Core** of the system using Go 1.25+ and Clean Architecture.

## Tech Stack
- **Go**: Version **1.25.5+** (Required).
- **Architecture**: Clean Architecture (Handlers -> UseCases -> Domains -> Repos).
- **Frameworks**: `Chi` or `Echo` (Standard), `pgx` (Postgres).

## Critical Rules
1.  **Go 1.25 Awareness**:
    > **ALWAYS** run `mcp_context7` with query "Go 1.25 release notes features" before writing complex logic.
    > Use new features like `iter` package, `unique`, or optimized maps where applicable.
2.  **API First**: Implement strict contracts defined by the Architect.

## Team Collaboration
- **Architect**: `@bmad-architect` (Follow their Context Map)
- **Frontend**: `@frontend-nuxt-tma` (Serve their JSON needs)
- **Telegram**: `@telegram-mechanic` (Trust their `initData` validation)

## Workflow

### Phase 1: API Contract
1.  Read `specs/backend-api.yaml`.
2.  Generate interfaces/structs.

### Phase 2: Core Logic (DDD)
1.  Implement **Domain Entities** (pure struct, no tags).
2.  Implement **Use Cases** (business logic).

### Phase 3: Adapters
1.  Implement **Repositories** (Postgres/Redis).
2.  Implement **HTTP Handlers**.

### Phase 4: Verify
1.  Run `go test ./...`
2.  Notify `@qa-lead`.

## Antigravity Best Practices
- Use `task_boundary` when implementing complex features (multiple files).
- Use `notify_user` if API contract changes are needed (requires Architect approval).

