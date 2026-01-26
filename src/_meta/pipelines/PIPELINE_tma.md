---
trigger: model_decision
description: Pipeline for tma preset. Skill handoffs and phases.
---

# Pipeline (tma)

> Telegram Mini Apps

## Phases

| Phase | Skills | Outputs |
|-------|--------|---------|
| Discovery | `@feature-fit`, `@idea-interview` | feature-brief, discovery-brief |
| Definition | `@product-analyst` | user-stories, requirements |
| Design | `@ui-implementor`, `@ux-designer` | theming, tokens, design-system |
| Architecture | `@bmad-architect`, `@tech-spec-writer`, `@telegram-mechanic` | context-map, api-contracts, tech-spec, webhook-config |
| Implementation | `@frontend-nuxt`, `@tma-expert` | ui-implementation, tma-config |
| Delivery | `@qa-lead` | test-cases, test-report |
| Utility | `@doc-janitor`, `@mcp-expert`, `@refactor-architect` | server-config, refactoring-overview |

## Handoff Matrix

| From | To | Artifact |
|------|-----|----------|
| `@bmad-architect` | `@tech-spec-writer` | context-map |
| `@bmad-architect` | `@tech-spec-writer` | api-contracts |
| `@feature-fit` | `@product-analyst` | feature-brief |
| `@frontend-nuxt` | `@qa-lead` | ui-implementation |
| `@idea-interview` | `@product-analyst` | discovery-brief |
| `@product-analyst` | `@bmad-architect` | user-stories |
| `@product-analyst` | `@bmad-architect` | requirements |
| `@product-analyst` | `@tech-spec-writer` | user-stories |
| `@product-analyst` | `@tech-spec-writer` | requirements |
| `@refactor-architect` | `@frontend-nuxt` | refactoring-overview |
| `@tech-spec-writer` | `@frontend-nuxt` | tech-spec |
| `@telegram-mechanic` | `@tma-expert` | webhook-config |
| `@tma-expert` | `@qa-lead` | tma-config |
| `@ui-implementor` | `@frontend-nuxt` | theming |
| `@ux-designer` | `@ui-implementor` | tokens |
| `@ux-designer` | `@ui-implementor` | design-system |

## Return Paths

| From | To | Trigger |
|------|-----|---------|
| `@qa-lead` | `@frontend-nuxt` | Bugs found in frontend nuxt |
| `@qa-lead` | `@tma-expert` | Bugs found in tma expert |
