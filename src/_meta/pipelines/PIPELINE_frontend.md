---
trigger: model_decision
description: Pipeline for frontend preset. Skill handoffs and phases.
---

# Pipeline (frontend)

> Nuxt/Vue frontend

## Phases

| Phase | Skills | Outputs |
|-------|--------|---------|
| Discovery | `@feature-fit`, `@idea-interview` | feature-brief, discovery-brief |
| Definition | `@product-analyst` | user-stories, requirements |
| Design | `@ui-implementor`, `@ux-designer` | theming, tokens, design-system |
| Architecture | `@bmad-architect`, `@tech-spec-writer` | context-map, api-contracts, tech-spec |
| Implementation | `@frontend-nuxt` | ui-implementation |
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
| `@ui-implementor` | `@frontend-nuxt` | theming |
| `@ux-designer` | `@ui-implementor` | tokens |
| `@ux-designer` | `@ui-implementor` | design-system |

## Return Paths

| From | To | Trigger |
|------|-----|---------|
| `@qa-lead` | `@frontend-nuxt` | Bugs found in frontend nuxt |
