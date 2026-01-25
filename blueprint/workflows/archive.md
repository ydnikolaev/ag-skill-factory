---
description: Archive completed work unit
---

# Archive Workflow

Close and archive a completed work unit.

## Steps

1. Activate `@doc-janitor` skill
2. Verify all documents have status = Approved
3. Move documents: `active/` → `closed/<work-unit>/`
4. Update `ARTIFACT_REGISTRY.md` with closed section
5. Commit with `chore(docs):` prefix

## Trigger

- "Archive this sprint"
- "Close this feature"
- Work unit completed and approved
