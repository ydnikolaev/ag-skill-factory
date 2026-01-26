---
trigger: model_decision
description: Pipeline for core preset. Skill handoffs and phases.
---

# Pipeline (core)

> Pipeline essentials (7 skills)

## Phases

| Phase | Skills | Outputs |
|-------|--------|---------|
| Discovery | `@idea-interview` | discovery-brief.md |
| Definition | `@product-analyst` | roadmap.md, user-stories.md, requirements.md |
| Architecture | `@bmad-architect`, `@tech-spec-writer` | context-map.md, api-contracts.yaml, tech-spec.md |
| Delivery | `@qa-lead` | test-cases.md, test-report.md |
| Utility | `@doc-janitor`, `@refactor-architect` | refactoring-overview.md |

## Handoff Matrix

| From | To | Artifact |
|------|-----|----------|
| `@bmad-architect` | `@tech-spec-writer` | context-map.md |
| `@bmad-architect` | `@tech-spec-writer` | api-contracts.yaml |
| `@idea-interview` | `@product-analyst` | discovery-brief.md |
| `@product-analyst` | `@bmad-architect` | roadmap.md |
| `@product-analyst` | `@bmad-architect` | user-stories.md |
| `@product-analyst` | `@bmad-architect` | requirements.md |
| `@product-analyst` | `@tech-spec-writer` | roadmap.md |
| `@product-analyst` | `@tech-spec-writer` | user-stories.md |
| `@product-analyst` | `@tech-spec-writer` | requirements.md |
