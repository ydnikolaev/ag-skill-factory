---
status: Draft
owner: @devops-sre
work_unit: {WORK_UNIT}

upstream:
  - doc_type: server-config
    owner: @mcp-expert
  - doc_type: test-cases
    owner: @qa-lead
  - doc_type: test-report
    owner: @qa-lead
  - doc_type: refactoring-overview
    owner: @refactor-architect

created: {DATE}
updated: {DATE}
---

# Deployment Guide: {WORK_UNIT}

## Prerequisites

- [ ] Docker installed
- [ ] Access to registry
- [ ] Environment variables configured

---

## Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `DATABASE_URL` | Postgres connection string | ✅ |
| `SECRET_KEY` | Application secret | ✅ |

---

## Build

```bash
docker build -t app:latest .
```

---

## Deploy

```bash
docker compose -f docker-compose.prod.yml up -d
```

---

## Health Checks

| Endpoint | Expected |
|----------|----------|
| `/health` | 200 OK |

