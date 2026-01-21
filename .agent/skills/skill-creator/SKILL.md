---
name: skill-creator
description: Defines the philosophy and process for creating high-quality Antigravity skills. Use this skill when asked to create a new skill or standardize an existing one.
---

# Skill Creator: The Design System

This skill uses a "Design First, Code Second" approach. Your goal is not just to create a folder, but to design an effective tool for an autonomous agent.

## Core Philosophy

1.  **Concise is Key**: The context window is precious. Do not overload `SKILL.md`.
    -   **Limit**: Keep `SKILL.md` under 500 lines.
    -   **Strategy**: Move large prompt chains, examples, or schemas to `resources/` or `references/`.
2.  **Progressive Disclosure**:
    -   **Level 1**: Metadata (Name/Description) - Agent sees this first.
    -   **Level 2**: `SKILL.md` - Agent reads this only when activated.
    -   **Level 3**: Scripts/References - Agent reads/runs these only when needed.
3.  **IDE Awareness**:
    -   Always use absolute paths.
    -   Use `task_boundary` for long-running workflows.
    -   Assume the agent knows nothing about the file structure until it runs `ls`.

## Skill Creation Workflow

### Phase 1: design (Mental Step)
Before running any script, answer these questions:
1.  **What is the Trigger?** What user intent triggers this skill? (Write this in the Description).
2.  **What is the Decision Tree?** Does the skill have one path or multiple? (If multiple, plan a "Decision Tree" section).
3.  **What are the Resources?** Do I need a Python script for logic? A Markdown template for output?

### Phase 2: Scaffold
Run the initialization script to create the standard structure. This script pulls a high-quality template.

```bash
python3 .agent/skills/skill-creator/scripts/init_skill.py <skill-name>
# Add --global to install to ~/.gemini/antigravity/skills/
```

### Phase 3: Refine
Edit the generated `SKILL.md`.
1.  **Fill the Decision Tree**: Map user inputs to actions.
2.  **Write the Workflow**: Step-by-step instructions.
3.  **Clean up**: Remove unused sections from the template.

### Phase 4: Verify
1.  Read `.agent/skills/<skill-name>/references/checklist.md`.
2.  Verify your skill against every item in that checklist.

## Anatomy of a Skill
```
skill-name/
├── SKILL.md          # The "Brain": Logic, Decisions, Workflow.
├── scripts/          # The "Hands": Python/Bash scripts for execution.
├── resources/        # The "Tools": Templates, Configs (e.g., Dockerfile).
└── references/       # The "Library": Documentation, Cheatsheets.
```
