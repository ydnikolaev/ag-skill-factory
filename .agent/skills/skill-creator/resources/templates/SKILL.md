---
name: {{SKILL_NAME}}
description: {{SKILL_DESCRIPTION}}
---

# {{SKILL_TITLE}}

{{SKILL_PURPOSE_SUMMARY}}

> [!IMPORTANT]
> ## First Step: Read Project Config
> Before making technical decisions, **always check**:
> ```
> project/CONFIG.yaml
> ```
> This file defines: stack versions, modules, architecture style, features.
> **Never assume defaults — verify against CONFIG.yaml first.**

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
> - Keep Q&A and sketches here — do NOT pollute `docs/`
>
> **Phase 2: Persist on Approval**
> - ONLY after "Looks good" → write final to `docs/` path
> - Update file status: `Draft` → `Approved` in header
> - The `docs/` folder is the ONLY memory that survives sessions

## Artifact Ownership

- **Creates**: `docs/<category>/<artifact>.md`
- **Reads**: `<previous skill artifacts>`
- **Updates**: `docs/AGENTS.md` (status + timestamp)

## Handoff Protocol

> [!CAUTION]
> **BEFORE delegating to next skill:**
> 1. ✅ Final document exists in `docs/` (not just brain artifact)
> 2. ✅ File header changed from `Draft` to `Approved`
> 3. ✅ `docs/AGENTS.md` updated to ✅ Done
> 4. ✅ User approved via `notify_user`
> 5. THEN delegate

## Resources

- `scripts/`:
    - `example_script.py`: {{SCRIPT_DESC}}
- `references/`:
    - `reference.md`: {{REF_DESC}}
