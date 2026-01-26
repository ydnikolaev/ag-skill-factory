---
trigger: model_decision
description: Pipeline for cli preset. Skill handoffs and phases.
---

# Pipeline (cli)

> CLI/TUI applications (8 skills)

## Phases

| Phase | Skills | Outputs |
|-------|--------|---------|
| Discovery | `@idea-interview` | discovery-brief.md |
| Definition | `@product-analyst` | roadmap.md, user-stories.md, requirements.md |
| Architecture | `@bmad-architect`, `@cli-architect`, `@tech-spec-writer` | context-map.md, api-contracts.yaml, cli-design.md, tech-spec.md |
| Implementation | `@backend-go-expert`, `@tui-charm-expert` | service-implementation.md, tui-design.md |
| Delivery | `@qa-lead` | test-cases.md, test-report.md |
| Utility | `@doc-janitor`, `@refactor-architect` | refactoring-overview.md |

## Handoff Matrix

| From | To | Artifact |
|------|-----|----------|
| `@backend-go-expert` | `@qa-lead` | service-implementation.md |
| `@bmad-architect` | `@tech-spec-writer` | context-map.md |
| `@bmad-architect` | `@tech-spec-writer` | api-contracts.yaml |
| `@cli-architect` | `@tui-charm-expert` | cli-design.md |
| `@cli-architect` | `@backend-go-expert` | cli-design.md |
| `@idea-interview` | `@product-analyst` | discovery-brief.md |
| `@product-analyst` | `@bmad-architect` | roadmap.md |
| `@product-analyst` | `@bmad-architect` | user-stories.md |
| `@product-analyst` | `@bmad-architect` | requirements.md |
| `@product-analyst` | `@tech-spec-writer` | roadmap.md |
| `@product-analyst` | `@tech-spec-writer` | user-stories.md |
| `@product-analyst` | `@tech-spec-writer` | requirements.md |
| `@refactor-architect` | `@backend-go-expert` | refactoring-overview.md |
| `@tech-spec-writer` | `@backend-go-expert` | tech-spec.md |
| `@tui-charm-expert` | `@qa-lead` | tui-design.md |

## Return Paths

| From | To | Trigger |
|------|-----|---------|
| `@qa-lead` | `@backend-go-expert` | Bugs found in backend go expert |
