---
trigger: model_decision
description: Pipeline for tma preset. Skill handoffs and phases.
---

# Pipeline (tma)

> Telegram Mini Apps (10 skills)

## Phases

| Phase | Skills | Outputs |
|-------|--------|---------|
| Discovery | `@idea-interview` | discovery-brief.md |
| Definition | `@product-analyst` | roadmap.md, user-stories.md, requirements.md |
| Design | `@ui-implementor`, `@ux-designer` | theming.md, tokens.json, design-system.md |
| Architecture | `@bmad-architect`, `@tech-spec-writer`, `@telegram-mechanic` | context-map.md, api-contracts.yaml, tech-spec.md, webhook-config.md |
| Implementation | `@frontend-nuxt`, `@tma-expert` | ui-implementation.md, tma-config.md |
| Delivery | `@qa-lead` | test-cases.md, test-report.md |
| Utility | `@doc-janitor`, `@refactor-architect` | refactoring-overview.md |

## Handoff Matrix

| From | To | Artifact |
|------|-----|----------|
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
| `@refactor-architect` | `@frontend-nuxt` | refactoring-overview.md |
| `@tech-spec-writer` | `@frontend-nuxt` | tech-spec.md |
| `@telegram-mechanic` | `@tma-expert` | webhook-config.md |
| `@tma-expert` | `@qa-lead` | tma-config.md |
| `@ui-implementor` | `@frontend-nuxt` | theming.md |
| `@ux-designer` | `@ui-implementor` | tokens.json |
| `@ux-designer` | `@ui-implementor` | design-system.md |

## Return Paths

| From | To | Trigger |
|------|-----|---------|
| `@qa-lead` | `@frontend-nuxt` | Bugs found in frontend nuxt |
