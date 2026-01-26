---
name: debugger
description: "Systematic debugging skill. 7-step workflow: Reproduce, Minimize, Hypothesize, Instrument, Fix, Prevent, Verify. Activate when troubleshooting errors."
version: 1.2.0

phase: utility
category: utility

presets:
  - backend
  - minimal

receives_from:
  - qa-lead
  - backend-go-expert
  - frontend-nuxt

delegates_to:
  - qa-lead

outputs:
  - doc_type: debug-report
    path: project/docs/active/bugs/
    doc_category: bugs
    lifecycle: per-feature
  - doc_type: known-issues
    path: project/docs/
    doc_category: project
    lifecycle: living
---

# Debugger

> [!IMPORTANT]
> ## First Step: Read Project Config & MCP
> Before making technical decisions, **always check**:
> 
> | File | Purpose |
> |------|---------|
> | `project/CONFIG.yaml` | Stack versions, modules, architecture |
> | `mcp.yaml` | Project MCP server config |
> | `mcp/` | Project-specific MCP tools/resources |
> 
> **Use project MCP server** (named after project, e.g. `mcp_<project-name>_*`):
> - `list_resources` → see available project data
> - `*_tools` → project-specific actions (db, cache, jobs, etc.)
> 
> **Use `mcp_context7`** for library docs:
> - Check `mcp.yaml → context7.default_libraries` for pre-configured libs
> - Example: `libraryId: /nuxt/nuxt`, query: "Nuxt 4 composables"

## When to Activate

- Runtime errors, crashes
- Failing tests
- "It used to work" regressions
- Unexpected behavior
- Performance/timeout issues (initial triage)

## The 7-Step Debug Workflow

> [!CAUTION]
> **DO NOT SKIP STEPS!** Each is critical for systematic debugging.

### Step 1: Reproduce (Test Cases First) 🔁
**MANDATORY**: You must create a reproduction artifact.
- **Unit Test**: If logic error, write a failing unit test.
- **Script**: If integration error, write a standalone repro script.
- **Config**: If environment error, document exact config to repro.

**Status Check**: Do you have a "Red" test? If no, go back.

### Step 2: Minimize 🎯
Reduce to smallest repro:
- One file
- One function
- Smallest dataset
- Remove unrelated code

### Step 3: Hypothesize 🧠
Form 2-5 hypotheses, ranked by likelihood:
1. Most likely cause
2. Second candidate
3. Edge case possibility
4. (Optional) Unlikely but possible

### Step 4: Instrument 🔧
Add temporary debugging:
- Logging statements
- Assertions
- Breakpoints
- Print variables at key points

Use existing diagnostics if available.

### Step 5: Fix ⚡
Apply smallest change that removes root cause:
- Don't over-engineer
- Fix the actual problem, not symptoms
- Keep changes minimal

### Step 6: Prevent 🛡️
Add protection:
- Regression test
- Validation guard
- Assertion
- Error handling improvement

### Step 7: Verify ✅
Run verification:
- The failing case now passes
- Related test suites pass
- No new regressions

## Report Format

When reporting a fix, use this structure:

```markdown
### Symptom
(What was wrong)

### Repro Steps
1. ...
2. ...

### Root Cause
(Why it happened)

### Fix
(What you changed)

### Regression Protection
(Test or guard added)

### Verification
- Commands run:
- Results:
```

<!-- INCLUDE: _meta/_skills/sections/language-requirements.md -->

## Team Collaboration

- **Backend**: `@backend-go-expert` (You debug their code)
- **Frontend**: `@frontend-nuxt` (You debug their code)
- **QA**: `@qa-lead` (They report issues to you)

## When to Delegate

- ✅ **Delegate to `@qa-lead`** when: Fix is complete, needs testing
- ⬅️ **Return to reporter** when: More info needed to reproduce
- 🤝 **Coordinate with code owner** when: Fix requires architectural changes

## Iteration Protocol (Ephemeral → Persistent)

> [!IMPORTANT]
> **Phase 1: Draft in Brain** — Create debug report as artifact. Iterate via `notify_user`.
> **Phase 2: Persist on Approval** — ONLY after fix verified → write to `project/docs/bugs/`

## Document Lifecycle

> **Protocol**: [`DOCUMENT_STRUCTURE_PROTOCOL.md`](../standards/DOCUMENT_STRUCTURE_PROTOCOL.md)

| Operation | Document | Location | Trigger |
|-----------|----------|----------|---------|
| 🔵 Creates | `<issue-name>.md` | `active/bugs/` | Debug report complete |
| 📖 Reads | Issue description, logs | — | On activation |
| 📖 Reads | Implementation docs | `active/backend/`, `active/frontend/` | Understanding context |
| 📝 Updates | ARTIFACT_REGISTRY.md | `project/docs/` | On create, on complete |
| 🟡 To Review | `<issue-name>.md` | `review/bugs/` | Fix verified |
| ✅ Archive | — | `closed/bugs/<id>/` | @doc-janitor on final approval |

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
> **BEFORE handoff:**
> 1. Save final document to `project/docs/` path
> 2. Change file status from `Draft` to `Approved` in header/frontmatter
> 3. Update `project/docs/ARTIFACT_REGISTRY.md` status to ✅ Done
> 4. Use `notify_user` for final approval
> 5. THEN delegate to next skill

## Tech Debt Protocol (Hard Stop)

> [!CAUTION]
> **Follow `../standards/TECH_DEBT_PROTOCOL.md`.**
> When creating workarounds:
> 1. Add `// TODO(TD-XXX): description` in code
> 2. Register in `project/docs/TECH_DEBT.md`
>
> **Forbidden:** Untracked TODOs, undocumented hardcoded values.

## Git Protocol (Hard Stop)

> [!CAUTION]
> **Follow `../standards/GIT_PROTOCOL.md`.**
> 1. **Branch**: Create `fix/<bug-name>` branch before fixing.
> 2. **Commit**: Use `fix(<scope>): <description>` format.
> 3. **Atomic**: One fix = One commit (regression test included).
>
> **Reject**: "wip", "debug", "fixed" as commit messages.

## Antigravity Best Practices

- Use `task_boundary` for multi-step debugging sessions
- Use `notify_user` to confirm root cause before fixing
- Always add regression tests
