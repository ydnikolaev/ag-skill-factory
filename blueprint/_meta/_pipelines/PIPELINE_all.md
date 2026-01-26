---
trigger: model_decision
description: Pipeline for all preset. Skill handoffs and phases.
---

# Pipeline (all)

> Full blueprint (21 skills, all extras)

## Phases

| Phase | Skills | Outputs |
|-------|--------|---------|
| Discovery | `@feature-fit`, `@idea-interview` | feature-brief.md, discovery-brief.md |
| Definition | `@product-analyst` | roadmap.md, user-stories.md, requirements.md, backlog.md |
| Design | `@ui-implementor`, `@ux-designer` | theming.md, tokens.json, design-system.md |
| Architecture | `@bmad-architect`, `@cli-architect`, `@tech-spec-writer`, `@telegram-mechanic` | context-map.md, api-contracts.yaml, decision-log.md, cli-design.md, tech-spec.md, webhook-config.md |
| Implementation | `@backend-go-expert`, `@frontend-nuxt`, `@tma-expert`, `@tui-charm-expert` | service-implementation.md, ui-implementation.md, tma-config.md, tui-design.md |
| Delivery | `@devops-sre`, `@qa-lead` | deployment-guide.md, test-cases.md, test-report.md |
| Utility | `@debugger`, `@doc-janitor`, `@mcp-expert`, `@project-bro`, `@refactor-architect` | debug-report.md, known-issues.md, server-config.md, refactoring-overview.md |

## Handoff Matrix

| From | To | Artifact |
|------|-----|----------|
| `@backend-go-expert` | `@qa-lead` | service-implementation.md |
| `@bmad-architect` | `@tech-spec-writer` | context-map.md |
| `@bmad-architect` | `@tech-spec-writer` | api-contracts.yaml |
| `@bmad-architect` | `@tech-spec-writer` | decision-log.md |
| `@cli-architect` | `@tui-charm-expert` | cli-design.md |
| `@cli-architect` | `@backend-go-expert` | cli-design.md |
| `@debugger` | `@qa-lead` | debug-report.md |
| `@debugger` | `@qa-lead` | known-issues.md |
| `@feature-fit` | `@product-analyst` | feature-brief.md |
| `@frontend-nuxt` | `@qa-lead` | ui-implementation.md |
| `@idea-interview` | `@product-analyst` | discovery-brief.md |
| `@mcp-expert` | `@backend-go-expert` | server-config.md |
| `@mcp-expert` | `@devops-sre` | server-config.md |
| `@product-analyst` | `@bmad-architect` | roadmap.md |
| `@product-analyst` | `@bmad-architect` | user-stories.md |
| `@product-analyst` | `@bmad-architect` | requirements.md |
| `@product-analyst` | `@bmad-architect` | backlog.md |
| `@product-analyst` | `@tech-spec-writer` | roadmap.md |
| `@product-analyst` | `@tech-spec-writer` | user-stories.md |
| `@product-analyst` | `@tech-spec-writer` | requirements.md |
| `@product-analyst` | `@tech-spec-writer` | backlog.md |
| `@qa-lead` | `@devops-sre` | test-cases.md |
| `@qa-lead` | `@devops-sre` | test-report.md |
| `@refactor-architect` | `@backend-go-expert` | refactoring-overview.md |
| `@refactor-architect` | `@frontend-nuxt` | refactoring-overview.md |
| `@refactor-architect` | `@devops-sre` | refactoring-overview.md |
| `@tech-spec-writer` | `@backend-go-expert` | tech-spec.md |
| `@tech-spec-writer` | `@frontend-nuxt` | tech-spec.md |
| `@telegram-mechanic` | `@backend-go-expert` | webhook-config.md |
| `@telegram-mechanic` | `@tma-expert` | webhook-config.md |
| `@tma-expert` | `@qa-lead` | tma-config.md |
| `@tui-charm-expert` | `@qa-lead` | tui-design.md |
| `@ui-implementor` | `@frontend-nuxt` | theming.md |
| `@ux-designer` | `@ui-implementor` | tokens.json |
| `@ux-designer` | `@ui-implementor` | design-system.md |

## Return Paths

| From | To | Trigger |
|------|-----|---------|
| `@qa-lead` | `@backend-go-expert` | Bugs found in backend go expert |
| `@qa-lead` | `@frontend-nuxt` | Bugs found in frontend nuxt |
