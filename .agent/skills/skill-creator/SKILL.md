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

## Repository Structure

This repository is a **Skill Factory** that separates the **creator** from the **products**:

```
ag-skill-factory/
├── .agent/skills/         # Factory tooling (only skill-creator)
│   └── skill-creator/     # This skill - creates other skills
├── squads/                # Generated skills live here
│   ├── backend-go-expert/
│   ├── frontend-nuxt/
│   └── ...
└── Makefile               # Links skills to global brain
```

> **IMPORTANT**: All new skills MUST be created in `squads/`, NOT in `.agent/skills/` or directly in the global brain.

## Skill Creation Workflow

### Phase 1: Design (Mental Step)
Before running any script, answer these questions:
1.  **What is the Trigger?** What user intent triggers this skill? (Write this in the Description).
2.  **What is the Decision Tree?** Does the skill have one path or multiple? (If multiple, plan a "Decision Tree" section).
3.  **What are the Resources?** Do I need a Python script for logic? A Markdown template for output?

### Phase 2: Scaffold
Run the initialization script to create the standard structure in `squads/`:

```bash
# ALWAYS create skills in squads/ directory
python3 .agent/skills/skill-creator/scripts/init_skill.py <skill-name>
```

This creates:
```
squads/<skill-name>/
├── SKILL.md          # The "Brain": Logic, Decisions, Workflow.
├── scripts/          # The "Hands": Python/Bash scripts for execution.
├── resources/        # The "Tools": Templates, Configs (e.g., Dockerfile).
├── examples/         # The "Demos": Usage examples.
└── references/       # The "Library": Documentation, Cheatsheets.
```

### Phase 3: Refine
Edit the generated files:
1.  **Fill the Decision Tree**: Map user inputs to actions.
2.  **Write the Workflow**: Step-by-step instructions.
3.  **Clean up**: Remove unused sections from the template.
4.  **Adapt the checklist**: Rewrite `references/checklist.md` for this skill's domain.
    - The template checklist is generic — make it specific!
    - Example: For an MCP skill, add checks for "stdio transport", "tool descriptions", etc.

### Phase 4: Verify
Run the validation script:

```bash
make validate SKILL=<skill-name>
```

This checks:
- Frontmatter, length (<500 lines)
- Team Collaboration & When to Delegate sections
- Customized checklist in `references/`
- No large embedded code blocks

Fix all errors before installing.

### Phase 5: Install (Link to Global Brain)
After creating and refining the skill, link it to the global brain:

```bash
make install    # Validates all, generates TEAM.md, then links
```

This copies skills from `squads/<skill-name>/` → `~/.gemini/antigravity/global_skills/<skill-name>/`.

## Makefile Commands

| Command | Description |
|---------|-------------|
| `make install` | Validate all, generate TEAM.md, link all skills |
| `make validate SKILL=<name>` | Validate a single skill (detailed output) |
| `make validate-all` | Validate all skills |
| `make generate-team` | Regenerate `squads/TEAM.md` |
| `make install-factory` | Install only skill-creator |
| `make install-squads` | Install only skills from `squads/` |
| `make uninstall` | Remove all symlinks |

## Anti-Patterns

❌ **DO NOT** create skills directly in `~/.gemini/antigravity/global_skills/`
❌ **DO NOT** use the `--global` flag in init_skill.py
❌ **DO NOT** place new skills in `.agent/skills/` (reserved for factory tooling)

✅ **DO** create skills in `squads/`
✅ **DO** use `make install-squads` to link them globally

## Content Organization Rules

**SKILL.md should contain:**
- Decisions, workflows, and logic (the "brain")
- Brief inline examples (max 5-10 lines)
- References to detailed examples: `See examples/python-server.py`

**examples/ should contain:**
- Full working code examples (complete files)
- Configuration samples
- Complete templates

**references/ should contain:**
- Cheatsheets and quick reference guides
- External documentation links
- Troubleshooting guides

**resources/ should contain:**
- Templates for generation
- Config files
- Static assets

> **CRITICAL**: Do NOT embed large code blocks (>10 lines) in SKILL.md. 
> Instead, create files in `examples/` and reference them.

## Team Collaboration

Every skill should know about the team. Add these sections to SKILL.md:

**Team Collaboration** — List related skills by role:
```markdown
## Team Collaboration
- **<Role>**: `@<skill-name>` (Brief description of collaboration)
- **<Role>**: `@<skill-name>` (Brief description of collaboration)
```

**When to Delegate** — When to hand off to another skill:
```markdown
## When to Delegate
- ✅ **Delegate to `@<skill-name>`** when: <Condition for handoff>
- ⬅️ **Return to `@<skill-name>`** if: <Condition to return>
```

> **Team Reference**: See `squads/TEAM.md` for the full list of available skills.

