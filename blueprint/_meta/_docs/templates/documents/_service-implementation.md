---
status: Draft
owner: @backend-go-expert
work_unit: {WORK_UNIT}

upstream:
  - doc: tech-spec.md
    owner: @tech-spec-writer
downstream:
  - doc: test-report.md
    owner: @qa-lead

created: {DATE}
updated: {DATE}
---

# Service Implementation: {WORK_UNIT}

## Upstream Documents

| Document | Owner | Status |
|----------|-------|--------|
| [tech-spec](../specs/{WORK_UNIT}.md) | @tech-spec-writer | ✅ |

---

## Implementation Summary

<!-- Brief overview of what was implemented -->

---

## Components Created

| Package | Files | Purpose |
|---------|-------|---------|
| `internal/domain/...` | ... | Domain entities |
| `internal/app/...` | ... | Use cases |
| `internal/infra/...` | ... | Repositories |

---

## API Endpoints Implemented

| Method | Path | Handler | Status |
|--------|------|---------|--------|
| GET | /api/v1/... | ... | ✅ |

---

## Database Migrations

| Migration | Description | Status |
|-----------|-------------|--------|
| 001_... | Initial schema | ✅ |

---

## Tests

| Type | Coverage | Status |
|------|----------|--------|
| Unit | ...% | ✅ |
| Integration | ...% | ✅ |

---

## Requirements Coverage

| ID | Requirement | Implemented | Notes |
|----|-------------|-------------|-------|
| US-01 | ... | ✅ | |

---

## Known Issues

<!-- List any known issues or tech debt -->
