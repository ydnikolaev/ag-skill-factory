---
status: Draft
owner: @tech-spec-writer
work_unit: {WORK_UNIT}

upstream:
  - doc: user-stories.md
    owner: @product-analyst
  - doc: context-map.md
    owner: @bmad-architect
downstream:
  - doc: service-implementation.md
    owner: @backend-go-expert
  - doc: ui-implementation.md
    owner: @frontend-nuxt

created: {DATE}
updated: {DATE}
---

# Tech Spec: {WORK_UNIT}

## Upstream Documents

| Document | Owner | Status |
|----------|-------|--------|
| [user-stories](../product/{WORK_UNIT}.md) | @product-analyst | ✅ |
| [context-map](../architecture/{WORK_UNIT}.md) | @bmad-architect | ✅ |

---

## Scope

<!-- What is being built in this work unit -->

---

## Technical Decisions

| Decision | Rationale |
|----------|-----------|
| ... | ... |

---

## API Contracts

### Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | /api/v1/... | ... |

---

## Data Model

### New Tables

```sql
-- Add schema changes here
```

### Migrations

- [ ] Migration file created
- [ ] Tested locally

---

## Requirements Checklist

| ID | Requirement | Covered | Notes |
|----|-------------|---------|-------|
| US-01 | ... | ⬜ | |
| US-02 | ... | ⬜ | |

---

## Open Questions

<!-- List any unresolved questions -->
