---
trigger: model_decision
description: Pipeline for backend preset. Skill handoffs and phases.
---

# Pipeline (backend)

> Go backend development

## Phases

| Phase | Skills | Outputs |
|-------|--------|---------|
| Discovery | `@feature-fit`, `@idea-interview` | feature-brief, work-unit-registry, discovery-brief, work-unit-registry |
| Definition | `@product-analyst` | roadmap, user-stories, requirements, backlog |
| Architecture | `@bmad-architect`, `@tech-spec-writer` | context-map, api-contracts, decision-log, tech-spec |
| Implementation | `@backend-go-expert` | service-implementation |
| Delivery | `@devops-sre`, `@qa-lead` | deployment-guide, test-cases, test-report |
| Utility | `@debugger`, `@doc-janitor`, `@mcp-expert`, `@project-bro`, `@refactor-architect` | debug-report, known-issues, server-config, refactoring-overview |

## Handoff Matrix

| From | To | Artifact |
|------|-----|----------|
| `@backend-go-expert` | `@qa-lead` | service-implementation |
| `@bmad-architect` | `@tech-spec-writer` | context-map |
| `@bmad-architect` | `@tech-spec-writer` | api-contracts |
| `@bmad-architect` | `@tech-spec-writer` | decision-log |
| `@debugger` | `@qa-lead` | debug-report |
| `@debugger` | `@qa-lead` | known-issues |
| `@feature-fit` | `@product-analyst` | feature-brief |
| `@feature-fit` | `@product-analyst` | work-unit-registry |
| `@idea-interview` | `@product-analyst` | discovery-brief |
| `@idea-interview` | `@product-analyst` | work-unit-registry |
| `@mcp-expert` | `@backend-go-expert` | server-config |
| `@mcp-expert` | `@devops-sre` | server-config |
| `@product-analyst` | `@bmad-architect` | roadmap |
| `@product-analyst` | `@bmad-architect` | user-stories |
| `@product-analyst` | `@bmad-architect` | requirements |
| `@product-analyst` | `@bmad-architect` | backlog |
| `@product-analyst` | `@tech-spec-writer` | roadmap |
| `@product-analyst` | `@tech-spec-writer` | user-stories |
| `@product-analyst` | `@tech-spec-writer` | requirements |
| `@product-analyst` | `@tech-spec-writer` | backlog |
| `@qa-lead` | `@devops-sre` | test-cases |
| `@qa-lead` | `@devops-sre` | test-report |
| `@refactor-architect` | `@backend-go-expert` | refactoring-overview |
| `@refactor-architect` | `@devops-sre` | refactoring-overview |
| `@tech-spec-writer` | `@backend-go-expert` | tech-spec |

## Return Paths

| From | To | Trigger |
|------|-----|---------|
| `@qa-lead` | `@backend-go-expert` | Bugs found in backend go expert |
