---
description: Deploy after tests pass
---

# Deploy Workflow

Deploy to production after QA approval.

## Steps

1. Verify `test-report.md` shows all tests passing
2. Activate `@devops-sre` skill
3. Create/update deployment guide in `active/infrastructure/`
4. Execute deployment (CI/CD or manual)
5. If Timeweb hosting → also activate `@timeweb-sysadmin`
6. Verify deployment success

## Trigger

- QA approved
- "Deploy to production"
- "Ship it"
