---
description: Run QA cycle
---

# QA Workflow

Execute quality assurance testing cycle.

## Steps

1. Activate `@qa-lead` skill
2. Read implementation docs from `active/backend/` and `active/frontend/`
3. Create test cases in `active/qa/test-cases.md`
4. Execute tests (E2E, API, UI)
5. Create test report in `active/qa/test-report.md`
6. If all pass → approve for deployment
7. If failures → return to implementation skills with bug reports

## Output

- `test-cases.md` — test plan
- `test-report.md` — results

## Trigger

- Implementation complete
- "Run QA"
- "Test this feature"
