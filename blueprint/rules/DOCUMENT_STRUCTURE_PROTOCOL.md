---
trigger: model_decision
description: Document naming and lifecycle. Apply when creating ANY project document.
---

# Document Structure Protocol

> **HARD RULES:**
> 1. Filename = `{work-unit}.md` (from branch name)
> 2. Path = `project/docs/active/{category}/`
> 3. **NEVER** create generic names like `tech-spec.md`

---

## Creating Documents

1. Get `work_unit` from branch: `git branch --show-current | sed 's|/|-|g'`
2. Get `doc_type` from skill outputs
3. Copy template: `_meta/_docs/templates/documents/_{doc_type}.md`
4. Save to: `active/{doc_category}/{work_unit}.md`
5. Fill placeholders: `{WORK_UNIT}`, `{DATE}`

> **No template?** See `DOCUMENT_CREATION.md` rule.

---

## Lifecycle

```
Draft → Review → Approved → Archived
  ↓        ↓         ↓          ↓
active/  review/   review/   closed/
```

| Action | What to Do |
|--------|------------|
| **Create** | Put in `active/{category}/` |
| **Submit** | Move to `review/`, set `status: Review`, call `notify_user` |
| **Archive** | (User action) Move to `closed/{type}/{work-unit}/` |

---

## Categories

| Category | Doc Types |
|----------|-----------|
| `discovery/` | discovery-brief, feature-brief |
| `product/` | user-stories, roadmap, backlog, requirements |
| `design/` | tokens, design-system, theming |
| `architecture/` | context-map, api-contracts, cli-design, webhook-config, server-config |
| `specs/` | tech-spec |
| `backend/` | service-implementation, tui-design |
| `frontend/` | ui-implementation, tma-config |
| `qa/` | test-cases, test-report |
| `infrastructure/` | deployment-guide |
| `bugs/` | debug-report |
| `refactoring/` | refactoring-overview |
| `project/` | decision-log, known-issues |

---

## Rules Summary

1. **NEVER** create document with generic name
2. **NEVER** modify files in `closed/`
3. **ALWAYS** use template when available
4. **ALWAYS** use `notify_user` when moving to `review/`
