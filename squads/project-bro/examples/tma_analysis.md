# Project-Bro Example: Analyzing a TMA Project

This example shows how project-bro analyzes a typical TMA project state.

## Scenario
User asks: "bro, where are we with xlinefitness?"

## Analysis Process

### Step 1: Read AGENTS.md
```markdown
# From docs/AGENTS.md:
| Artifact | Status | Owner |
|----------|--------|-------|
| discovery-brief.md | ✅ Done | idea-interview |
| roadmap.md | ✅ Done | product-analyst |
| context-map.md | 🔄 In Progress | bmad-architect |
| backend-api.yaml | ⏳ Pending | bmad-architect |
```

### Step 2: Check Code Structure
```
xlinefitness/
├── backend/           # Go backend
│   ├── cmd/api/       # Main entry
│   ├── internal/      # Business logic
│   │   ├── domain/    # 2 aggregates defined
│   │   └── app/       # 3 use cases
│   └── go.mod
├── frontend/          # Nuxt 4 TMA
│   ├── pages/         # 4 pages
│   └── components/    # 8 components
└── docs/              # Documentation
```

### Step 3: Summarize

**Response:**
> **XLineFitness Status:**
> 
> | Phase | Status | Notes |
> |-------|--------|-------|
> | Discovery | ✅ Done | Brief complete |
> | Definition | ✅ Done | Roadmap approved |
> | Architecture | 🔄 70% | Context map needs review |
> | Backend | 🔄 40% | 2/5 aggregates done |
> | Frontend | 🔄 30% | Basic pages, no styling |
> | QA | ⏳ 0% | Not started |
> 
> **Blockers:**
> - Architecture review pending → blocks backend completion
> 
> **Recommendation:**
> Activate `@bmad-architect` to finish context-map, then continue backend.
