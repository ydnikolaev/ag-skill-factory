---
status: Draft
owner: @mcp-expert
lifecycle: per-feature
work_unit: {WORK_UNIT}

downstream:
  - skill: @backend-go-expert
  - skill: @devops-sre

created: {DATE}
updated: {DATE}
---

# Server Config: {WORK_UNIT}

## Server Definition

```yaml
servers:
  my-server:
    type: stdio
    command: go run ./cmd/mcp
```

---

## Tools

| Tool | Description | Parameters |
|------|-------------|------------|
| `db_query` | Execute SQL query | `query: string` |

---

## Resources

| URI | Description |
|-----|-------------|
| `db://schema` | Database schema |

