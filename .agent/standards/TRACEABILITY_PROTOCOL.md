# Traceability Protocol

> [!CAUTION]
> **Every pipeline artifact MUST trace back to original requirements.**

## Why This Matters

Requirements get lost at each pipeline transition. Without traceability:
- `@bmad-architect` doesn't reference User Stories
- `@tech-spec-writer` doesn't verify coverage
- `@qa-lead` tests against tech-spec, not original requirements

## Required Sections in Artifacts

### 1. Upstream Documents

Every artifact MUST start with:

```markdown
## Upstream Documents
- **Input**: `project/docs/...` (the artifact this skill received)
- **Original**: `project/docs/features/...` (source of truth)
```

### 2. Requirements Checklist

Every artifact MUST include:

```markdown
## Requirements Checklist

> Source: [path to user-stories or feature brief]

| Req ID | Description | Status | Notes |
|--------|-------------|--------|-------|
| US-001 | Feature description | ✅ Covered | Section 1.1 |
| US-002 | Another feature | ⚠️ Partial | Only AC-1,2 covered |
| US-003 | Third feature | ❌ Deferred | Out of scope (approved) |
```

### 3. Status Values

| Status | Meaning |
|--------|---------|
| ✅ Covered | Fully specified in this document |
| ⚠️ Partial | Needs more detail — MUST call out to user |
| ❌ Deferred | Explicitly out of scope (with reason) |
| ➡️ Upstream | Already handled by previous skill |

## Before Handoff Checklist

1. ✅ Upstream Documents section filled
2. ✅ Requirements Checklist complete
3. ✅ No ❌ without explicit reason
4. ✅ Any ⚠️ called out to user via `notify_user`

## Responsible Skills

| Skill | Responsibility |
|-------|----------------|
| `@product-analyst` | Creates US-XXX with numbered ACs — this IS the REQ list |
| `@bmad-architect` | Maps each Implementation Phase to US-XXX |
| `@tech-spec-writer` | Checklist + gap analysis before handoff |
| `@qa-lead` | Tests against User Stories, NOT tech-spec |
