---
trigger: model_decision
description: Document lifecycle and naming conventions. Apply when creating or moving ANY project document.
---

# Document Structure Protocol

> Apply when creating, moving, or archiving project documents.

---

## Quick Reference

| Action | Path Pattern | Example |
|--------|--------------|---------|
| **Create** | `active/{category}/{work-unit}.md` | `active/specs/feat-forum.md` |
| **Review** | `review/{category}/{work-unit}.md` | `review/specs/feat-forum.md` |
| **Archive** | `closed/{type}/{work-unit}/` | `closed/features/feat-forum/` |

---

## Naming Convention

**Format:** `{category}/{work-unit}.md`

```
active/specs/feat-forum-topics.md     ✅ Correct
active/specs/tech-spec.md             ❌ Wrong (no work-unit)
```

**Work Unit = Branch Name:**
- `feat/forum-topics` → `feat-forum-topics`
- `fix/login-bug` → `fix-login-bug`
- Sprint → `sprint-03`

---

## Document Frontmatter

**Required in every document:**

```yaml
---
status: Draft | Review | Approved
owner: @skill-name
work_unit: feat-forum-topics

upstream:
  - doc: user-stories.md
    owner: @product-analyst
downstream:
  - doc: service-implementation.md
    owner: @backend-go-expert

created: 2026-01-26
updated: 2026-01-26
---
```

---

## Lifecycle

```
Draft → Review → Approved → Archived
  ↓        ↓         ↓          ↓
active/  review/   review/   closed/
```

| Status | Folder | Action |
|--------|--------|--------|
| Draft | `active/` | Being worked on |
| Review | `review/` | Ready for approval |
| Approved | `review/` | User said "looks good" |
| Archived | `closed/` | Moved by user or current skill |

---

## Registry System

### Main Registry (`ARTIFACT_REGISTRY.md`)
Lists all work units with status. Light-weight.

### Per-Work-Unit Module (`registry/{work-unit}.md`)  
Full document list for each work unit. Detailed.

```
project/docs/
├── ARTIFACT_REGISTRY.md          # Light: work units list
├── registry/
│   └── feat-forum-topics.md      # Heavy: all docs for feature
└── active/...
```

---

## Skill Workflow

### Creating a Document

1. Use document template from your skill (see skill SKILL.md)
2. Create in `active/{category}/{work-unit}.md`
3. Add to `registry/{work-unit}.md`

### Completing a Document

1. Set `status: Review` in frontmatter
2. Move to `review/{category}/{work-unit}.md`
3. Update `registry/{work-unit}.md`
4. Use `notify_user` for approval

### After User Approval

1. Move document to `closed/{type}/{work-unit}/`
2. Updates main `ARTIFACT_REGISTRY.md`

---

## Categories

| Category | Documents |
|----------|-----------|
| `discovery/` | discovery-brief |
| `product/` | user-stories, roadmap |
| `design/` | design-tokens, design-system |
| `architecture/` | context-map, api-contracts |
| `specs/` | tech-spec, requirements |
| `backend/` | service-implementation |
| `frontend/` | ui-implementation, tma-config |
| `qa/` | test-cases, test-report |

---

## Rules

1. **Never** create document without `work_unit` in name
2. **Never** modify files in `closed/`
3. **Always** update `registry/{work-unit}.md` when document status changes
4. **Always** follow document frontmatter format shown above
