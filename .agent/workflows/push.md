---
description: Run full commit workflow, install, and push to remote
---

# /push Workflow

Complete push pipeline: commit checks + install + push.

// turbo-all

## Steps

### 1. Run /commit Workflow
Execute the full `/commit` workflow first:
- Branch protection check
- Self-evolve sync
- Validate all skills
- Update CHANGELOG
- Create commit

If no pending changes, skip to step 3.

### 2. Final Install
After successful commit, run full installation to verify everything works:
```bash
sudo make install
```

This ensures:
- All skills are properly installed
- CLI is built and linked
- Completions are in place

### 3. Push to Remote
```bash
git push origin HEAD
```

### 4. Summary Report
Report what was pushed:
- Branch name
- Commits pushed (count)
- Remote URL
- Any CI/CD links if applicable
