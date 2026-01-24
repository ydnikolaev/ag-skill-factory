---
description: Document cleanup and structure enforcement via @doc-janitor
---

# /doc-cleanup Workflow

Activate `@doc-janitor` to audit and fix document structure.

## Steps

### 1. Activate Doc Janitor
Activate `@doc-janitor` skill.

### 2. Dry-Run Audit
Doc Janitor will:
- Scan `project/docs/` structure
- Detect legacy format
- Find orphan files
- Identify archive candidates
- Generate change report

### 3. Review Report
Review the dry-run report showing:
- 🔧 Structure fixes (folder moves)
- 📝 Format fixes (frontmatter)
- 📦 Archive actions (closed work)
- 📋 ARTIFACT_REGISTRY.md updates

### 4. Apply Changes
Reply "apply" to execute all planned changes.

Doc Janitor will:
- Create missing folders
- Move files to correct locations
- Add missing frontmatter
- Archive completed Work Units
- Update ARTIFACT_REGISTRY.md

// turbo
### 5. Commit
```bash
git add project/docs/
git commit -m "chore(docs): doc-janitor cleanup"
```

### 6. Summary
Doc Janitor reports completion summary.
