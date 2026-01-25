---
description: Pre-commit checks, changelog update, and smart commit for ag-skill-factory
---

# /commit Workflow

Full pre-commit pipeline: sync docs, validate, update changelog, commit.

// turbo-all

## Steps

### 1. Branch Protection Check
```bash
current_branch=$(git branch --show-current)
if [ "$current_branch" = "main" ]; then
  echo "⚠️ On main branch - creating feature branch..."
  # Generate branch name from current work context
  git checkout -b feat/$(date +%Y%m%d)-update
fi
echo "✅ On branch: $(git branch --show-current)"
```

If on `main`, automatically create a feature branch with date-based name.
Agent should use a descriptive name based on session context (e.g., `feat/meta-skills`, `fix/pipeline-paths`).

### 2. Self-Evolve Sync (Embedded)
Critical checks from `/self-evolve` workflow:

#### 2.1 Check Skill Count
```bash
# Count actual skills
ls blueprint/skills/ | grep -v -E "\\.md$|^_|^references$" | wc -l
```
Compare against `skill-factory-expert/SKILL.md` roster count. If mismatch → update expert.

#### 2.2 Regenerate TEAM.md
```bash
make generate-team
git diff blueprint/skills/TEAM.md
```
If diff exists → TEAM.md was outdated, now fixed.

#### 2.3 Verify project/docs/ Convention
Spot-check skill-creator templates use `project/docs/` not `docs/`.

#### 2.4 Validate All Skills
```bash
make validate-all
```
All must pass ✅

### 3. Diff Analysis
```bash
git diff --stat
git diff --name-only
```
Study the diff to understand what changed.

### 4. Generate Commit Message
Create a Conventional Commits message based on the diff:

**Format:** `<type>(<scope>): <description>`

**Types:**
| Type | Use for |
|------|---------|
| `feat` | New feature or skill |
| `fix` | Bug fix |
| `docs` | Documentation only |
| `refactor` | Code restructure |
| `chore` | Maintenance tasks |

**Examples:**
- `feat(skills): add workflow-creator meta-skill`
- `fix(pipeline): correct handoff paths in PIPELINE.md`
- `docs(readme): update installation instructions`

### 5. Stage and Commit
```bash
git add -A
git commit -m "<generated-message>"
```

### 6. Summary Report
Report what was committed:
- Files changed
- Commit hash
