---
status: Draft
owner: @qa-lead
work_unit: {WORK_UNIT}

upstream:
  - doc: service-implementation.md
    owner: @backend-go-expert
  - doc: ui-implementation.md
    owner: @frontend-nuxt
downstream:
  - doc: deployment-guide.md
    owner: @devops-sre

created: {DATE}
updated: {DATE}
---

# Test Report: {WORK_UNIT}

## Upstream Documents

| Document | Owner | Status |
|----------|-------|--------|
| [backend](../backend/{WORK_UNIT}.md) | @backend-go-expert | ✅ |
| [frontend](../frontend/{WORK_UNIT}.md) | @frontend-nuxt | ✅ |

---

## Test Summary

| Type | Total | Passed | Failed | Skipped |
|------|-------|--------|--------|---------|
| Unit | | | | |
| Integration | | | | |
| E2E | | | | |

---

## Test Cases

### TC-01: {Test Case Name}

| Attribute | Value |
|-----------|-------|
| **Preconditions** | ... |
| **Steps** | 1. ... 2. ... |
| **Expected** | ... |
| **Result** | ✅ Pass / ❌ Fail |

---

## Coverage Report

| Module | Coverage |
|--------|----------|
| backend | ...% |
| frontend | ...% |

---

## Issues Found

| ID | Description | Severity | Status |
|----|-------------|----------|--------|
| BUG-01 | ... | High | Fixed |

---

## Recommendations

<!-- Any recommendations for improvement -->

---

## Sign-off

- [ ] All critical tests pass
- [ ] No high-severity bugs open
- [ ] Ready for deployment
