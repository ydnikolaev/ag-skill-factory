---
status: Draft
owner: @qa-lead
lifecycle: per-feature
work_unit: {WORK_UNIT}

upstream:
  - doc_type: service-implementation
    owner: @backend-go-expert
  - doc_type: debug-report
    owner: @debugger
  - doc_type: known-issues
    owner: @debugger
  - doc_type: ui-implementation
    owner: @frontend-nuxt
  - doc_type: tma-config
    owner: @tma-expert
  - doc_type: tui-design
    owner: @tui-charm-expert
downstream:
  - skill: @devops-sre

created: {DATE}
updated: {DATE}
---

# Test Cases: {WORK_UNIT}

## Test Summary

| Type | Count | Pass | Fail |
|------|-------|------|------|
| Unit | ... | ... | ... |
| Integration | ... | ... | ... |
| E2E | ... | ... | ... |

---

## Test Cases

### TC-01: ...

| Field | Value |
|-------|-------|
| **Preconditions** | ... |
| **Steps** | 1. ... 2. ... |
| **Expected** | ... |
| **Status** | ⬜ |

