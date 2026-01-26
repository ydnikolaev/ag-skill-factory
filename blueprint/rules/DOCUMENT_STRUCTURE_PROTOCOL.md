---
trigger: model_decision
description: Document naming and lifecycle. Apply when creating ANY project document.
---

# Document Structure Protocol

> **HARD RULES:**
> 1. Check template `lifecycle:` field first
> 2. **per-feature** → filename = `{work-unit}.md`
> 3. **living** → filename = `{doc_type}.md` (fixed name)
> 4. **NEVER** create generic names for per-feature docs

---

## Document Lifecycle Types

| Lifecycle | Naming | Location | When to Use |
|-----------|--------|----------|-------------|
| **per-feature** | `{work-unit}.md` | `active/{category}/` | Created for each feature/sprint, archived when done |
| **living** | `{doc_type}.md` | `project/docs/` | One per project, updated continuously |

### Living Documents

These are long-lived project documents with fixed names:

| Doc Type | Location | Owner |
|----------|----------|-------|
| `roadmap.md` | `project/docs/` | @product-analyst |
| `backlog.md` | `project/docs/` | @product-analyst |
| `decision-log.md` | `project/docs/` | @bmad-architect |
| `known-issues.md` | `project/docs/` | @debugger |

### Per-Feature Documents

Created for each work unit (feature/sprint/fix):

```
project/docs/active/specs/feat-forum.md      ← per-feature
project/docs/backlog.md                      ← living
```

---

## Creating Documents

### Step 1: Check Template

```bash
# Find template for your doc_type
cat project/docs/templates/_{doc_type}.md

# Check lifecycle field in template
lifecycle: per-feature  # or: living
```

### Step 2: Copy & Rename

**For per-feature docs:**
```bash
work_unit=$(git branch --show-current | sed 's|/|-|g')
cp project/docs/templates/_{doc_type}.md project/docs/active/{category}/${work_unit}.md
```

**For living docs:**
```bash
cp project/docs/templates/_{doc_type}.md project/docs/{doc_type}.md
# Example: project/docs/backlog.md
```

### Step 3: Fill Placeholders

- `{WORK_UNIT}` → branch name (per-feature only)
- `{DATE}` → current date

---

## Lifecycle Flow

### Per-Feature Documents

```
Draft → Review → Approved → Archived
  ↓        ↓         ↓          ↓
active/  review/   review/   closed/
```

### Living Documents

```
Draft → Updated → Updated → Updated...
         ↓          ↓          ↓
      project/docs/ (stays in place)
```

---

## Categories (Per-Feature Docs)

| Category | Doc Types |
|----------|-----------|
| `discovery/` | discovery-brief, feature-brief |
| `product/` | user-stories |
| `design/` | tokens, design-system, theming |
| `architecture/` | context-map, api-contracts, cli-design, webhook-config, server-config |
| `specs/` | tech-spec, requirements |
| `backend/` | service-implementation, tui-design |
| `frontend/` | ui-implementation, tma-config |
| `qa/` | test-cases, test-report |
| `infrastructure/` | deployment-guide |
| `bugs/` | debug-report |
| `refactoring/` | refactoring-overview |

---

## Rules Summary

1. **ALWAYS** check template `lifecycle:` field before creating
2. **NEVER** use work-unit naming for living docs
3. **NEVER** use fixed naming for per-feature docs
4. **ALWAYS** use `notify_user` when moving to `review/`
