---
trigger: model_decision
description: Pipeline for core preset. Skill handoffs and phases.
---

# Pipeline (core)

> Pipeline essentials

## Phases

| Phase | Skills | Outputs |
|-------|--------|---------|
| Discovery | `@feature-fit`, `@idea-interview` | feature-brief, work-unit-registry, discovery-brief, work-unit-registry |
| Definition | `@product-analyst` | roadmap, user-stories, requirements, backlog |
| Architecture | `@bmad-architect`, `@tech-spec-writer` | context-map, api-contracts, decision-log, tech-spec |
| Delivery | `@qa-lead` | test-cases, test-report |
| Utility | `@doc-janitor`, `@mcp-expert`, `@refactor-architect` | server-config, refactoring-overview |

## Handoff Matrix

| From | To | Artifact |
|------|-----|----------|
| `@bmad-architect` | `@tech-spec-writer` | context-map |
| `@bmad-architect` | `@tech-spec-writer` | api-contracts |
| `@bmad-architect` | `@tech-spec-writer` | decision-log |
| `@feature-fit` | `@product-analyst` | feature-brief |
| `@feature-fit` | `@product-analyst` | work-unit-registry |
| `@idea-interview` | `@product-analyst` | discovery-brief |
| `@idea-interview` | `@product-analyst` | work-unit-registry |
| `@product-analyst` | `@bmad-architect` | roadmap |
| `@product-analyst` | `@bmad-architect` | user-stories |
| `@product-analyst` | `@bmad-architect` | requirements |
| `@product-analyst` | `@bmad-architect` | backlog |
| `@product-analyst` | `@tech-spec-writer` | roadmap |
| `@product-analyst` | `@tech-spec-writer` | user-stories |
| `@product-analyst` | `@tech-spec-writer` | requirements |
| `@product-analyst` | `@tech-spec-writer` | backlog |
