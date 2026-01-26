---
trigger: model_decision
description: Pipeline for cli preset. Skill handoffs and phases.
---

# Pipeline (cli)

> CLI/TUI applications

## Phases

| Phase | Skills | Outputs |
|-------|--------|---------|
| Discovery | `@feature-fit`, `@idea-interview` | feature-brief, work-unit-registry, discovery-brief, work-unit-registry |
| Definition | `@product-analyst` | user-stories, requirements |
| Architecture | `@bmad-architect`, `@cli-architect`, `@tech-spec-writer` | context-map, api-contracts, cli-design, tech-spec |
| Implementation | `@tui-charm-expert` | tui-design |
| Delivery | `@qa-lead` | test-cases, test-report |
| Utility | `@doc-janitor`, `@mcp-expert`, `@refactor-architect` | server-config, refactoring-overview |

## Handoff Matrix

| From | To | Artifact |
|------|-----|----------|
| `@bmad-architect` | `@tech-spec-writer` | context-map |
| `@bmad-architect` | `@tech-spec-writer` | api-contracts |
| `@cli-architect` | `@tui-charm-expert` | cli-design |
| `@feature-fit` | `@product-analyst` | feature-brief |
| `@feature-fit` | `@product-analyst` | work-unit-registry |
| `@idea-interview` | `@product-analyst` | discovery-brief |
| `@idea-interview` | `@product-analyst` | work-unit-registry |
| `@product-analyst` | `@bmad-architect` | user-stories |
| `@product-analyst` | `@bmad-architect` | requirements |
| `@product-analyst` | `@tech-spec-writer` | user-stories |
| `@product-analyst` | `@tech-spec-writer` | requirements |
| `@tui-charm-expert` | `@qa-lead` | tui-design |

## Return Paths

| From | To | Trigger |
|------|-----|---------|
| `@qa-lead` | `@tui-charm-expert` | Bugs found in tui charm expert |
