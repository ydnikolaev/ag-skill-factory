# Document Schema

> Entry point for AI agents creating document templates.

## Quick Start

1. Copy existing template as base (e.g., `_tech-spec.md`)
2. Update frontmatter: owner, upstream, downstream
3. Define sections appropriate for doc type
4. Run `make doc-validate` to check

## Template Structure

```markdown
---
status: Draft
owner: @skill-name
lifecycle: per-feature
work_unit: {WORK_UNIT}

upstream:
  - doc_type: context-map
    owner: @bmad-architect
downstream:
  - skill: @backend-go-expert

created: {DATE}
updated: {DATE}
---

# Doc Type: {WORK_UNIT}

## Overview
<!-- Brief description -->

## Details
<!-- Main content -->

## Checklist
- [ ] Item 1
```

## Key Files

| File | Purpose |
|------|---------|
| `document-schema.yaml` | Core schema |
| `enums.yaml` | Statuses, lifecycles, categories |
| `examples/` | Example templates |

## Frontmatter Fields

| Field | Required | Description |
|-------|----------|-------------|
| `status` | ✅ | Draft, In Progress, Review, Approved, Archived |
| `owner` | ✅ | @skill-name that creates this doc |
| `lifecycle` | ✅ | per-feature, per-work-unit, living |
| `work_unit` | ✅ | Work unit ID |
| `upstream` | ❌ | Docs this depends on |
| `downstream` | ❌ | Skills/docs that consume this |
| `created` | ✅ | Creation date |
| `updated` | ✅ | Last update date |

## Related Files

- [doc-types.yaml](../../doc-types.yaml) — Auto-generated doc type registry
- [DOCUMENT_FLOW.md](../../../../project/docs/DOCUMENT_FLOW.md) — Document lifecycle
