---
description: Run full commit workflow, install, and push to remote
---

# /push Workflow

Complete push pipeline: commit + merge to main + cleanup.

// turbo-all

## Steps

### 1. Run /commit Workflow
Execute the full `/commit` workflow first:
- Branch protection (auto-creates feature branch if on main)
- Self-evolve sync checks
- Validate all skills
- Create commit

If no pending changes, skip to step 2.

### 2. Final Install
After successful commit, run full installation to verify everything works:
```bash
sudo make install
```

### 3. Merge to Main
```bash
git checkout main
git merge <feature-branch> --no-edit
```
Fast-forward merge preferred.

### 4. Cleanup Branches
```bash
# Delete local feature branch
git branch -d <feature-branch>

# Delete remote feature branch (if exists)
git push origin --delete <feature-branch>
```

### 5. Update Changelog
```bash
make changelog
git add CHANGELOG.md && git commit -m "chore: update changelog" || true
```
Generates CHANGELOG.md from git history using git-cliff.

### 6. Push to Main
```bash
git push origin main
```

### 7. Summary Report
Report what was pushed:
- Branch merged
- Commits count
- Remote branch cleanup status
