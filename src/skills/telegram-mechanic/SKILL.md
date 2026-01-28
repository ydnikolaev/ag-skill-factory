---
# === SECTION 1: IDENTITY ===
name: telegram-mechanic
description: Expert in Telegram Bot API, Webhooks, and Mini App Authentication.
version: 3.0.0
phase: architecture
category: technical
scope: project
tags:
  - telegram
  - bot-api
  - webhooks
  - tma-auth

# === SECTION 2: CAPABILITIES ===
mcp_servers:
  - context7
  - telegram-docs
allowed_tools:
  - notify_user
  - view_file
  - write_to_file
  - grep_search
  - run_command
dependencies: []
context:
  required:
    - path: project/docs/active/architecture/
      purpose: Context map
  optional:
    - path: project/CONFIG.yaml
      purpose: Stack decisions
reads:
  - type: context_map
    from: project/docs/active/architecture/
produces:
  - type: webhook_config
  - type: auth_middleware_spec

# === SECTION 3: WORKFLOW ===
presets:
  - tma
receives_from:
  - skill: bmad-architect
    docs:
      - doc_type: context-map
        trigger: design_complete
delegates_to:
  - skill: backend-go-expert
    docs:
      - doc_type: webhook-config
        trigger: design_complete
  - skill: tma-expert
    docs:
      - doc_type: webhook-config
        trigger: design_complete
return_paths: []

# === SECTION 4: DOCUMENTS ===
requires:
  - doc_type: context-map
    status: Approved
creates:
  - doc_type: webhook-config
    path: project/docs/active/architecture/
    doc_category: architecture
    lifecycle: per-feature
    initial_status: Draft
    trigger: design_complete
updates:
  - doc_type: artifact-registry
    path: project/docs/
    lifecycle: living
    trigger: on_complete
archives:
  - doc_type: webhook-config
    destination: project/docs/closed/<work-unit>/
    trigger: qa_signoff

# === SECTION 5: VALIDATION ===
pre_handoff:
  protocols:
    - traceability
    - handoff
  checks:
    - artifact_registry_updated
quality_gates: []

# === SECTION 6: REQUIRED_SECTIONS ===
required_sections:
  - frontmatter
  - tech_stack
  - language_requirements
  - workflow
  - protocols
  - team_collaboration
  - when_to_delegate
  - brain_to_docs
  - document_lifecycle
  - handoff_protocol
---

# Telegram Mechanic

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


This skill is the **Gateway**. It manages the Bot, Webhooks, and Security.

## Responsibilities
1.  **Bot API**: Menu setup, Commands, Deep Links.
2.  **Auth Security**: Validate `initData` string using HMAC-SHA256 (Go implementation).
3.  **Entry Point**: Ensure the Mini App opens correctly (`setChatMenuButton`).

<!-- INCLUDE: _meta/_skills/sections/language-requirements.md -->

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


<!-- INCLUDE: _meta/_skills/sections/brain-to-docs.md -->

## Document Lifecycle

> **Protocol**: [`DOCUMENT_STRUCTURE_PROTOCOL.md`](../standards/DOCUMENT_STRUCTURE_PROTOCOL.md)

| Operation | Document | Location | Trigger |
|-----------|----------|----------|---------|
| 🔵 Creates | webhook-config.md | `active/bot/` | Bot setup complete |
| 📖 Reads | context-map.md | `active/architecture/` | On activation |
| 📝 Updates | ARTIFACT_REGISTRY.md | `project/docs/` | On create, on complete |
| 🟡 To Review | webhook-config.md | `review/bot/` | Ready for implementation |
| ✅ Archive | — | `closed/<work-unit>/` | @doc-janitor on final approval |

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

## Antigravity Best Practices
- Use `task_boundary` when setting up a new Bot from scratch.
- Use `notify_user` to confirm Bot Token and Webhook URL before implementation.

