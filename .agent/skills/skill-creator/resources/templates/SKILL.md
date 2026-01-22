---
name: {{SKILL_NAME}}
description: {{SKILL_DESCRIPTION}}
---

# {{SKILL_TITLE}}

{{SKILL_PURPOSE_SUMMARY}}

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

## Artifact Ownership

- **Creates**: `docs/<category>/<artifact>.md`
- **Reads**: `<previous skill artifacts>`
- **Updates**: `docs/AGENTS.md` (status + timestamp)

## Handoff Protocol

> [!CAUTION]
> **BEFORE handoff:**
> 1. Save final document to `docs/` path
> 2. Update `docs/AGENTS.md` status to ✅ Done
> 3. Use `notify_user` for final approval
> 4. THEN delegate to next skill

## Resources

- `scripts/`:
    - `example_script.py`: {{SCRIPT_DESC}}
- `references/`:
    - `reference.md`: {{REF_DESC}}
