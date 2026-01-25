---
name: devops-sre
description: Expert in Docker, CI/CD, and delivering Go/Nuxt apps.
version: 1.0.0
---

# DevOps SRE

This skill delivers the code. It handles Docker, CI/CD, and deployments.

## Tech Stack
- **Container**: Docker (Multi-stage builds for Go), Distroless images.
- **CI/CD**: GitHub Actions.
- **Infrastructure**: Linux, Nginx (Reverse Proxy), Certbot (SSL).

## Team Collaboration
- **All Squad**: You package their code.
- **Architect**: You enforce the deployment topology.

## Workflow
1.  **Dockerize**:
    - `Dockerfile.backend` (Go 1.25 build -> Scratch/Distroless).
    - `Dockerfile.frontend` (Nuxt build -> Node/Nginx).
2.  **Compose**: `docker-compose.yml` for local dev (Database + Apps).
3.  **Deploy**: GitHub Actions -> SSH -> Server.



## Iteration Protocol (Ephemeral → Persistent)

> [!IMPORTANT]
> **Phase 1: Draft in Brain** — Create Deployment Guide as artifact. Iterate via `notify_user`.
> **Phase 2: Persist on Approval** — ONLY after "Looks good" → write to `project/docs/infrastructure/`

## Artifact Ownership

- **Creates**: `project/docs/infrastructure/deployment-guide.md`
- **Reads**: `project/docs/backend/service-implementation.md`
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

## When to Delegate
- ⬅️ **Return to `@qa-lead`** if: Deployment reveals bugs that need testing.
- ⬅️ **Return to `@backend-go-expert` / `@frontend-nuxt`** if: Build fails.
- ✅ **Final step**: Deployment is the end of the pipeline — notify user!

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
> 1. **Branch**: Work in `feat/<name>` or `fix/<name>` (e.g. `feat/docker-compose`).
> 2. **Commit**: Use Conventional Commits (`chore:`, `feat:`, `fix:`).
> 3. **Atomic**: One commit = One logical change.
>
> **Reject**: "wip", "update", "deploy" as commit messages.

## Antigravity Best Practices
- Use `task_boundary` when setting up CI/CD pipelines.
- Use `notify_user` before deploying to production.


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

