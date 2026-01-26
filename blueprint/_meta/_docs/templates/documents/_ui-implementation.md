---
status: Draft
owner: @frontend-nuxt
work_unit: {WORK_UNIT}

upstream:
  - doc_type: refactoring-overview
    owner: @refactor-architect
  - doc_type: tech-spec
    owner: @tech-spec-writer
  - doc_type: theming
    owner: @ui-implementor
downstream:
  - skill: @qa-lead

created: {DATE}
updated: {DATE}
---

# Ui Implementation: {WORK_UNIT}

## Pages

| Route | Component | Status |
|-------|-----------|--------|
| `/` | `pages/index.vue` | ⬜ |
| `/dashboard` | `pages/dashboard.vue` | ⬜ |

---

## Components

| Component | Location | Props |
|-----------|----------|-------|
| `Button` | `components/ui/Button.vue` | `variant`, `size` |

---

## API Integration

| Endpoint | Composable | Status |
|----------|------------|--------|
| `GET /api/items` | `useItems()` | ⬜ |

---

## Testing

- [ ] Component tests
- [ ] E2E tests

