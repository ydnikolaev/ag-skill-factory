---
status: Draft
owner: @product-analyst
work_unit: {WORK_UNIT}

upstream:
  - doc: discovery-brief.md
    owner: @idea-interview
downstream:
  - doc: context-map.md
    owner: @bmad-architect
  - doc: tech-spec.md
    owner: @tech-spec-writer

created: {DATE}
updated: {DATE}
---

# User Stories: {WORK_UNIT}

## Upstream Documents

| Document | Owner | Status |
|----------|-------|--------|
| [discovery-brief](../discovery/{WORK_UNIT}.md) | @idea-interview | ✅ |

---

## Epic: {EPIC_NAME}

### US-01: {Story Title}

**As a** {user type}
**I want to** {action}
**So that** {benefit}

#### Acceptance Criteria

- [ ] Given ... When ... Then ...
- [ ] Given ... When ... Then ...

#### Priority: High | Medium | Low

---

### US-02: {Story Title}

**As a** {user type}
**I want to** {action}
**So that** {benefit}

#### Acceptance Criteria

- [ ] Given ... When ... Then ...

#### Priority: High | Medium | Low

---

## Non-Functional Requirements

| ID | Requirement | Priority |
|----|-------------|----------|
| NFR-01 | Performance: ... | High |
| NFR-02 | Security: ... | High |

---

## Story Map

| Release | Stories |
|---------|---------|
| MVP | US-01, US-02 |
| v1.1 | US-03 |
