# Artifact Registry System

Reference document for managing project artifacts across all skills.

## Single Source of Truth

Every project using Antigravity skills MUST have `docs/AGENTS.md` — the artifact registry.

## Standard Structure

```
docs/
├── AGENTS.md              # 📋 Artifact Registry (required)
├── discovery/
│   └── discovery-brief.md
├── product/
│   ├── roadmap.md
│   └── user-stories.md
├── specs/
│   └── requirements.md
├── architecture/
│   └── context-map.md
└── qa/
    └── test-cases.md
```

## Ownership Matrix

| Artifact | Owner | Path |
|----------|-------|------|
| Discovery Brief | `@idea-interview` | `docs/discovery/discovery-brief.md` |
| Roadmap | `@product-analyst` | `docs/product/roadmap.md` |
| User Stories | `@product-analyst` | `docs/product/user-stories.md` |
| Requirements | `@product-analyst` | `docs/specs/requirements.md` |
| Context Map | `@bmad-architect` | `docs/architecture/context-map.md` |
| API Contracts | `@bmad-architect` | `docs/architecture/api-contracts.yaml` |
| Test Cases | `@qa-lead` | `docs/qa/test-cases.md` |
| Design Tokens | `@ux-designer` | `docs/design/tokens.json` |

## Lifecycle Rules

### Create
1. Owner skill creates artifact in designated path
2. Owner updates `docs/AGENTS.md` with new entry

### Update
1. Only owner can update artifact
2. Owner updates "Last Updated" in `docs/AGENTS.md`
3. Git commit = version history

### Handoff
1. Downstream skill reads from designated path
2. If artifact missing → return to owner skill

## Versioning

**Git-based** — no manual archive folders.
- Each commit to artifact = version
- Use `git log docs/<path>` to see history
- Tag major versions with git tags if needed

## AGENTS.md Template

```markdown
# Artifact Registry

> Single source of truth. Updated by each skill on create/update.

| Artifact | Owner | Status | Last Updated |
|----------|-------|--------|--------------|
| discovery-brief.md | @idea-interview | ✅ Done | YYYY-MM-DD |
| roadmap.md | @product-analyst | 📝 Draft | YYYY-MM-DD |
| requirements.md | @product-analyst | 🔄 Review | YYYY-MM-DD |
| context-map.md | @bmad-architect | ⏳ Pending | - |
| test-cases.md | @qa-lead | ⏳ Pending | - |

## Status Legend
- ⏳ Pending — not started
- 📝 Draft — in progress
- 🔄 Review — awaiting approval
- ✅ Done — approved, ready for downstream
```

## Skill Integration

Each skill that creates artifacts MUST have:

```markdown
## Artifact Ownership
- **Creates**: `docs/<path>/<file>.md`
- **Reads**: `docs/<path>/<file>.md`
- **Updates**: `docs/AGENTS.md`
```
