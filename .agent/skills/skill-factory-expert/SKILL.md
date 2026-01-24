---
name: skill-factory-expert
description: Expert on the ag-skill-factory project. Knows the entire codebase, skill creation workflow, Makefile commands, and project architecture. Activate this skill when working within the skill-factory repository.
---

# Skill Factory Expert

This skill has deep knowledge of the **ag-skill-factory** project — the Skill Factory that produces high-quality Antigravity agent skills.

> [!IMPORTANT]
> ## Role Boundary
> **I KNOW the factory. I DON'T CREATE skills.**
> 
> | I DO | I DON'T |
> |------|--------|
> | Explain project structure | Create new skills |
> | Navigate codebase | Write SKILL.md files |
> | Answer "how does X work?" | Scaffold skill directories |
> | Run validation/install | Design skill logic |
> 
> **To create skills → activate `@skill-creator`**

> [!CAUTION]
> ## 🚨 MANDATORY SELF-EVOLUTION PROTOCOL
> 
> **BEFORE completing ANY task in this repository, you MUST:**
> 
> 1. **Check if repository structure changed** (new skills, modified Makefile, updated scripts)
> 2. **If changes detected → UPDATE THIS FILE** before finishing your task
> 3. **Failure to update = INCOMPLETE TASK** — the user will reject your work
> 
> **Self-Evolution Triggers (check EVERY TIME):**
> - [ ] New skill added to `squads/`? → Update "Current Squad Roster" table
> - [ ] Makefile commands changed? → Update "Makefile Commands" table
> - [ ] New scripts in `skill-creator/scripts/`? → Update "Key Files" section
> - [ ] New factory skills in `.agent/skills/`? → Update "Project Overview" diagram
> - [ ] Workflow process changed? → Update "Skill Creation Workflow" section
> - [ ] **skill-creator outdated?** → Update `.agent/skills/skill-creator/SKILL.md` too! (younger brother 👶)
> 
> **How to Self-Evolve:**
> ```bash
> # 1. List current squads
> ls squads/
> 
> # 2. Check for new make targets
> grep -E "^[a-z].*:" Makefile | head -20
> 
> # 3. Update this file AND skill-creator if needed
> # Edit: .agent/skills/skill-factory-expert/SKILL.md
> # Edit: .agent/skills/skill-creator/SKILL.md
> ```
> 
> ⚠️ **If you skip this step, the skill becomes outdated and USELESS.**

## Project Overview

**Purpose**: The Skill Factory separates the **creator** (skill-creator) from the **products** (skills in `squads/`).

```
ag-skill-factory/
├── .agent/skills/           # 🏭 Factory tooling
│   ├── skill-creator/       # Meta-skill that creates other skills
│   ├── skill-factory-expert/# THIS SKILL - project expert
│   ├── skill-interviewer/   # Creative partner for skill ideation
│   ├── skill-updater/       # Mass updates to existing skills
│   └── workflow-creator/    # Designs automation workflows
├── .agent/workflows/        # 🔄 Automation workflows
│   ├── commit.md            # Pre-commit checks + changelog
│   ├── push.md              # Merge + push pipeline
│   └── self-evolve.md       # Factory synchronization
├── squads/                  # 👥 Generated skills (gitignored)
│   ├── backend-go-expert/
│   ├── frontend-nuxt/
│   ├── mcp-expert/
│   └── ...
├── Makefile                 # Links skills to global brain
├── ARTIFACT_REGISTRY.md                # 📋 AI Agent Context (READ THIS FIRST)
├── README.md                # Project documentation
└── docs/                    # Documentation
```

> [!TIP]
> **Always read `ARTIFACT_REGISTRY.md` first!** It contains full project context, CLI commands, and development rules.

## Core Concepts

### 1. Skill Categories
- **Factory Skills** (`.agent/skills/`): `skill-creator`, `skill-factory-expert`, `skill-interviewer`, `workflow-creator` — internal tooling
- **Squad Skills** (`squads/`): All generated skills live here, gitignored, user-specific
- **Workflows** (`.agent/workflows/`): Automation scripts — `/commit`, `/push`, `/self-evolve`

### 2. Skill Structure
Every skill follows this pattern:
```
<skill-name>/
├── SKILL.md          # Brain: Logic, Decisions, Workflow (<500 lines)
├── scripts/          # Hands: Python/Bash for execution
├── resources/        # Tools: Templates, Configs
├── examples/         # Demos: Usage examples
└── references/       # Library: Docs, Cheatsheets
```

