---
name: {{SKILL_NAME}}
description: {{SKILL_DESCRIPTION}}
---

# {{SKILL_TITLE}}

{{SKILL_PURPOSE_SUMMARY}}

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
> **Use project MCP server** (named after project, e.g. `mcp_xlinefitness-bot_*`):
> - `list_resources` → see available project data
> - `*_tools` → project-specific actions (db, cache, jobs, etc.)
> 
> **Use `mcp_context7`** for library docs:
> - Check `mcp.yaml → context7.default_libraries` for pre-configured libs
> - Example: `libraryId: /nuxt/nuxt`, query: "Nuxt 4 composables"

## When to use this skill

- **Trigger 1**: Description of situation...
- **Trigger 2**: Description of situation...
- **Anti-pattern**: Do NOT use this skill when...

## Decision Tree

Before acting, determine the specific scenario:

1.  **IF** [Condition A]:
    - Use `scripts/script_a.py`
    - Refer to `references/guide_a.md`
2.  **IF** [Condition B]:
    - Follow the manual workflow below
3.  **ELSE**:
    - Ask user for clarification

## Workflow

### Phase 1: Preparation
1.  Verify prerequisites...
2.  Check for existing files...

### Phase 2: Execution
> **Tip**: Use `task_boundary` to track progress if this is a long task.

1.  Step 1...
2.  Step 2...

### Phase 3: Verification
1.  Run validation checks...
2.  Notify user...


## Iteration Protocol (Ephemeral → Persistent)

> [!IMPORTANT]
> **Phase 1: Draft in Brain**
> - Create drafts as UI Artifacts (in `brain/` conversation directory)
> - Iterate with user via `notify_user` until approved
> - Keep Q&A and sketches here — do NOT pollute `project/docs/`
>
> **Phase 2: Persist on Approval**
> - ONLY after "Looks good" → write final to `project/docs/` path
> - Update file status: `Draft` → `Approved` in header
> - The `project/docs/` folder is the ONLY memory that survives sessions

## Artifact Ownership

- **Creates**: `project/docs/<category>/<artifact>.md`
- **Reads**: `<previous skill artifacts>`
- **Updates**: `project/docs/AGENTS.md` (status + timestamp)

## TDD Protocol

> [!CAUTION]
> **NO CODE WITHOUT FAILING TEST.**
> - **Red**: Write test → Fail.
> - **Green**: Write code → Pass.
> - **Refactor**: Clean up.
> **NEVER write implementation code without a failing test.**
> **NEVER fix a bug without a reproduction case.**

## Handoff Protocol

> [!CAUTION]
> **BEFORE delegating to next skill:**
> 1. ✅ Final document exists in `project/docs/` (not just brain artifact)
> 2. ✅ File header changed from `Draft` to `Approved`
> 3. ✅ `project/docs/AGENTS.md` updated to ✅ Done
> 4. ✅ User approved via `notify_user`
> 5. THEN delegate

## Resources

- `scripts/`:
    - `example_script.py`: {{SCRIPT_DESC}}
- `references/`:
    - `reference.md`: {{REF_DESC}}
