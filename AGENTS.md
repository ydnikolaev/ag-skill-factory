# AGENTS.md - AI Agent Context

> This file provides context for AI agents working on this project.
> Read this first before making any changes.

## Project Overview

**ag-skill-factory** is a skill management system for Antigravity AI agents.

### What it does:
1. **Defines skills** - Markdown-based agent skill definitions in `squads/`
2. **Validates skills** - Python validator ensures skills meet standards
3. **Installs skills** - Go CLI (`skills`) deploys to workspaces
4. **Syncs skills** - Bidirectional sync between factory and projects

---

## Project Structure

```
ag-skill-factory/
├── squads/                    # 🧠 SKILL DEFINITIONS
│   ├── backend-go-expert/     # Each skill has SKILL.md + resources
│   ├── frontend-nuxt/
│   ├── tech-spec-writer/
│   ├── _standards/            # Shared protocols (TDD, Git)
│   ├── TEAM.md                # Auto-generated skill roster
│   └── PIPELINE.md            # Visual workflow diagram
│
├── cmd/skills/                # 🔧 CLI COMMANDS
│   ├── root.go                # Main command setup
│   ├── install.go             # skills install
│   ├── update.go              # skills update
│   ├── backport.go            # skills backport <name>
│   └── list.go                # skills list
│
├── internal/                  # 📦 CORE LOGIC
│   ├── installer/             # Install/update/backport logic
│   ├── diff/                  # Directory comparison
│   └── coverage/              # Test coverage enforcement
│
├── .agent/skills/             # Local skills (dev)
├── docs/                      # Project documentation
│   ├── discovery/             # Discovery briefs
│   └── specs/                 # Technical specifications
│
└── Makefile                   # Build commands
```

---

## Key Commands

### Makefile (in this repo)
```bash
make install        # Validate all + install skills + build CLI
make validate-all   # Validate all skill definitions
make test           # Run all Go tests
make build-skills   # Build the skills CLI binary
```

### Skills CLI (in any workspace)
```bash
skills install      # Install all skills to .agent/skills/
skills update       # Update from factory (shows diff)
skills backport X   # Push local changes back to factory
skills list         # Show installed vs available skills
```

---

## Development Rules

### TDD Protocol
> [!CAUTION]
> **No code without failing test.**
> - Write test first
> - See it fail
> - Write implementation
> - See it pass

### Git Protocol
> [!CAUTION]
> **Use feature branches and Conventional Commits.**
> - `feat/...`, `fix/...`, `chore/...`
> - `feat(cli): add backport command`
> - Never push directly to main

### Test Coverage
The test `internal/coverage/coverage_test.go` will **FAIL** if:
- A Go package has no `_test.go` files
- Expected test files are missing

---

## Skill Format

Skills are defined in `squads/<skill-name>/SKILL.md`:

```yaml
---
name: skill-name
description: What the skill does
---

# Skill Name

## When to Activate
- Trigger conditions

## Workflow
1. Step 1
2. Step 2

## Team Collaboration
- `@other-skill` - How to collaborate

## Handoff Protocol
...
```

---

## Config

The CLI reads from `~/.config/ag-skills/config.yaml`:

```yaml
source: /path/to/ag-skill-factory/squads
global_path: ~/.gemini/antigravity/global_skills
```

---

## Important Files

| File | Purpose |
|------|---------|
| `squads/TEAM.md` | Auto-generated skill roster |
| `squads/PIPELINE.md` | Visual workflow diagram |
| `squads/_standards/` | Shared protocols (TDD, Git) |
| `go.mod` | Go module definition |
| `Makefile` | Build/test/install automation |

---

## When Working on This Project

1. **Adding a new skill**: Use `python3 .agent/skills/skill-creator/scripts/init_skill.py <name>`
2. **Adding a new CLI command**: Create `cmd/skills/<name>.go` + add tests
3. **Modifying installer**: Edit `internal/installer/installer.go` + update tests
4. **Checking coverage**: Run `make test` - will fail if tests missing

---

## Dependencies

- Go 1.23+
- Python 3.x (for skill validator)
- Cobra (CLI framework)
- Viper (config management)
