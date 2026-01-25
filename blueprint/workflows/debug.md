---
description: Bug investigation and fix
---

# Debug Workflow

Systematic debugging using 7-step workflow.

## Steps

1. Activate `@debugger` skill
2. Reproduce (mandatory — create failing test)
3. Minimize (smallest repro case)
4. Hypothesize (2-5 ranked causes)
5. Instrument (add logging/assertions)
6. Fix (minimal change)
7. Prevent (add regression test)
8. Verify (all tests pass)

## Output

- Debug report in `active/bugs/<issue-name>.md`
- Regression test added
- Return to `@backend-go-expert` or `@frontend-nuxt` if architectural changes needed

## Trigger

- "There's a bug"
- "This is broken"
- Test failure investigation
