---
description: Reload active skill to refresh agent context when instructions drift
---

# /reload-skill Workflow

Reloads the SKILL.md of the current or specified skill to restore agent behavior when context drifts.

## Usage

```
/reload-skill                    # Reload currently active skill
/reload-skill @skill-name        # Reload a specific skill
```

## Steps

// turbo-all

### 1. Identify the Skill

**If explicitly specified** (`@skill-name`): use that skill.

**If not specified**: detect from context:
- Check the last `view_file` call for any `SKILL.md`
- Or identify by current activity (which skill was mentioned)

### 2. Locate SKILL.md

```bash
# Path: .agent/skills/<skill-name>/SKILL.md
ls .agent/skills/*/SKILL.md
```

### 3. Re-read SKILL.md

Use `view_file` on the full path to SKILL.md:
```
view_file: .agent/skills/<skill-name>/SKILL.md
```

### 4. Confirm Reload

After reading the file, the agent must:
1. Output brief confirmation: "✅ Skill `<skill-name>` reloaded"
2. Remind itself of the key points from the skill
3. Continue working according to the instructions
