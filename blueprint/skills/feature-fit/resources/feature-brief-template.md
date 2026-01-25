# Feature Brief Template

Use this template when creating Feature Briefs.

---

```markdown
---
status: Draft
author: feature-fit
date: {{DATE}}
---

# Feature: [Feature Name]

## Context

**Project Stack** (from CONFIG.yaml):
- Backend: [e.g., Go 1.25, PostgreSQL]
- Frontend: [e.g., Nuxt 4, TailwindCSS]
- Integrations: [e.g., Telegram Bot, MCP servers]

**Where it fits**:
- Bounded Context: [e.g., User Management]
- Existing Components: [e.g., Uses User table]

## Requirements

**User Goal**: [What problem does this solve?]

**Key Features**:
1. [Feature 1]
2. [Feature 2]
3. [Feature 3]

**Scope**:
- ✅ MVP: [What's included]
- ❌ V2: [What's deferred]

## Gap Analysis

### Backend
- [ ] [Change 1]
- [ ] [Change 2]

### Frontend
- [ ] [Change 1]
- [ ] [Change 2]

### MCP / Integrations
- [ ] [Change 1 or "None needed"]

## Impact / Risks

| Area | Impact | Risk Level |
|------|--------|------------|
| Database | [Tables affected] | 🟡 Medium |
| API | [Breaking changes?] | 🔴 High / 🟢 Low |
| Security | [New permissions?] | 🟢 Low |

## TDD Strategy

- **Unit Tests**: [What to mock?]
- **Integration Tests**: [Which endpoints?]
- **E2E Tests**: [User flow to validate?]

## Next Steps

Delegate to `@product-analyst` for User Stories and Specs.
```
