---
trigger: model_decision
description: Track TODOs and workarounds in TECH_DEBT.md registry. Apply when creating temporary solutions, hardcoded values, or deferred work.
---

# Technical Debt Protocol

> [!CAUTION]
> **When you create a workaround, TODO, or temporary solution — TRACK IT!**

## Why This Matters

TODOs and workarounds disappear between sessions. Without tracking:
- They become hidden bugs
- Next developer doesn't know they exist
- Technical debt accumulates silently

## The Rule

**Every workaround MUST be registered in `project/docs/TECH_DEBT.md`.**

## How to Track

### 1. In Code
Add comment with ID:
```go
// TODO(TD-001): Replace hardcoded clubs with ClubRepository
clubs := []string{"Club A", "Club B"}
```

### 2. In Registry
Add entry to `project/docs/TECH_DEBT.md`:

```markdown
| ID | Created | Skill | Description | Status |
|----|---------|-------|-------------|--------|
| TD-001 | 2026-01-24 | @debugger | Hardcoded clubs in InterviewUseCase | Open |
```

## Registry Template

Create `project/docs/TECH_DEBT.md` if it doesn't exist:

```markdown
# Technical Debt Registry

> Track all TODOs, workarounds, and deferred work.

| ID | Created | Skill | Description | Status |
|----|---------|-------|-------------|--------|
| TD-001 | YYYY-MM-DD | @skill-name | Description | Open |

## Status Legend
- **Open** — Not addressed yet
- **In Progress** — Being worked on
- **Resolved** — Fixed, can be removed
```

## Forbidden Patterns

❌ `// TODO: fix later` — no ID, no registry
❌ `// FIXME` — what exactly? when? who?
❌ Hardcoded values without comment
❌ Temporary mocks without documentation

## Pipeline Integration

### At Handoff (Every Skill)
Before handing off to next skill:
1. Check if you created any workarounds
2. If yes → register in `TECH_DEBT.md`
3. If no → proceed

### At QA (Quality Gate)
`@qa-lead` checks:
1. Scan for unregistered `// TODO` comments
2. Flag any missing registry entries
3. Block release if critical debt untracked

## Cleanup Triggers

Tech debt should be resolved when:
- Feature work touches the same file
- Sprint specifically targets debt reduction
- Debt causes a bug (forced cleanup)
