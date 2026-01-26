---
trigger: model_decision
description: Pipeline for core preset. Skill handoffs and phases.
---

# Pipeline (core)

> Pipeline essentials

## Phases

| Phase | Skills | Outputs |
|-------|--------|---------|
| Discovery | `@feature-fit`, `@idea-interview` | feature-brief, discovery-brief |
| Definition | `@product-analyst` | user-stories, requirements |
| Architecture | `@bmad-architect`, `@tech-spec-writer` | context-map, api-contracts, tech-spec |
| Delivery | `@qa-lead` | test-cases, test-report |
| Utility | `@doc-janitor`, `@mcp-expert`, `@refactor-architect` | server-config, refactoring-overview |

## Handoff Matrix

| From | To | Artifact |
|------|-----|----------|
| `@bmad-architect` | `@tech-spec-writer` | context-map |
| `@bmad-architect` | `@tech-spec-writer` | api-contracts |
| `@feature-fit` | `@product-analyst` | feature-brief |
| `@idea-interview` | `@product-analyst` | discovery-brief |
| `@product-analyst` | `@bmad-architect` | user-stories |
| `@product-analyst` | `@bmad-architect` | requirements |
| `@product-analyst` | `@tech-spec-writer` | user-stories |
| `@product-analyst` | `@tech-spec-writer` | requirements |
