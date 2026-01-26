# Document Flow Audit

> Complete document lifecycle map: who creates → who reads → who archives

---

## Document Types Summary

| Metric | Value |
|--------|-------|
| Total doc types | 27 |
| Per-feature docs | 22 |
| Living docs | 4 |
| Meta docs (registries) | 2 |
| Skills with outputs | 16 |
| Handoffs | 37 |
| Return paths | 4 |

---

## Per-Feature Documents (22 types)

| Doc Type | Creator | Phase | Consumers | Archival |
|----------|---------|-------|-----------|----------|
| discovery-brief | @idea-interview | Discovery | @product-analyst | @doc-janitor |
| feature-brief | @feature-fit | Discovery | @product-analyst | @doc-janitor |
| work-unit-registry | @idea-interview, @feature-fit | Discovery | All skills (update) | @doc-janitor |
| user-stories | @product-analyst | Definition | @bmad-architect, @tech-spec-writer | @doc-janitor |
| requirements | @product-analyst | Definition | @bmad-architect, @tech-spec-writer | @doc-janitor |
| tokens | @ux-designer | Design | @ui-implementor | @doc-janitor |
| design-system | @ux-designer | Design | @ui-implementor | @doc-janitor |
| theming | @ui-implementor | Design | @frontend-nuxt | @doc-janitor |
| context-map | @bmad-architect | Architecture | @tech-spec-writer | @doc-janitor |
| api-contracts | @bmad-architect | Architecture | @tech-spec-writer | @doc-janitor |
| cli-design | @cli-architect | Architecture | @tui-charm-expert, @backend-go-expert | @doc-janitor |
| webhook-config | @telegram-mechanic | Architecture | @backend-go-expert, @tma-expert | @doc-janitor |
| server-config | @mcp-expert | Architecture | @backend-go-expert, @devops-sre | @doc-janitor |
| tech-spec | @tech-spec-writer | Architecture | @backend-go-expert, @frontend-nuxt | @doc-janitor |
| service-implementation | @backend-go-expert | Implementation | @qa-lead | @doc-janitor |
| ui-implementation | @frontend-nuxt | Implementation | @qa-lead | @doc-janitor |
| tma-config | @tma-expert | Implementation | @qa-lead | @doc-janitor |
| tui-design | @tui-charm-expert | Implementation | @qa-lead | @doc-janitor |
| test-cases | @qa-lead | Delivery | @devops-sre | @doc-janitor |
| test-report | @qa-lead | Delivery | @devops-sre | @doc-janitor |
| deployment-guide | @devops-sre | Delivery | — (end of chain) | @doc-janitor |
| debug-report | @debugger | Utility | @qa-lead | @doc-janitor |
| refactoring-overview | @refactor-architect | Utility | @backend-go-expert, @frontend-nuxt, @devops-sre | @doc-janitor |

---

## Living Documents (4 types)

| Doc Type | Creator | Updaters | Location |
|----------|---------|----------|----------|
| roadmap | @product-analyst | @product-analyst | `project/docs/roadmap.md` |
| backlog | @product-analyst | All skills (add items) | `project/docs/backlog.md` |
| decision-log | @bmad-architect | @bmad-architect, @tech-spec-writer | `project/docs/decision-log.md` |
| known-issues | @debugger | @debugger, @qa-lead | `project/docs/known-issues.md` |

---

## Meta Documents (2 types)

| Doc Type | Creator | Updaters | Location |
|----------|---------|----------|----------|
| artifact-registry | @idea-interview, @feature-fit | All skills (on handoff) | `project/docs/ARTIFACT_REGISTRY.md` |
| work-unit-registry | @idea-interview, @feature-fit | All skills (on doc create) | `project/docs/registry/{work-unit}.md` |

---

## Pipeline Flow

```
DISCOVERY (Entry Points)
├── @idea-interview → discovery-brief, work-unit-registry
├── @feature-fit → feature-brief, work-unit-registry
└── Handoff → @product-analyst

DEFINITION
├── @product-analyst → user-stories, requirements
├── Updates: roadmap, backlog (living)
└── Handoff → @bmad-architect, @tech-spec-writer

DESIGN
├── @ux-designer → tokens, design-system
├── @ui-implementor → theming
└── Handoff → @frontend-nuxt

ARCHITECTURE
├── @bmad-architect → context-map, api-contracts, decision-log
├── @cli-architect → cli-design
├── @telegram-mechanic → webhook-config
├── @tech-spec-writer → tech-spec
└── Handoff → Implementation skills

IMPLEMENTATION
├── @backend-go-expert → service-implementation
├── @frontend-nuxt → ui-implementation
├── @tma-expert → tma-config
├── @tui-charm-expert → tui-design
└── Handoff → @qa-lead

DELIVERY
├── @qa-lead → test-cases, test-report
├── Return paths → @backend-go-expert, @frontend-nuxt, @tma-expert, @tui-charm-expert
├── @devops-sre → deployment-guide
└── End of chain

UTILITY (Parallel)
├── @debugger → debug-report, known-issues
├── @refactor-architect → refactoring-overview
├── @doc-janitor → Archives per-feature docs (manual trigger)
└── @project-bro → Read-only helper
```

---

## Archival Flow

```
per-feature doc created
       ↓
 active/{category}/{work-unit}.md
       ↓
 (skill completes work)
       ↓
 review/{category}/{work-unit}.md  ← notify_user for approval
       ↓
 (user approves)
       ↓
 closed/{category|work-unit}/      ← @doc-janitor archives
```

**Trigger for archival:** Manual via `/doc-cleanup` workflow or "archive completed work" command.

---

## Chain Integrity

| Check | Status |
|-------|--------|
| Every doc has creator | ✅ |
| Every per-feature doc has archival path | ✅ |
| Every handoff has receiver | ✅ |
| Return paths defined | ✅ |
| Living docs have updaters | ✅ |
| Meta docs (registries) have workflow | ✅ |
