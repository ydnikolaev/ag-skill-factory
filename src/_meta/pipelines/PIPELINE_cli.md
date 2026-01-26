---
trigger: model_decision
description: Pipeline for cli preset. Skill handoffs and phases.
---

# Pipeline (cli)

> CLI/TUI applications

## Phases

| Phase | Skills | Outputs |
|-------|--------|---------|
| Discovery | `@feature-fit`, `@idea-interview` | feature-brief, discovery-brief |
| Definition | `@product-analyst` | roadmap, user-stories, requirements, backlog |
| Architecture | `@bmad-architect`, `@cli-architect`, `@tech-spec-writer` | context-map, api-contracts, decision-log, cli-design, tech-spec |
| Implementation | `@tui-charm-expert` | tui-design |
| Delivery | `@qa-lead` | test-cases, test-report |
| Utility | `@doc-janitor`, `@mcp-expert`, `@refactor-architect` | server-config, refactoring-overview |

## Handoff Matrix

| From | To | Artifact |
|------|-----|----------|
| `@bmad-architect` | `@tech-spec-writer` | context-map |
| `@bmad-architect` | `@tech-spec-writer` | api-contracts |
| `@bmad-architect` | `@tech-spec-writer` | decision-log |
| `@cli-architect` | `@tui-charm-expert` | cli-design |
| `@feature-fit` | `@product-analyst` | feature-brief |
| `@idea-interview` | `@product-analyst` | discovery-brief |
| `@product-analyst` | `@bmad-architect` | roadmap |
| `@product-analyst` | `@bmad-architect` | user-stories |
| `@product-analyst` | `@bmad-architect` | requirements |
| `@product-analyst` | `@bmad-architect` | backlog |
| `@product-analyst` | `@tech-spec-writer` | roadmap |
| `@product-analyst` | `@tech-spec-writer` | user-stories |
| `@product-analyst` | `@tech-spec-writer` | requirements |
| `@product-analyst` | `@tech-spec-writer` | backlog |
| `@tui-charm-expert` | `@qa-lead` | tui-design |
