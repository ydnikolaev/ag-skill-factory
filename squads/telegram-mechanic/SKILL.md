---
name: telegram-mechanic
description: Expert in Telegram Bot API, Webhooks, and Mini App Authentication.
---

# Telegram Mechanic

This skill is the **Gateway**. It manages the Bot, Webhooks, and Security.

## Responsibilities
1.  **Bot API**: Menu setup, Commands, Deep Links.
2.  **Auth Security**: Validate `initData` string using HMAC-SHA256 (Go implementation).
3.  **Entry Point**: Ensure the Mini App opens correctly (`setChatMenuButton`).

## Team Collaboration
- **Architect**: `@bmad-architect` (Activates for TMA/Bot projects)
- **Backend**: `@backend-go-expert` (You provide the Auth Middleware logic)
- **TMA**: `@tma-expert` (You provide the `start_param` parsing)

## Workflow
1.  Register Bot in `@BotFather`.
2.  Configure Webhook (SSL required).
3.  Implement `POST /webhook` handler in Backend.
4.  Implement `initData` validation helper for Backend.

## When to Delegate
- ✅ **Provide to `@backend-go-expert`**: Auth middleware code and validation logic.
- ✅ **Provide to `@tma-expert`**: Deep link format and start_param parsing.

- ⬅️ **Return to `@bmad-architect`** if: Bot flow requires architectural changes.


## Iteration Protocol (Ephemeral → Persistent)

> [!IMPORTANT]
> **Phase 1: Draft in Brain** — Create Webhook Config as artifact. Iterate via `notify_user`.
> **Phase 2: Persist on Approval** — ONLY after "Looks good" → write to `project/docs/bot/`

## Artifact Ownership

- **Creates**: `project/docs/bot/webhook-config.md`
- **Reads**: `project/docs/architecture/context-map.md`
- **Updates**: `project/docs/AGENTS.md` (status + timestamp)

## Handoff Protocol


> [!CAUTION]
> **BEFORE handoff:**
> 1. Save final document to `project/docs/` path
> 2. Change file status from `Draft` to `Approved` in header/frontmatter
> 3. Update `project/docs/AGENTS.md` status to ✅ Done
> 4. Use `notify_user` for final approval
> 5. THEN delegate to next skill

## Antigravity Best Practices
- Use `task_boundary` when setting up a new Bot from scratch.
- Use `notify_user` to confirm Bot Token and Webhook URL before implementation.


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