### 3. Design Philosophy
1. **Concise is Key**: SKILL.md must be under 500 lines
2. **Progressive Disclosure**: Metadata → SKILL.md → Scripts/References
3. **IDE Awareness**: Absolute paths, `task_boundary` for long tasks
4. **Dual-Write Pattern**: Drafts in `brain/`, finals in `project/docs/`

### 4. Project Docs Convention

> [!CAUTION]
> **All skills MUST use `project/docs/` NOT `docs/`!**
> 
> - ✅ Correct: `project/docs/features/`, `project/docs/architecture/`
> - ❌ Wrong: `docs/features/`, `docs/architecture/`
> 
> **Why?** User projects have their docs in `project/docs/` folder. Using `docs/` creates files in wrong location.

### 5. Project Config Awareness
Every skill must check `project/CONFIG.yaml` before making technical decisions:

```yaml
# Key sections in CONFIG.yaml:
stack:
  backend:
    framework: go-stdlib  # NOT chi, NOT fiber!
    modules: [pgx/v5, river, ...]
  frontend:
    framework: nuxt4
runtime:
  go: "1.25"
  node: "22"
architecture:
  style: ddd
```

**Why this matters:** Prevents skills from suggesting wrong libraries (e.g., Chi instead of stdlib).

## Makefile Commands

| Command | Description |
|---------|-------------|
| `make install` | **Full install**: validate-all → generate-team → install-factory → install-squads → build CLI |
| `make validate SKILL=<name>` | Validate a single skill (detailed output with errors) |
| `make validate-all` | Validate all skills in `squads/` (summary output) |
| `make generate-team` | Regenerate `squads/TEAM.md` from skill descriptions |
| `make install-factory` | Copy `.agent/skills/*` → `~/.gemini/antigravity/global_skills/` |
| `make install-squads` | Copy `squads/*` → `~/.gemini/antigravity/global_skills/` |
| `make build-skills` | Build the `skills` CLI binary |
| `make install-skills` | Install CLI to `/usr/local/bin/skills` |
| `make test` | Run all Go tests |
| `make lint` | Run linters |
| `make uninstall` | Remove all installed skills from global brain |

## Skills CLI

The factory includes a Go CLI (`skills`) for managing skills in workspaces:

```bash
skills install      # Install all skills to .agent/skills/
skills update       # Update from factory (shows diff)
skills backport X   # Push local changes back to factory
skills list         # Show installed vs available skills
```

> [!NOTE]
> CLI reads config from `~/.config/ag-skills/config.yaml`

## Skill Creation Workflow

### Phase 1: Design
Before creating a skill, answer:
1. **What is the Trigger?** What user intent activates this skill?
2. **What is the Decision Tree?** Single path or multiple?
3. **What Resources are needed?** Scripts, templates, references?

### Phase 2: Scaffold
```bash
python3 .agent/skills/skill-creator/scripts/init_skill.py <skill-name>
```
Creates the standard structure in `squads/<skill-name>/`.

### Phase 3: Refine
1. Fill the **Decision Tree** in SKILL.md
2. Write the **Workflow** with clear phases

3. **Adapt the checklist** in `references/checklist.md`
4. Move large code examples to `examples/`
5. **Enforce Handoff Protocol**: Ensure the "Draft -> Approved" status change step is present


### Phase 4: Verify
```bash
make validate SKILL=<skill-name>
```
Checks: frontmatter, length (<500), team sections, checklist customization.

### Phase 5: Install
```bash
make install-squads   # Copy squads to global brain
# OR
make install          # Full pipeline: validate → generate-team → install all
```
Physically copies `squads/<skill-name>/` → `~/.gemini/antigravity/global_skills/<skill-name>/`

## Key Files

### Validation Script
`Path: .agent/skills/skill-creator/scripts/validate_skill.py`
- Validates SKILL.md frontmatter
- Checks line count (<500)
- Ensures Team Collaboration & When to Delegate sections exist
- Verifies checklist in references/

### Scripts
`Path: .agent/skills/skill-creator/scripts/`
- `init_skill.py` — Creates skill skeleton in `squads/`
- `validate_skill.py` — Validates SKILL.md against standards
- `add_config_awareness.py` — Adds CONFIG.yaml awareness to skills
- `add_mcp_awareness.py` — Adds MCP context awareness to skills

### Templates
- `Path: .agent/skills/skill-creator/resources/templates/SKILL.md` — Base template
- `Path: .agent/skills/skill-creator/resources/references/checklist.md` — QA checklist

