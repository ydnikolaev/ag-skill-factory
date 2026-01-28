---
# === SECTION 1: IDENTITY ===
name: doc-janitor
description: Enforces document structure, archives completed work, migrates legacy format to lifecycle-based folders. Dry-run first, then apply.
version: 3.0.0
phase: utility
category: utility
scope: project
tags:
  - cleanup
  - structure
  - archive
  - migration

# === SECTION 2: CAPABILITIES ===
mcp_servers: []
allowed_tools:
  - list_dir
  - view_file
  - run_command
  - write_to_file
  - notify_user
dependencies: []
context:
  required:
    - path: project/docs/
      purpose: Document root to organize
  optional:
    - path: project/docs/ARTIFACT_REGISTRY.md
      purpose: Registry to update after cleanup
reads:
  - type: document_frontmatter
    from: all-project-docs
  - type: folder_structure
    from: project/docs/
produces:
  - type: cleanup_report
  - type: archive_structure

# === SECTION 3: WORKFLOW ===
presets:
  - core
receives_from: []
delegates_to: []
return_paths: []

# === SECTION 4: DOCUMENTS ===
requires: []
creates: []
updates:
  - doc_type: artifact-registry
    path: project/docs/
    lifecycle: living
    trigger: on_complete
archives:
  - doc_type: all-per-feature
    destination: project/docs/closed/<work-unit>/
    trigger: user_approval
    final_status: archived

# === SECTION 5: VALIDATION ===
pre_handoff:
  protocols:
    - handoff
  checks:
    - artifact_registry_updated
quality_gates: []

# === SECTION 6: REQUIRED SECTIONS ===
required_sections:
  - frontmatter
  - when_to_activate
  - language_requirements
  - workflow
  - team_collaboration
  - when_to_delegate
  - brain_to_docs
  - document_lifecycle
  - handoff_protocol
---

# Doc Janitor

> **MODE**: AUTONOMOUS EXECUTOR. You clean up and organize project docs.
> ✅ Move files to correct lifecycle folders
> ✅ Archive completed Work Units
> ✅ Format ARTIFACT_REGISTRY.md
> ❌ Do NOT write document content
> ❌ Do NOT make approval decisions

## When to Activate

- "Clean up the docs"
- "Check document structure"
- "Migrate to new format"
- "Archive completed work"
- "/doc-cleanup" workflow

## Role Boundary

| DOES ✅ | DOES NOT ❌ |
|---------|-------------|
| Move files to lifecycle folders | Write document content |
| Create folder structure | Make approval decisions |
| Update ARTIFACT_REGISTRY.md | Create new artifacts |
| Archive completed Work Units | Delete without confirmation |
| Add missing sections | Change document meaning |

> **Protocol**: Enforces `../standards/DOCUMENT_STRUCTURE_PROTOCOL.md`

## Workflow

### Phase 0: Dry-Run (MANDATORY FIRST)

> [!CAUTION]
> **ALWAYS start with dry-run. NEVER apply without user approval.**

1. Scan `project/docs/` structure
2. Generate change report
3. Present via `notify_user`
4. Wait for explicit "apply" command

### Phase 1: Structure Audit

```bash
# Check expected folders
ls -la project/docs/active/ project/docs/review/ project/docs/closed/

# Find orphan files in root
ls project/docs/*.md | grep -v ARTIFACT_REGISTRY
```

Detect legacy signals:
- No `active/` folder → full migration needed
- Files in `project/docs/specs/` → move to `active/specs/`
- AGENTS.md exists → rename to ARTIFACT_REGISTRY.md

### Phase 2: Document Validation

For each document check:
1. YAML frontmatter with `status`, `owner`, `created`, `updated`
2. `## Upstream Documents` section (if applicable)
3. Status matches location (Draft→active, Review→review)

### Phase 3: Archive Identification

**Archive candidates:**
- Status = `Approved` in ARTIFACT_REGISTRY.md
- All requirements marked ✅
- User confirmed completion

**Archive paths:**
| Work Type | Path |
|-----------|------|
| Sprint | `closed/sprints/sprint-XX/` |
| Feature | `closed/features/<name>/` |
| Refactoring | `closed/refactoring/<name>/` |
| Bug | `closed/bugs/<id>/` |

### Phase 4: Report Generation

Create report in brain artifact:

