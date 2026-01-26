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
| Definition | `@product-analyst`, `@ux-designer` | roadmap.md, user-stories.md, requirements.md |
| Architecture | `@bmad-architect`, `@tech-spec-writer` | context-map.md, api-contracts.yaml, tech-spec.md |
| Implementation | `@backend-go-expert`, `@frontend-nuxt`, `@ui-implementor` | service-implementation.md, ui-implementation.md, theming.md |
| Delivery | `@devops-sre`, `@qa-lead` | deployment-guide.md, test-cases.md, test-report.md |
| Utility | `@doc-janitor`, `@project-bro`, `@refactor-architect` | refactoring-overview.md |

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
| `@refactor-architect` | `@backend-go-expert` | refactoring-overview.md |
| `@refactor-architect` | `@frontend-nuxt` | refactoring-overview.md |
| `@refactor-architect` | `@devops-sre` | refactoring-overview.md |
| `@tech-spec-writer` | `@backend-go-expert` | tech-spec.md |
| `@tech-spec-writer` | `@frontend-nuxt` | tech-spec.md |
| `@ui-implementor` | `@frontend-nuxt` | theming.md |
| `@ux-designer` | `@ui-implementor` | tokens.json |
| `@ux-designer` | `@ui-implementor` | design-system.md |

## Return Paths

| From | To | Trigger |
|------|-----|---------|
| `@qa-lead` | `@backend-go-expert` | Bugs found in backend go expert |
| `@qa-lead` | `@frontend-nuxt` | Bugs found in frontend nuxt |
