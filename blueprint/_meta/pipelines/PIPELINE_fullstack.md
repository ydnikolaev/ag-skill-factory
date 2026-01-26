---
trigger: model_decision
description: Pipeline for fullstack preset. Skill handoffs and phases.
---

# Pipeline (fullstack)

> Full stack development (12 skills)

## Phases

| Phase | Skills | Outputs |
|-------|--------|---------|
| Discovery | `@idea-interview` | discovery-brief.md |
| Definition | `@product-analyst` | roadmap.md, user-stories.md, requirements.md |
| Architecture | `@bmad-architect`, `@tech-spec-writer` | context-map.md, api-contracts.yaml, tech-spec.md |
| Implementation | `@backend-go-expert`, `@frontend-nuxt` | service-implementation.md, ui-implementation.md |
| Delivery | `@devops-sre`, `@qa-lead` | deployment-guide.md, test-cases.md, test-report.md |
| Utility | `@doc-janitor`, `@project-bro`, `@refactor-architect`, `@ui-implementor`, `@ux-designer` | — |

## Handoff Matrix

| From | To | Artifact |
|------|-----|----------|
| `@backend-go-expert` | `@qa-lead` | service-implementation.md |
| `@bmad-architect` | `@tech-spec-writer` | context-map.md |
| `@bmad-architect` | `@tech-spec-writer` | api-contracts.yaml |
| `@frontend-nuxt` | `@qa-lead` | ui-implementation.md |
| `@idea-interview` | `@product-analyst` | discovery-brief.md |
| `@product-analyst` | `@bmad-architect` | roadmap.md |
| `@product-analyst` | `@bmad-architect` | user-stories.md |
| `@product-analyst` | `@bmad-architect` | requirements.md |
| `@product-analyst` | `@tech-spec-writer` | roadmap.md |
| `@product-analyst` | `@tech-spec-writer` | user-stories.md |
| `@product-analyst` | `@tech-spec-writer` | requirements.md |
| `@qa-lead` | `@devops-sre` | test-cases.md |
| `@qa-lead` | `@devops-sre` | test-report.md |
| `@tech-spec-writer` | `@backend-go-expert` | tech-spec.md |
| `@tech-spec-writer` | `@frontend-nuxt` | tech-spec.md |

## Return Paths

| From | To | Trigger |
|------|-----|---------|
| `@qa-lead` | `@backend-go-expert` | Bugs found in backend go expert |
| `@qa-lead` | `@frontend-nuxt` | Bugs found in frontend nuxt |
