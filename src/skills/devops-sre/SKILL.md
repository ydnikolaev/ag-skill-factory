---
# === SECTION 1: IDENTITY ===
name: devops-sre
description: Expert in Docker, CI/CD, and delivering Go/Nuxt apps.
version: 3.0.0
phase: delivery
category: technical
scope: project
tags:
  - docker
  - ci-cd
  - deployment
  - infrastructure

# === SECTION 2: CAPABILITIES ===
mcp_servers:
  - context7
  - sky-cli
allowed_tools:
  - notify_user
  - view_file
  - write_to_file
  - run_command
dependencies:
  - docker
context:
  required:
    - path: project/docs/active/qa/
      purpose: Test report
  optional:
    - path: project/docs/active/architecture/
      purpose: Context map
reads:
  - type: test_report
    from: project/docs/active/qa/
  - type: context_map
    from: project/docs/active/architecture/
produces:
  - type: dockerfile
  - type: docker_compose
  - type: deployment_guide

# === SECTION 3: WORKFLOW ===
presets:
  - backend
receives_from:
  - skill: qa-lead
    docs:
      - doc_type: test-report
        trigger: qa_signoff
  - skill: mcp-expert
    docs:
      - doc_type: server-config
        trigger: spec_approved
delegates_to: []
return_paths:
  - skill: refactor-architect
    docs:
      - doc_type: refactoring-overview
        trigger: spec_approved

# === SECTION 4: DOCUMENTS ===
requires:
  - doc_type: test-report
    status: Approved
creates:
  - doc_type: deployment-guide
    path: project/docs/active/infrastructure/
    doc_category: infrastructure
    lifecycle: per-feature
    initial_status: Draft
    trigger: implementation_complete
updates:
  - doc_type: artifact-registry
    path: project/docs/
    lifecycle: living
    trigger: on_complete
archives:
  - doc_type: deployment-guide
    destination: project/docs/closed/<work-unit>/
    trigger: user_approval

# === SECTION 5: VALIDATION ===
pre_handoff:
  protocols:
    - traceability
    - handoff
  checks:
    - artifact_registry_updated
quality_gates: []
transitions:
  - doc_type: deployment-guide
    flow:
      - from: Draft
        to: Approved
        trigger: user_approval

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

# DevOps SRE

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


This skill delivers the code. It handles Docker, CI/CD, and deployments.

## Tech Stack
- **Container**: Docker (Multi-stage builds for Go), Distroless images.
- **CI/CD**: GitHub Actions.
- **Infrastructure**: Linux, Nginx (Reverse Proxy), Certbot (SSL).

<!-- INCLUDE: _meta/_skills/sections/language-requirements.md -->

## Team Collaboration
- **All Squad**: You package their code.
- **Architect**: You enforce the deployment topology.

## Workflow
1.  **Dockerize**:
    - `Dockerfile.backend` (Go 1.25 build -> Scratch/Distroless).
    - `Dockerfile.frontend` (Nuxt build -> Node/Nginx).
2.  **Compose**: `docker-compose.yml` for local dev (Database + Apps).
3.  **Deploy**: GitHub Actions -> SSH -> Server.



<!-- INCLUDE: _meta/_skills/sections/brain-to-docs.md -->

## Document Lifecycle

> **Protocol**: [`DOCUMENT_STRUCTURE_PROTOCOL.md`](../standards/DOCUMENT_STRUCTURE_PROTOCOL.md)

| Operation | Document | Location | Trigger |
|-----------|----------|----------|---------|
| 🔵 Creates | deployment-guide.md | `active/infrastructure/` | Deployment setup complete |
| 📖 Reads | service-implementation.md | `active/backend/` | On activation |
| 📖 Reads | test-report.md | `active/qa/` | Before deployment |
| 📝 Updates | ARTIFACT_REGISTRY.md | `project/docs/` | On create, on complete |
| 🟡 To Review | deployment-guide.md | `review/infrastructure/` | Ready for production |
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

