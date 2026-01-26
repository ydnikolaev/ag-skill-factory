# Document Structure Protocol

> **Status**: Active
> **Enforced by**: `@doc-janitor`
> **Last Updated**: 2026-01-25

This protocol defines the **mandatory** document structure for all projects using Antigravity skills.

---

## Folder Structure (Lifecycle-Based)

```
project/docs/
├── ARTIFACT_REGISTRY.md       # 📋 Single Source of Truth
│
├── active/                    # 🔵 Currently being worked on
│   ├── discovery/             # Discovery phase docs
│   ├── product/               # Roadmap, user stories
│   ├── specs/                 # Tech specs
│   ├── architecture/          # Context maps, API contracts
│   ├── design/                # Tokens, design system
│   ├── backend/               # Implementation reports
│   ├── frontend/              # UI implementation
│   └── qa/                    # Test cases, reports
│
├── review/                    # 🟡 Awaiting approval
│   └── (same subfolders)
│
└── closed/                    # ✅ Done — archived, read-only
    ├── sprints/sprint-XX/
    ├── features/<feature-name>/
    ├── refactoring/<refactor-name>/
    └── bugs/<bug-id>/
```

---

## Lifecycle States

| State | Folder | Meaning | Who moves here |
|-------|--------|---------|----------------|
| 🔵 **Active** | `active/` | Being worked on | Creator skill |
| 🟡 **Review** | `review/` | Ready for approval | Owner skill |
| ✅ **Closed** | `closed/` | Done, archived | User or `@doc-janitor` |

---

## Movement Rules

1. **Create** → always in `active/<category>/`
2. **Ready for handoff** → move to `review/<category>/`
3. **User approves** → move to `closed/<context>/`
4. Skills read from `active/` or `review/`
5. `closed/` is **read-only** — never modified

---

## Document Requirements

### YAML Frontmatter (Required)

Every document MUST have:

```yaml
---
status: Draft | Review | Approved
owner: @skill-name
next: @skill-name
work_unit: sprint-03 | feature-name | refactor-name
created: YYYY-MM-DD
updated: YYYY-MM-DD
---
```

### Required Sections

| Section | When Required |
|---------|---------------|
| `## Upstream Documents` | When document references other docs |
| `## Requirements Checklist` | When implementing user stories |

---

## ARTIFACT_REGISTRY.md Format

### Work Units Structure

```markdown
# Artifact Registry

> **Project**: <project-name>
> **Current Focus**: `🔵 <active-work-unit>`

---

## 🔵 Active: <Work Unit Name>

> **Type**: Feature | Sprint | Refactoring
> **Started**: YYYY-MM-DD
> **Lead**: @skill-name

### Artifacts

| Phase | Document | Owner | Status |
|-------|----------|-------|--------|
| Discovery | [discovery-brief.md](active/discovery/) | @idea-interview | ✅ |
| Definition | [user-stories.md](active/product/) | @product-analyst | ✅ |
| Architecture | [context-map.md](active/architecture/) | @bmad-architect | 🔵 |

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
| Sprint 01 | Sprint | ✅ Closed | — |
```

---

## Archive Rules

### Archive Paths

| Work Type | Path Template | Example |
|-----------|---------------|---------|
| Sprint | `closed/sprints/sprint-XX/` | `closed/sprints/sprint-03/` |
| Feature | `closed/features/<name>/` | `closed/features/forum-topics/` |
| Refactoring | `closed/refactoring/<name>/` | `closed/refactoring/test-coverage/` |
| Bug | `closed/bugs/<id>/` | `closed/bugs/BUG-042/` |

### What Gets Archived

All documents with:
- `status: Approved` in frontmatter
- Status = ✅ in ARTIFACT_REGISTRY.md
- User confirmation

### Archive Is Read-Only

> [!CAUTION]
> Files in `closed/` are NEVER modified.
> If changes needed → copy back to `active/`, modify, re-archive.

---

## Status Icons

| Icon | Meaning | Folder |
|------|---------|--------|
| 🔵 | Active — in progress | `active/` |
| 🟡 | Review — awaiting approval | `review/` |
| ✅ | Done / Closed | `closed/` |
| ⏳ | Pending — not started | (no file yet) |

---

## Skill Responsibilities

### Creating Documents

1. Create in `active/<category>/`
2. Add YAML frontmatter
3. Update ARTIFACT_REGISTRY.md

### Completing Documents

1. Change frontmatter `status: Review`
2. Move to `review/<category>/`
3. Update ARTIFACT_REGISTRY.md status

### After User Approval

1. `@doc-janitor` moves to `closed/<context>/`
2. Updates ARTIFACT_REGISTRY.md with `<details>` block
3. Commits changes

---

## Enforcement

This protocol is enforced by:
- `@doc-janitor` skill (manual cleanup)
- Pre-Handoff Validation in all skills
- `/doc-cleanup` workflow

Violations will be flagged and corrected automatically.
