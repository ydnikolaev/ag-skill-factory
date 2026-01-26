---
trigger: model_decision
description: Pipeline for fullstack preset. Skill handoffs and phases.
---

# Pipeline (fullstack)

> Full stack development

## Phases

| Phase | Skills | Outputs |
|-------|--------|---------|
| Discovery | `@feature-fit`, `@idea-interview` | feature-brief, discovery-brief |
| Definition | `@product-analyst` | user-stories, requirements |
| Design | `@ui-implementor`, `@ux-designer` | theming, tokens, design-system |
| Architecture | `@bmad-architect`, `@tech-spec-writer` | context-map, api-contracts, tech-spec |
| Implementation | `@backend-go-expert`, `@frontend-nuxt` | service-implementation, ui-implementation |
| Delivery | `@devops-sre`, `@qa-lead` | deployment-guide, test-cases, test-report |
| Utility | `@debugger`, `@doc-janitor`, `@mcp-expert`, `@project-bro`, `@refactor-architect` | debug-report, server-config, refactoring-overview |

## Handoff Matrix

| From | To | Artifact |
|------|-----|----------|
| `@backend-go-expert` | `@qa-lead` | service-implementation |
| `@bmad-architect` | `@tech-spec-writer` | context-map |
| `@bmad-architect` | `@tech-spec-writer` | api-contracts |
| `@debugger` | `@qa-lead` | debug-report |
| `@feature-fit` | `@product-analyst` | feature-brief |
| `@frontend-nuxt` | `@qa-lead` | ui-implementation |
| `@idea-interview` | `@product-analyst` | discovery-brief |
| `@mcp-expert` | `@backend-go-expert` | server-config |
| `@mcp-expert` | `@devops-sre` | server-config |
| `@product-analyst` | `@bmad-architect` | user-stories |
| `@product-analyst` | `@bmad-architect` | requirements |
| `@product-analyst` | `@tech-spec-writer` | user-stories |
| `@product-analyst` | `@tech-spec-writer` | requirements |
| `@qa-lead` | `@devops-sre` | test-cases |
| `@qa-lead` | `@devops-sre` | test-report |
| `@refactor-architect` | `@backend-go-expert` | refactoring-overview |
| `@refactor-architect` | `@frontend-nuxt` | refactoring-overview |
| `@refactor-architect` | `@devops-sre` | refactoring-overview |
| `@tech-spec-writer` | `@backend-go-expert` | tech-spec |
| `@tech-spec-writer` | `@frontend-nuxt` | tech-spec |
| `@ui-implementor` | `@frontend-nuxt` | theming |
| `@ux-designer` | `@ui-implementor` | tokens |
| `@ux-designer` | `@ui-implementor` | design-system |

## Return Paths

| From | To | Trigger |
|------|-----|---------|
| `@qa-lead` | `@backend-go-expert` | Bugs found in backend go expert |
| `@qa-lead` | `@frontend-nuxt` | Bugs found in frontend nuxt |
