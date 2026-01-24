# Artifact Registry System

Reference document for managing project artifacts across all skills.

## Single Source of Truth

Every project using Antigravity skills MUST have `docs/ARTIFACT_REGISTRY.md` — the artifact registry.

## Standard Structure

```
docs/
├── ARTIFACT_REGISTRY.md       # 📋 Artifact Registry (required)
│
├── active/                    # 🔵 Current work — in progress
│   ├── discovery/
│   ├── product/
│   ├── specs/
│   ├── architecture/
│   └── design/
│
├── review/                    # 🟡 Awaiting approval
│   └── (same subfolders)
│
└── closed/                    # ✅ Done — archived, read-only
    ├── sprints/
    ├── features/
    └── refactoring/
```

## Lifecycle States

| State | Folder | Meaning |
|-------|--------|---------|
| 🔵 **Active** | `active/` | Currently being worked on |
| 🟡 **Review** | `review/` | Finished, awaiting approval |
| ✅ **Closed** | `closed/` | Approved, archived |

## Movement Rules

1. **Create** → always in `active/`
2. **Ready for handoff** → move to `review/`
3. **User approves** → move to `closed/<context>/`
4. `closed/` is **read-only** — never modified

## Ownership Matrix

| Artifact | Owner | Path |
|----------|-------|------|
| Discovery Brief | `@idea-interview` | `active/discovery/discovery-brief.md` |
| Roadmap | `@product-analyst` | `active/product/roadmap.md` |
| User Stories | `@product-analyst` | `active/product/user-stories.md` |
| Requirements | `@product-analyst` | `active/specs/requirements.md` |
| Context Map | `@bmad-architect` | `active/architecture/context-map.md` |
| API Contracts | `@bmad-architect` | `active/architecture/api-contracts.yaml` |
| Tech Spec | `@tech-spec-writer` | `active/specs/<feature>-tech-spec.md` |
| Test Cases | `@qa-lead` | `active/qa/test-cases.md` |
| Design Tokens | `@ux-designer` | `active/design/tokens.json` |

## ARTIFACT_REGISTRY.md Template

```markdown
# Artifact Registry

> **Project**: <project-name>
> **Current Focus**: `🔵 <active-work-unit>`

---

## 🔵 Active: <Work Unit Name>

> **Type**: Feature | **Started**: YYYY-MM-DD | **Lead**: @skill-name

### Artifacts

| Phase | Document | Owner | Status |
|-------|----------|-------|--------|
| Discovery | [discovery-brief.md](active/discovery/) | @idea-interview | ✅ |
| Definition | [user-stories.md](active/product/) | @product-analyst | ✅ |
| Architecture | [context-map.md](active/architecture/) | @bmad-architect | 🔵 IN PROGRESS |

---

## ✅ Closed

<details>
<summary><b>Sprint 01: Name</b> (YYYY-MM-DD)</summary>

| Document | Owner | Archive |
|----------|-------|---------|
| discovery-brief.md | @idea-interview | `closed/sprints/01/` |

</details>

---

## Quick Links

| Work Unit | Type | Status | Lead |
|-----------|------|--------|------|
| feature-name | Feature | 🔵 Active | @skill |

## Legend

| Icon | Meaning |
|------|---------|
| 🔵 | Active — in progress |
| 🟡 | Review — awaiting approval |
| ✅ | Done / Closed |
```

## Skill Integration

Each skill that creates artifacts MUST have:

```markdown
## Artifact Ownership
- **Creates**: `active/<path>/<file>.md`
- **Reads**: `active/<path>/<file>.md`
- **Updates**: `ARTIFACT_REGISTRY.md`
```
