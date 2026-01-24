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

### 2. Self-Evolve Sync
Run the `/self-evolve` workflow to ensure documentation is synchronized:
- Skill count matches reality
- TEAM.md regenerated
- PIPELINE.md validated
- skill-factory-expert up to date

### 3. Validate All Skills
```bash
make validate-all
```

### 4. Diff Analysis
```bash
git diff --stat
git diff --name-only
```
Study the diff to understand what changed.

### 5. Update CHANGELOG.md
Based on the diff, add an entry to `CHANGELOG.md` under `## [Unreleased]`:

**Format (Keep a Changelog):**
```markdown
## [Unreleased]

### Added
- New feature description

### Changed  
- Modified feature description

### Fixed
- Bug fix description
```

**Categories:**
- `Added` — new features, skills, workflows
- `Changed` — modifications to existing functionality
- `Fixed` — bug fixes
- `Removed` — deleted features
- `Deprecated` — soon-to-be removed features

### 6. Generate Commit Message
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

### 7. Stage and Commit
```bash
git add -A
git commit -m "<generated-message>"
```

### 8. Summary Report
Report what was committed:
- Files changed
- CHANGELOG entry added
- Commit hash
