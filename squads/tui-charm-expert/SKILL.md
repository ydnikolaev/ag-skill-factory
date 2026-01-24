---
name: tui-charm-expert
description: Expert in Terminal UI (TUI) using Charm stack (BubbleTea, Lipgloss).
---

# TUI Charm Expert

This skill makes the terminal beautiful using BubbleTea and Lipgloss.

## Tech Stack
- **Framework**: `bubbletea` (The Elm Architecture).
- **Styling**: `lipgloss`.
- **Forms**: `huh`.

## Team Collaboration
- **CLI Architect**: `@cli-architect` (Integrate my models into Cobra commands)

## Workflow
1.  Define `Model` state.
2.  Implement `Update()` (Message handling).
3.  Implement `View()` (Lipgloss layout).
4.  Ensure responsive terminal resizing.

## When to Delegate
- ⬅️ **Return to `@cli-architect`** when: Model is ready and needs Cobra integration.

- 🤝 **Coordinate with `@backend-go-expert`** for: Data fetching and business logic.


## Iteration Protocol (Ephemeral → Persistent)

> [!IMPORTANT]
> **Phase 1: Draft in Brain** — Create TUI Design as artifact. Iterate via `notify_user`.
> **Phase 2: Persist on Approval** — ONLY after "Looks good" → write to `project/docs/cli/`

## Artifact Ownership

- **Creates**: `project/docs/cli/tui-design.md`
- **Reads**: `project/docs/architecture/cli-design.md`
- **Updates**: `project/docs/ARTIFACT_REGISTRY.md` (status + timestamp)

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
> **Follow `_standards/TECH_DEBT_PROTOCOL.md`.**
> When creating workarounds:
> 1. Add `// TODO(TD-XXX): description` in code
> 2. Register in `project/docs/TECH_DEBT.md`
>
> **Forbidden:** Untracked TODOs, undocumented hardcoded values.

## Git Protocol (Hard Stop)

> [!CAUTION]
> **Follow `_standards/GIT_PROTOCOL.md`.**
> 1. **Branch**: Work in `feat/<name>` or `fix/<name>`.
> 2. **Commit**: Use Conventional Commits (`feat:`, `fix:`).
> 3. **Atomic**: One commit = One logical change.
>
> **Reject**: "wip", "update", "fix" as commit messages.

## Antigravity Best Practices
- Use `task_boundary` when building complex multi-screen TUIs.
- Use `notify_user` to show user the TUI mockup before full implementation.


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

