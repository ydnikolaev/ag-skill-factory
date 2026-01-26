---
status: Draft
owner: @bmad-architect
work_unit: {WORK_UNIT}

upstream:
  - doc_type: roadmap
    owner: @product-analyst
  - doc_type: user-stories
    owner: @product-analyst
  - doc_type: requirements
    owner: @product-analyst
  - doc_type: backlog
    owner: @product-analyst
downstream:
  - skill: @tech-spec-writer

created: {DATE}
updated: {DATE}
---

# Api Contracts: {WORK_UNIT}

## Endpoints

| Method | Path | Description | Auth |
|--------|------|-------------|------|
| GET | /api/v1/... | ... | Bearer |

---

## Request/Response Schemas

### GET /api/v1/example

**Request:** `{}`

**Response:**
```json
{"id": "string", "data": {}}
```

---

## Error Codes

| Code | Message | When |
|------|---------|------|
| 400 | Bad Request | Invalid input |
| 401 | Unauthorized | Missing/invalid token |
| 404 | Not Found | Resource doesn't exist |

