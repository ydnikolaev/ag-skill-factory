# AGENTS.md - AI Agent Context

> This file provides context for AI agents working on this project.
> Read this first before making any changes.

## Project Overview

**antigravity-factory** is a blueprint management system for Antigravity AI agents.

### What it does:
1. **Defines skills** - 20 expert agent skills in `blueprint/skills/`
2. **Provides workflows** - Automation scripts in `blueprint/workflows/`
3. **Sets team rules** - TEAM.md roster and PIPELINE.md in `blueprint/rules/`
4. **Shares standards** - TDD, Git, Tech Debt protocols in `blueprint/standards/`
5. **Installs blueprints** - Go CLI (`factory`) deploys to project `.agent/`

---

## Project Structure

```
antigravity-factory/
├── blueprint/                   # 📦 COPIED TO PROJECTS on install
│   ├── skills/                  # 20 expert skills
│   ├── workflows/               # doc-cleanup, refactor
│   ├── rules/                   # TEAM.md, PIPELINE.md
│   ├── standards/               # TDD, Git, Tech Debt protocols
│   ├── _meta/                   # Presets config
│   └── private/                 # Private skills (gitignored)
│
├── .agent/                      # 🏭 FACTORY-INTERNAL (NOT copied)
│   ├── skills/                  # skill-creator, skill-factory-expert, etc.
│   └── workflows/               # commit, push, self-evolve
│
├── website/                     # 📚 VitePress skill catalog
│
├── cmd/factory/                 # 🔧 CLI COMMANDS
│   ├── install.go               # factory install (--preset, TUI)
│   ├── doctor.go                # factory doctor
│   ├── list.go                  # factory list
│   └── version.go               # factory version
│
├── internal/
│   ├── installer/               # Blueprint copy logic
│   ├── doctor/                  # Link checker, diagnostics
│   ├── presets/                 # Preset loader (YAML)
│   └── config/                  # Config handler
│
├── scripts/                     # Python utilities
└── Makefile                     # Build commands
```

---

## Key Commands

### Makefile (in this repo)
```bash
make install        # Validate all + build CLI + install symlink
make validate-all   # Validate all skill definitions in blueprint/
make test           # Run all Go tests
make build-factory  # Build the factory CLI binary
make generate-team  # Regenerate TEAM.md from skills
```

### Factory CLI (in any workspace)
```bash
factory install              # Interactive preset selection (TUI)
factory install --preset=backend  # Install specific preset
factory doctor               # Check for broken links and issues
factory list                 # Show installed inventory
factory version              # Show version
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

---

## Skill Format

Skills are defined in `blueprint/skills/<skill-name>/SKILL.md`:

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

The CLI reads from `~/.config/factory/config.yaml`:

```yaml
source: ~/Developer/antigravity/antigravity-factory/blueprint
```

---

## Important Files

| File | Purpose |
|------|---------|
| `blueprint/rules/TEAM.md` | Auto-generated skill roster |
| `blueprint/rules/PIPELINE.md` | Visual workflow diagram |
| `blueprint/standards/` | Shared protocols (TDD, Git, etc.) |
| `go.mod` | Go module definition |
| `Makefile` | Build/test/install automation |

---

## When Working on This Project

### Blueprint Content
1. **Adding a skill**: Use `@skill-creator` or create in `blueprint/skills/<name>/`
2. **Adding a workflow**: Create `blueprint/workflows/<name>.md`
3. **Updating rules**: Edit `blueprint/rules/TEAM.md` or `PIPELINE.md`
4. **Adding a standard**: Create `blueprint/standards/<NAME>_PROTOCOL.md`

### Factory Tooling
5. **Adding a CLI command**: Create `cmd/factory/<name>.go` + add tests
6. **Modifying installer**: Edit `internal/installer/installer.go` + update tests
7. **Regenerating TEAM.md**: Run `make generate-team`
8. **Validating skills**: Run `make validate SKILL=<name>` or `make validate-all`

---

## Dependencies

- Go 1.25+
- Python 3.x (for skill validator)
- Cobra (CLI framework)
- Viper (config management)
- Afero (filesystem abstraction for testing)
