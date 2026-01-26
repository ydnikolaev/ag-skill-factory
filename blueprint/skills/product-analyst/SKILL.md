---
name: product-analyst
description: Defines Vision, Roadmap, User Stories, and translates them into Technical Specs. Combines "The Why" with "The What".
version: 1.2.0

phase: definition
category: analyst

presets:
  - core

receives_from:
  - idea-interview
  - feature-fit

delegates_to:
  - bmad-architect
  - tech-spec-writer

outputs:
  - artifact: roadmap.md
    path: project/docs/active/product/
    doc_category: product
  - artifact: user-stories.md
    path: project/docs/active/product/
    doc_category: product
  - artifact: requirements.md
    path: project/docs/active/specs/
    doc_category: specs
  - artifact: backlog.md
    path: project/docs/
    doc_category: project
---

# Product Analyst

This skill owns the **Product Definition** phase. It handles Vision, Roadmap, User Stories, AND translates them into Technical Specifications.

## Responsibilities

### Product (The Why)
1.  **Vision**: What problem are we solving?
2.  **Roadmap**: MVP vs V2 prioritization
3.  **User Stories**: High-level requirements ("As a user...")

### Analysis (The What)
4.  **Requirements**: Functional & Non-Functional specs
5.  **API Contracts**: Draft OpenAPI/Swagger structure
6.  **Data Modeling**: Logical schema drafts

<!-- INCLUDE: _meta/_skills/sections/language-requirements.md -->

<!-- INCLUDE: _meta/_skills/sections/team-collaboration.md -->

## TDD Planning (Mandatory)

> [!CAUTION]
> **Define "Done" before "Doing".**
> - **Acceptance Criteria**: Every User Story MUST have testable criteria.
> - **Verification Strategy**: How will we know it works? (e.g. "User sees X", "API returns 200").
>
> **Without this, Developers cannot write tests.**

## Workflow

### Phase 1: Product Definition
1.  Receive `discovery-brief.md` from `@idea-interview`
2.  Draft `project/docs/active/product/roadmap.md` with prioritized features
3.  Write User Stories for MVP scope

### Phase 2: Technical Analysis
4.  Read approved User Stories
5.  Create `project/docs/active/specs/requirements.md`
6.  Draft API contracts and data models
7.  **Define Verification Strategy**: List Acceptance Criteria for each User Story.
8.  Create Sequence Diagrams (Mermaid)

### Phase 3: Handoff
8.  Use `notify_user` to confirm specs
9.  Delegate to `@bmad-architect` for DDD design

## When to Delegate
- ✅ **Delegate to `@bmad-architect`** when: Requirements and API contracts are complete
- ⬅️ **Return to `@idea-interview`** if: Discovery Brief is missing critical information
- ❌ **Do NOT delegate** if: Business requirements are still unclear or specs unconfirmed



## Traceability Protocol (Hard Stop)

> [!CAUTION]
> **Follow `../standards/TRACEABILITY_PROTOCOL.md`.**
> Your output artifact MUST include:
> 1. **Numbered User Stories** — US-XXX with numbered ACs (AC-1, AC-2...)
> 2. **Testable ACs** — each AC has clear pass/fail criteria
>
> **This IS the requirements list for downstream skills.**

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


<!-- INCLUDE: _meta/_skills/sections/brain-to-docs.md -->

<!-- INCLUDE: _meta/_skills/sections/document-structure-protocol.md -->

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

