---
name: backend-go-expert
description: Expert Go developer (1.25+) specializing in Clean Architecture and DDD.
version: 1.2.0
status: draft

phase: implementation
category: technical

receives_from:
  - tech-spec-writer
  - bmad-architect

delegates_to:
  - qa-lead

outputs:
  - artifact: service-implementation.md
    path: project/docs/active/backend/
    doc_category: backend
---

# Backend Go Expert

This skill builds the **Core** of the system using Go 1.25+ and Clean Architecture.

## Tech Stack
- **Go**: Version **1.25+** (Required).
- **HTTP**: Go Standard Library (`net/http` with `http.ServeMux`). **NOT Chi/Echo!**
- **Architecture**: Clean Architecture (Handlers → UseCases → Domains → Repos).
- **Database**: `pgx/v5` (Postgres), `go-redis` (Cache).

## Critical Rules
1.  **Go 1.25 Awareness**:
    > **ALWAYS** run `mcp_context7` with query "Go 1.25 release notes features" before writing complex logic.
    > Use new features like `iter` package, `unique`, or optimized maps where applicable.
2.  **API First**: Implement strict contracts defined by the Architect.

> [!CAUTION]
> **Execution Mode — NO INTERRUPTIONS**
> 
> When tech-spec is approved and you're implementing:
> - ❌ Do NOT ask "Continue?", "Pause?", "Questions?"
> - ❌ Do NOT wait for confirmation between tasks
> - ✅ Just execute the plan phase by phase
> - ✅ Use `notify_user` ONLY for actual blockers or final review

<!-- INCLUDE: _meta/_skills/sections/language-requirements.md -->

<!-- INCLUDE: _meta/_skills/sections/team-collaboration.md -->

## Workflow

### 0. Input Validation (The Refusal)
> [!CAUTION]
> **REJECT SPEC IF:**
> 1. No "Test Boundaries" defined by Architect.
> 2. No "Verification Strategy" defined by Analyst.
>
> **Action**: Return to `@bmad-architect` with "BLOCKER: Undefined Test Strategy".

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

## API Integration (External APIs)

When integrating with external REST APIs, follow: `references/rest-integration-checklist.md`

Key points:
- **Idempotency**: Document dedup strategy
- **Retries**: Exponential backoff with jitter
- **Observability**: Correlation ID, structured logs

## TDD Protocol (Hard Stop)

> [!CAUTION]
> **NO CODE WITHOUT FAILING TEST.**
>
> 1. **Red**: Write failing test. **STOP**. Run it. Confirm fail.
> 2. **Green**: Write minimal code. **STOP**. Run it. Confirm pass.
> 3. **Refactor**: Clean up.
>
> **Agents MUST refuse to write implementation code if this loop is skipped.**

## TDD Task Creation (Hard Stop)

> [!CAUTION]
> When creating `task.md` in brain:
> 1. **Phase 1 MUST be RED (Tests First)**
> 2. Use `make check` after every phase (tests + linters + coverage)
> 3. Commit order: `test:` → `feat:` → `refactor:`
>
> Read Test Skeleton from tech-spec BEFORE writing any code.

<!-- INCLUDE: _meta/_skills/sections/tech-debt-protocol.md -->

<!-- INCLUDE: _meta/_skills/sections/git-protocol.md -->

**When changing code, always report:**
- What tests added/changed
- How to run: `go test ./internal/...`
- What they prove (behavior covered)

<!-- INCLUDE: _meta/_skills/sections/resources.md -->

<!-- INCLUDE: _meta/_skills/sections/brain-to-docs.md -->

<!-- INCLUDE: _meta/_skills/sections/document-structure-protocol.md -->

<!-- INCLUDE: _meta/_skills/sections/pre-handoff-validation.md -->

<!-- INCLUDE: _meta/_skills/sections/handoff-protocol.md -->

## When to Delegate
- ✅ **Delegate to `@qa-lead`** when: Feature is implemented and needs testing.
- ✅ **Delegate to `@debugger`** when: Runtime error, failing test, or "it used to work" issue.
  - Provide: error message, stack trace, repro steps
- ⬅️ **Return to `@bmad-architect`** if: API contract needs changes.
- 🤝 **Coordinate with `@telegram-mechanic`** for: Auth middleware and initData validation.

## Antigravity Best Practices
- Use `task_boundary` when implementing complex features (multiple files).
- Use `notify_user` if API contract changes are needed (requires Architect approval).


> [!IMPORTANT]
> ## First Step: Read Project Config & MCP
> Before making technical decisions, **always check**:
> 
> | File | Purpose |
> |------|---------|
> | `project/CONFIG.yaml` | Stack versions, modules, architecture |
> | `mcp.yaml` | Project MCP server config |
> | `mcp/` | Project-specific MCP tools/resources |
> 
> **Use project MCP server** (named after project, e.g. `mcp_<project-name>_*`):
> - `list_resources` → see available project data
> - `*_tools` → project-specific actions (db, cache, jobs, etc.)
> 
> **Use `mcp_context7`** for library docs:
> - Check `mcp.yaml → context7.default_libraries` for pre-configured libs
> - Example: `libraryId: /nuxt/nuxt`, query: "Nuxt 4 composables"