### Standards (Shared Protocols)
- `Path: squads/_standards/TDD_PROTOCOL.md` — Test-Driven Development rules
- `Path: squads/_standards/GIT_PROTOCOL.md` — Git workflow and Conventional Commits
- `Path: squads/_standards/TECH_DEBT_PROTOCOL.md` — TODO/workaround tracking
- `Path: squads/_standards/TRACEABILITY_PROTOCOL.md` — Pipeline requirements tracing

### Team Registry
`Path: squads/TEAM.md`
- Auto-generated via `make generate-team`
- Lists all skills with descriptions

## Anti-Patterns

❌ **NEVER** create skills directly in `~/.gemini/antigravity/global_skills/`
❌ **NEVER** use `--global` flag (removed from init_skill.py)
❌ **NEVER** place new skills in `.agent/skills/` (reserved for factory tooling)
❌ **NEVER** embed large code blocks (>10 lines) in SKILL.md


✅ **ALWAYS** create skills in `squads/`
✅ **ALWAYS** use `make install-squads` to link
✅ **ALWAYS** customize checklist for the skill's domain
✅ **ALWAYS** add Team Collaboration and When to Delegate sections
✅ **ALWAYS** enforce "Draft -> Approved" status change before handoff

## Content Organization

**SKILL.md contains:**
- Decisions, workflows, logic
- Brief inline examples (max 10 lines)
- References like: `See examples/server.py`

**examples/ contains:**
- Full working code
- Configuration samples

**references/ contains:**
- Cheatsheets
- External docs
- Troubleshooting guides

**resources/ contains:**
- Generation templates
- Config files

## Team Collaboration

Skills must include team awareness:

```markdown
## Team Collaboration
- **Role**: `@skill-name` (Description of collaboration)

## When to Delegate
- ✅ **Delegate to `@skill-name`** when: <condition>
- ⬅️ **Return to `@skill-name`** if: <condition>
```

See `squads/TEAM.md` for the full roster.

## Core Pipeline

The Discovery-to-Delivery pipeline flows through these phases:

```
idea-interview → product-analyst → bmad-architect → tech-spec-writer → implementation → delivery
```

**Core path:** 5 mandatory skills
**Optional paths:**
- TMA/Bot: `@telegram-mechanic`, `@tma-expert`
- CLI: `@cli-architect`, `@tui-charm-expert`  
- Design: `@ux-designer`, `@ui-implementor`
- Hosting: `@timeweb-sysadmin`
- **Utility (any phase):** `@project-bro`, `@debugger`

See `squads/PIPELINE.md` for visual diagram and handoff matrix.

## Current Squad Roster (20 skills)

| Skill | Focus |
|-------|-------|
| `idea-interview` | Discovery phase, extract project info from user |
| `feature-fit` | Analyzes new features for EXISTING projects |
| `product-analyst` | Vision, Roadmap, Specs (merged PM+SA) |
| `tech-spec-writer` | Converts architecture into detailed tech specs |
| `bmad-architect` | DDD, Context Maps, API Contracts |
| `backend-go-expert` | Go 1.25+, Clean Architecture, DDD |
| `frontend-nuxt` | Nuxt 4, TailwindCSS, SSR |
| `tma-expert` | Telegram Mini Apps |
| `telegram-mechanic` | Bot API, Webhooks, initData (optional) |
| `cli-architect` | Cobra, Viper, POSIX CLI (optional) |
| `tui-charm-expert` | BubbleTea, Lipgloss (optional) |
| `mcp-expert` | MCP servers (Go preferred) |
| `ux-designer` | Design systems, tokens (optional) |
| `ui-implementor` | Tailwind, shadcn/ui (optional) |
| `qa-lead` | E2E, API, UI testing, severity levels |
| `devops-sre` | Docker, CI/CD, deployments |
| `timeweb-sysadmin` | Timeweb Cloud, VPS (optional) |
| `project-bro` | Project awareness, "where are we?" (utility) |
| `refactor-architect` | Analyzes codebase, designs modular refactoring specs |
| `debugger` | Systematic 7-step bug investigation (utility) |

## When to Delegate

- ✅ **Delegate to `@skill-creator`** when: Creating a new skill
- ⬅️ **Return from `@skill-creator`** after: Skill is scaffolded
- 🤝 **Coordinate with squad skills** when: Understanding their capabilities

## Antigravity Best Practices

- Use `task_boundary` when performing multi-step operations
- Use `notify_user` for user review checkpoints
- Always use **absolute paths** in scripts and documentation
