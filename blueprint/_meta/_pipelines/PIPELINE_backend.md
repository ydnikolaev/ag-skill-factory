---
trigger: model_decision
description: Pipeline for backend preset. Skill handoffs and phases.
---

# Pipeline (backend)

> Go backend development (9 skills)

## Phases

| Phase | Skills | Outputs |
|-------|--------|---------|
| Discovery | `@idea-interview` | discovery-brief.md |
| Definition | `@product-analyst` | roadmap.md, user-stories.md, requirements.md |
| Architecture | `@bmad-architect`, `@tech-spec-writer` | context-map.md, api-contracts.yaml, tech-spec.md |
| Implementation | `@backend-go-expert` | service-implementation.md |
| Delivery | `@devops-sre`, `@qa-lead` | deployment-guide.md, test-cases.md, test-report.md |
| Utility | `@debugger`, `@doc-janitor`, `@project-bro`, `@refactor-architect` | debug-report.md, refactoring-overview.md |

## Handoff Matrix

| From | To | Artifact |
|------|-----|----------|
| `@backend-go-expert` | `@qa-lead` | service-implementation.md |
| `@bmad-architect` | `@tech-spec-writer` | context-map.md |
| `@bmad-architect` | `@tech-spec-writer` | api-contracts.yaml |
| `@debugger` | `@qa-lead` | debug-report.md |
| `@idea-interview` | `@product-analyst` | discovery-brief.md |
| `@product-analyst` | `@bmad-architect` | roadmap.md |
| `@product-analyst` | `@bmad-architect` | user-stories.md |
| `@product-analyst` | `@bmad-architect` | requirements.md |
| `@product-analyst` | `@tech-spec-writer` | roadmap.md |
| `@product-analyst` | `@tech-spec-writer` | user-stories.md |
| `@product-analyst` | `@tech-spec-writer` | requirements.md |
| `@qa-lead` | `@devops-sre` | test-cases.md |
| `@qa-lead` | `@devops-sre` | test-report.md |
| `@refactor-architect` | `@backend-go-expert` | refactoring-overview.md |
| `@refactor-architect` | `@devops-sre` | refactoring-overview.md |
| `@tech-spec-writer` | `@backend-go-expert` | tech-spec.md |

## Return Paths

| From | To | Trigger |
|------|-----|---------|
| `@qa-lead` | `@backend-go-expert` | Bugs found in backend go expert |