```markdown
# Doc Janitor Report (Dry Run)

## Statistics
- Files scanned: N
- Issues found: N
- Actions planned: N

## Planned Actions

### 🔧 Structure Fixes
| File | Current | New |
|------|---------|-----|
| discovery-brief.md | `docs/` | `docs/active/discovery/` |

### 📦 Archive Actions
| Work Unit | Files | Archive Path |
|-----------|-------|--------------|
| Sprint-03 | 4 | `closed/sprints/sprint-03/` |

### 📋 ARTIFACT_REGISTRY.md
- [ ] Migrate to Work Units format
- [ ] Add Quick Links table

## Approve?
Reply "apply" to execute.
```

### Phase 5: Apply Changes

After user approval:

**Trivial (no confirmation):**
- Create missing folders
- Add YAML frontmatter

**Requires confirmation (shown in report):**
- Move files between folders
- Archive to `closed/`
- Rewrite ARTIFACT_REGISTRY.md

### Phase 6: Commit

```bash
git add project/docs/
git commit -m "chore(docs): doc-janitor cleanup"
```

## Folder Structure Reference

```
project/docs/
├── ARTIFACT_REGISTRY.md       # 📋 Single Source of Truth
│
├── active/                    # 🔵 In progress
│   ├── discovery/
│   ├── product/
│   ├── specs/
│   ├── architecture/
│   ├── design/
│   ├── backend/
│   ├── frontend/
│   └── qa/
│
├── review/                    # 🟡 Awaiting approval
│   └── (same subfolders)
│
└── closed/                    # ✅ Archived, read-only
    ├── sprints/sprint-XX/
    ├── features/<name>/
    ├── refactoring/<name>/
    └── bugs/<id>/
```

## ARTIFACT_REGISTRY.md Format

Must follow Work Units structure:

```markdown
# Artifact Registry

> **Project**: <name>
> **Current Focus**: `🔵 <active-work>`

---

## 🔵 Active: <Work Unit Name>

| Phase | Document | Owner | Status |
|-------|----------|-------|--------|
| Discovery | discovery-brief.md | @idea-interview | ✅ |
| Implementation | impl.md | @backend-go-expert | 🔵 |

---

## ✅ Closed

<details>
<summary><b>Sprint 01</b></summary>

| Document | Owner | Archive |
|----------|-------|---------|
| discovery-brief.md | @idea-interview | `closed/sprints/01/` |

</details>
```

<!-- INCLUDE: _meta/_skills/sections/language-requirements.md -->

## Team Collaboration

- **User** (direct trigger via `/doc-cleanup`)
- **All skills** (they follow protocol during work)

## When to Delegate

- ✅ **Delegate to nothing** — autonomous skill
- ⬅️ **Return to user** when: Dry-run complete, need approval
- ⬅️ **Return to user** when: Ambiguous Work Unit ownership

## Pre-Handoff Validation (Hard Stop)

> [!CAUTION]
> **MANDATORY self-check before `notify_user` or delegation.**

| # | Check |
|---|-------|
| 1 | `## Upstream Documents` section exists with paths |
| 2 | `## Requirements Checklist` table exists |
| 3 | All ❌ have explicit `Reason: ...` |
| 4 | Document in `review/` folder |
| 5 | `ARTIFACT_REGISTRY.md` updated |

**If ANY unchecked → DO NOT PROCEED.**

## Handoff Protocol

> [!CAUTION]
> **BEFORE completing:**
> 1. Dry-run report shown to user
> 2. User explicitly said "apply"
> 3. Changes committed with `chore(docs):` prefix
> 4. Report summary via `notify_user`

<!-- INCLUDE: _meta/_skills/sections/brain-to-docs.md -->

## Document Lifecycle

> **Protocol**: [`DOCUMENT_STRUCTURE_PROTOCOL.md`](../standards/DOCUMENT_STRUCTURE_PROTOCOL.md)

| Operation | Document | Location | Trigger |
|-----------|----------|----------|---------|
| 📝 Updates | ARTIFACT_REGISTRY.md | `project/docs/` | On archive, on cleanup |
| 📁 Creates | `active/`, `review/`, `closed/` folders | `project/docs/` | Structure setup |
| 📁 Moves | Any document | `active/` → `review/` → `closed/` | Lifecycle transitions |
| 📖 Reads | All project docs | `project/docs/` | Audit phase |
| ✅ Archives | Completed documents | `closed/<work-unit>/` | User approves closure |

## Antigravity Best Practices

- Use `task_boundary` with mode EXECUTION during apply phase
- Use `notify_user` for dry-run report and completion
- Never delete files without explicit user confirmation
- Always backup ARTIFACT_REGISTRY.md before rewriting
