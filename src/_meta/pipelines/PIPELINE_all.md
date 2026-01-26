---
trigger: model_decision
description: Pipeline for all preset. Skill handoffs and phases.
---

# Pipeline (all)

> Full blueprint

## Phases

| Phase | Skills | Outputs |
|-------|--------|---------|
| Discovery | `@feature-fit`, `@idea-interview` | feature-brief, work-unit-registry, discovery-brief, work-unit-registry |
| Definition | `@product-analyst` | roadmap, user-stories, requirements, backlog |
| Design | `@ui-implementor`, `@ux-designer` | theming, tokens, design-system |
| Architecture | `@bmad-architect`, `@cli-architect`, `@tech-spec-writer`, `@telegram-mechanic` | context-map, api-contracts, decision-log, cli-design, tech-spec, webhook-config |
| Implementation | `@backend-go-expert`, `@frontend-nuxt`, `@tma-expert`, `@tui-charm-expert` | service-implementation, ui-implementation, tma-config, tui-design |
| Delivery | `@devops-sre`, `@qa-lead` | deployment-guide, test-cases, test-report |
| Utility | `@debugger`, `@doc-janitor`, `@mcp-expert`, `@project-bro`, `@refactor-architect` | debug-report, known-issues, server-config, refactoring-overview |

## Handoff Matrix

| From | To | Artifact |
|------|-----|----------|
| `@backend-go-expert` | `@qa-lead` | service-implementation |
| `@bmad-architect` | `@tech-spec-writer` | context-map |
| `@bmad-architect` | `@tech-spec-writer` | api-contracts |
| `@bmad-architect` | `@tech-spec-writer` | decision-log |
| `@cli-architect` | `@tui-charm-expert` | cli-design |
| `@cli-architect` | `@backend-go-expert` | cli-design |
| `@debugger` | `@qa-lead` | debug-report |
| `@debugger` | `@qa-lead` | known-issues |
| `@feature-fit` | `@product-analyst` | feature-brief |
| `@feature-fit` | `@product-analyst` | work-unit-registry |
| `@frontend-nuxt` | `@qa-lead` | ui-implementation |
| `@idea-interview` | `@product-analyst` | discovery-brief |
| `@idea-interview` | `@product-analyst` | work-unit-registry |
| `@mcp-expert` | `@backend-go-expert` | server-config |
| `@mcp-expert` | `@devops-sre` | server-config |
| `@product-analyst` | `@bmad-architect` | roadmap |
| `@product-analyst` | `@bmad-architect` | user-stories |
| `@product-analyst` | `@bmad-architect` | requirements |
| `@product-analyst` | `@bmad-architect` | backlog |
| `@product-analyst` | `@tech-spec-writer` | roadmap |
| `@product-analyst` | `@tech-spec-writer` | user-stories |
| `@product-analyst` | `@tech-spec-writer` | requirements |
| `@product-analyst` | `@tech-spec-writer` | backlog |
| `@qa-lead` | `@devops-sre` | test-cases |
| `@qa-lead` | `@devops-sre` | test-report |
| `@refactor-architect` | `@backend-go-expert` | refactoring-overview |
| `@refactor-architect` | `@frontend-nuxt` | refactoring-overview |
| `@refactor-architect` | `@devops-sre` | refactoring-overview |
| `@tech-spec-writer` | `@backend-go-expert` | tech-spec |
| `@tech-spec-writer` | `@frontend-nuxt` | tech-spec |
| `@telegram-mechanic` | `@backend-go-expert` | webhook-config |
| `@telegram-mechanic` | `@tma-expert` | webhook-config |
| `@tma-expert` | `@qa-lead` | tma-config |
| `@tui-charm-expert` | `@qa-lead` | tui-design |
| `@ui-implementor` | `@frontend-nuxt` | theming |
| `@ux-designer` | `@ui-implementor` | tokens |
| `@ux-designer` | `@ui-implementor` | design-system |

## Return Paths

| From | To | Trigger |
|------|-----|---------|
| `@qa-lead` | `@backend-go-expert` | Bugs found in backend go expert |
| `@qa-lead` | `@frontend-nuxt` | Bugs found in frontend nuxt |
| `@qa-lead` | `@tma-expert` | Bugs found in tma expert |
| `@qa-lead` | `@tui-charm-expert` | Bugs found in tui charm expert |
