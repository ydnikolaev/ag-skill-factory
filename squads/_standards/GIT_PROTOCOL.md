# Git Protocol 🐙

This document defines the Git workflow for all Antigravity projects.

## Branching Strategy (Solo GitHub Flow)

> **Simple, Deployable, Feature-Based.**

### Branches
| Branch | Purpose |
|--------|---------|
| `main` | Always deployable. Protected. |
| `feat/<name>` | New features |
| `fix/<name>` | Bug fixes |
| `chore/<name>` | Maintenance, refactoring |

### Workflow
1.  **Create Branch**: `git checkout -b feat/<name>` from `main`.
2.  **Work**: Make commits following Conventional Commits (below).
3.  **Merge**: `git checkout main && git merge feat/<name>` (squash optional).
4.  **Push**: `git push origin main`.
5.  **Delete Branch**: `git branch -d feat/<name>`.

> [!CAUTION]
> **Forbidden**: Pushing directly to `main` without a branch.
> **Reason**: Preserves clean history and allows easy rollback.

## Commit Messages (Conventional Commits)

> **Semantic, Automatable.**

### Format
```text
<type>(<scope>): <subject>

<body>
```

### Types
| Type | Description | SemVer Impact |
|------|-------------|---------------|
| `feat` | New feature | Minor (0.X.0) |
| `fix` | Bug fix | Patch (0.0.X) |
| `chore` | Maintenance | None |
| `docs` | Documentation | None |
| `refactor` | Code restructure | None |
| `test` | Test changes | None |
| `perf` | Performance | Patch |

### Rules
1.  **Imperative Mood**: "Add feature" not "Added feature".
2.  **Scope**: Component name (e.g. `auth`, `api`, `ui`).
3.  **Subject**: Max 50 chars, no period at end.

### Examples
```text
feat(auth): add jwt middleware
fix(ui): resolve hydration error on profile page
chore: update dependencies
refactor(api): extract validation logic to helper
```

## Breaking Changes

Use `!` after type or `BREAKING CHANGE:` in footer:
```text
feat(api)!: change user endpoint response format

BREAKING CHANGE: Response now returns `user` object instead of flat fields.
```

## Atomic Commits

> [!CAUTION]
> **One commit = One logical change.**
> - ❌ "Add feature and fix bug and update deps"
> - ✅ Three separate commits

## Enforcement

Agents MUST follow this protocol. Reject:
-   Commits like "wip", "fix", "update", "stuff".
-   Huge commits with unrelated changes.
-   Direct pushes to `main` without a branch.
