---
status: Draft
owner: @bmad-architect
work_unit: {WORK_UNIT}

upstream:
  - doc: user-stories.md
    owner: @product-analyst
downstream:
  - doc: tech-spec.md
    owner: @tech-spec-writer
  - doc: api-contracts.yaml
    owner: @bmad-architect

created: {DATE}
updated: {DATE}
---

# Context Map: {WORK_UNIT}

## Upstream Documents

| Document | Owner | Status |
|----------|-------|--------|
| [user-stories](../product/{WORK_UNIT}.md) | @product-analyst | ✅ |

---

## System Overview

```mermaid
graph TD
    subgraph "Bounded Contexts"
        A[Context A]
        B[Context B]
    end
    A --> B
```

---

## Bounded Contexts

### Context: {Name}

| Aspect | Description |
|--------|-------------|
| **Responsibility** | ... |
| **Owner** | @skill |
| **Key Entities** | ... |

---

## Context Relationships

| From | To | Relationship | Notes |
|------|-----|--------------|-------|
| Context A | Context B | Upstream-Downstream | ... |

---

## Aggregates

### {Aggregate Name}

- **Root Entity:** ...
- **Value Objects:** ...
- **Invariants:** ...

---

## Domain Events

| Event | Trigger | Consumers |
|-------|---------|-----------|
| UserCreated | User signs up | Notifications, Analytics |

---

## Integration Points

| System | Protocol | Purpose |
|--------|----------|---------|
| Telegram | Webhook | Bot commands |
| PostgreSQL | TCP | Persistence |
