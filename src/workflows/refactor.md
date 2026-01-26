---
description: Full refactoring workflow from analysis to execution
---

# /refactor Workflow

Comprehensive refactoring workflow. Analyzes codebase, creates modular specs, and delegates to domain executors.

## Steps

### 1. Activate Refactor Architect

Invoke `@refactor-architect` to analyze the codebase and create refactoring spec.

The skill will:
- Interview you for scope (project/module/before-feature)
- Interview you for depth (shallow/deep/custom)
- Query Context7 for best practices
- Run static analysis
- Generate modular spec in `project/docs/refactoring/`

### 2. Review Refactoring Spec

Review the generated artifacts:

| Artifact | Purpose |
|----------|---------|
| `project/docs/refactoring/overview.md` | Master plan with priorities |
| `project/docs/refactoring/modules/*.md` | Per-domain module specs |
| `project/docs/refactoring/enforcement/*.md` | Lint rules, CI, pre-commit |

Approve or request changes before proceeding.

### 3. Execute Module Specs

After approval, delegate modules to executors (can run in parallel):

| Module Domain | Executor |
|---------------|----------|
| Go backend | `@backend-go-expert` |
| Nuxt frontend | `@frontend-nuxt` |
| CLI | `@cli-architect` |
| Telegram bot | `@telegram-mechanic` |
| TMA | `@tma-expert` |
| MCP | `@mcp-expert` |

Each executor receives their module spec from `project/docs/refactoring/modules/`.

### 4. Apply Enforcement

After refactoring is complete, apply enforcement mechanisms:

```bash
# Apply golangci-lint rules from enforcement/lint-rules.md
# Apply pre-commit hooks from enforcement/pre-commit-hooks.md
# Apply CI changes from enforcement/ci-additions.md
```

Delegate to `@devops-sre` for CI/CD changes.

### 5. Validate Improvements

// turbo

```bash
# Run lint checks
make lint

# Run tests
make test

# Check coverage
make coverage
```

Delegate to `@qa-lead` for coverage validation.

### 6. Report

Summarize refactoring results:
- Modules completed
- Enforcement applied
- Test coverage changes
- Remaining items (if any)

Update `project/docs/AGENTS.md` with final status.
