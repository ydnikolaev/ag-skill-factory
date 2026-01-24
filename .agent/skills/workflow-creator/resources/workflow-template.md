# Workflow Template

Use this template when creating new workflows.

---

```markdown
---
description: [One-line description of what this workflow does]
---

# /[command-name] Workflow

[Brief purpose description - what problem does this solve?]

## Steps

// turbo-all

### 1. [Context Loading / Check State]
```bash
# Commands to understand current state
ls -la [relevant-path]
```

### 2. [Main Action]
[Instructions or commands for main task]

```bash
# If applicable
[commands]
```

### 3. [Verification]
[How to verify the action was successful]

```bash
# Verification commands
[commands]
```

### 4. [Report]
Summarize what was checked and any updates made.
```

---

## Notes

- Replace `[command-name]` with actual slash command (e.g., `self-evolve`)
- Use `// turbo-all` for fully autonomous workflows
- Use `// turbo` above individual steps for selective auto-run
- Always end with a Report step
