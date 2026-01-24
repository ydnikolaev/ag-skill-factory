---
description: Self-evolution check for ag-skill-factory project synchronization
---

# /self-evolve Workflow

Activate `@skill-factory-expert` and verify project consistency.

## Steps

// turbo-all

### 1. Check Current State
```bash
# Count skills
ls squads/ | grep -v -E "\.md$|^_|^references$" | wc -l

# List Makefile targets
grep -E "^[a-z].*:" Makefile | head -20

# List scripts
ls .agent/skills/skill-creator/scripts/
```

### 2. Verify skill-factory-expert
Compare output from Step 1 against:
- `## Current Squad Roster (N skills)` — count must match
- `## Makefile Commands` table — all targets documented
- `## Key Files` → `### Scripts` — all scripts listed

**If mismatch → Update `.agent/skills/skill-factory-expert/SKILL.md`**

### 3. Verify skill-creator (Younger Brother 👶)
Check `.agent/skills/skill-creator/SKILL.md`:
- Uses `project/docs/` (not `docs/`)
- Template in `resources/templates/SKILL.md` is up-to-date

**If mismatch → Update skill-creator**

### 4. Verify PIPELINE.md
Check `squads/PIPELINE.md`:
- All skills in Core Pipeline diagram
- Handoff Matrix has correct `project/docs/...` paths
- Return Paths table is complete

### 5. Verify TEAM.md
```bash
make generate-team
git diff squads/TEAM.md
```
If diff → TEAM.md was outdated, now fixed.

### 6. Validate All Skills
```bash
make validate-all
```
All must pass ✅

### 7. Install If Changes
```bash
sudo make install
```

### 8. Report
Summarize what was checked and any updates made.
