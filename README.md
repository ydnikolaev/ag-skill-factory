# Antigravity Factory 🚀

> **Build Complete Agent Infrastructure.**
> A framework for managing AI agent blueprints: skills, workflows, team rules, and development standards.

[![Antigravity](https://img.shields.io/badge/Antigravity-Native-purple)](https://antigravity.google)
[![Go](https://img.shields.io/badge/Go-1.25-00ADD8)](https://go.dev)
[![Agent Skills](https://img.shields.io/badge/Agent-Skills-blue)](https://github.com/anthropics/skills)

## What is this?

**Antigravity Factory** is a blueprint management system for AI agents. It provides:
- **20 Expert Skills** — from backend-go-expert to mcp-expert
- **Shared Standards** — TDD, Git, Tech Debt protocols
- **Team Structure** — TEAM.md roster and PIPELINE.md workflow
- **Factory Skills** — meta-skills for creating and maintaining the ecosystem

Unlike simple scaffolding scripts, this tool enforces a **Design-First Philosophy**:
1.  **Context-Optimized**: Enforces concise `SKILL.md` (<500 lines) to respect context windows.
2.  **IDE-Aware**: Generates skills that understand absolute paths, `task_boundary`, and local environments.
3.  **Self-Verifying**: Includes built-in QA checklists for agents to validate their own work.

## ✨ Features

-   **🧠 20 Expert Skills**: Backend, Frontend, DevOps, QA, MCP, CLI, TUI, and more
-   **🛡️ Strict Validation**: `validate_skill.py` enforces <500 lines and quality standards
-   **✅ Auto-Checklists**: Each skill has `checklist.md` for QA
-   **🛠️ Factory CLI**: Go-based `factory install` and `factory list`
-   **📝 Standards Library**: TDD, Git, Tech Debt, Traceability protocols
-   **📦 Blueprint Pattern**: Copy entire `.agent/` structure to any project
-   **🏗️ Architecture Tests**: Enforces Go Modern standards via AST analysis

### Factory Skills (Meta-Tooling)

| Skill | Purpose |
|-------|---------|
| `@skill-creator` | Creates new skills from specs |
| `@skill-factory-expert` | Knows the factory codebase, answers questions |
| `@skill-interviewer` | Creative partner for skill ideation |
| `@skill-updater` | Mass updates existing skills |
| `@workflow-creator` | Designs automation workflows |

## 📂 Repository Structure

```
antigravity-factory/
├── .agent/                      # 🏭 Factory-internal (NOT copied to projects)
│   ├── skills/
│   │   ├── skill-creator/       # Meta-skill that creates other skills
│   │   ├── skill-factory-expert/# Project expert
│   │   ├── skill-interviewer/   # Creative partner for skill ideation
│   │   ├── skill-updater/       # Mass updates to existing skills
│   │   └── workflow-creator/    # Designs automation workflows
│   └── workflows/
│       ├── commit.md            # Pre-commit checks
│       ├── push.md              # Merge + changelog + push
│       └── self-evolve.md       # Factory synchronization
│
├── blueprint/                   # 📦 Copied to .agent/ on install
│   ├── skills/                  # 20 expert skills
│   ├── workflows/               # doc-cleanup, refactor
│   ├── rules/                   # TEAM.md, PIPELINE.md
│   ├── standards/               # TDD, Git, Tech Debt protocols
│   ├── _meta/                   # Presets config (presets.yaml)
│   └── private/                 # Private skills (gitignored)
│
├── website/                     # 📚 VitePress skill catalog
│   ├── .vitepress/config.mts
│   ├── index.md
│   └── skills/                  # Generated skill pages
│
├── cmd/factory/                 # 🔧 CLI source code
│   ├── root.go, install.go, list.go
│   ├── doctor.go                # factory doctor
│   └── version.go
│
├── internal/
│   ├── installer/               # Blueprint copy logic
│   ├── doctor/                  # Link checker, diagnostics
│   ├── presets/                 # Preset loader
│   └── config/                  # Config handler
│
├── scripts/
│   └── generate_catalog.py      # Generate website skill pages
│
├── .github/workflows/
│   └── deploy-docs.yml          # Auto-deploy VitePress to GitHub Pages
│
├── Makefile
└── README.md
```

## 🔧 Factory CLI

The `factory` CLI copies the blueprint to any project's `.agent/` folder.

### Installation

**Quick install (requires Go 1.22+):**
```bash
go install github.com/ydnikolaev/antigravity-factory@latest
```

**Or build from source:**
```bash
git clone https://github.com/ydnikolaev/antigravity-factory.git
cd antigravity-factory
make install
```

### Commands

```bash
factory install              # Interactive preset selection (TUI)
factory install --preset=backend  # Install specific preset
factory doctor               # Check for broken links and issues
factory list                 # Show installed inventory
factory version              # Show version
```

### Presets

| Preset | Skills | Description |
|--------|--------|-------------|
| `all` | 20 | Full blueprint |
| `core` | 5 | Pipeline essentials |
| `backend` | 9 | Go backend |
| `frontend` | 8 | Nuxt/Vue |
| `fullstack` | 12 | Backend + Frontend |
| `tma` | 8 | Telegram Mini Apps |
| `cli` | 8 | CLI/TUI apps |
| `minimal` | 2 | Utilities only |

### Example Workflow

```bash
# 1. Go to your project
cd my-project

# 2. Install blueprint (interactive)
factory install
# Select preset via TUI or use --preset

# 3. Check inventory
factory list
# 📦 Installed Blueprint
# Skills (20)
```

### Configuration

Config file: `~/.config/factory/config.yaml`

```yaml
source: ~/Developer/antigravity/antigravity-factory/blueprint
```

## 🧪 Development

```bash
# Run linter (FASCIST MODE)
make lint

# Run all tests
make test

# Build CLI
make build-factory

# Full install (build + install + completions)
make install
```

### Architecture Enforcement

The project includes `architecture_test.go` that enforces Go Modern standards:

| Rule | Enforcement |
|------|-------------|
| `NO_ANY` | Forbid `interface{}`/`any` — use generics |
| `NO_LEGACY_LOG` | Forbid `log` package — use `log/slog` |
| `NO_FMT_PRINT` | Forbid `fmt.Print*` in library code |
| `NO_GLOBALS` | Forbid exported mutable globals |
| `MODERN_ITER` | Enforce Go 1.22+ range syntax |
| `CTX_HYGIENE` | `context.Context` must be first param |
| `DOC_GO` | Every package must have `doc.go` |

## 🔧 Makefile Commands

| Command | Description |
|---------|-------------|
| `make install` | Build CLI, install to PATH, add completions |
| `make build-factory` | Build CLI binary to `bin/factory` |
| `make install-factory` | Symlink binary to `/usr/local/bin` |
| `make install-completions` | Add shell completions for zsh/bash |
| `make lint` | Run golangci-lint (FASCIST MODE) |
| `make test` | Run all tests |
| `make clean` | Remove build artifacts |
| `make changelog` | Generate CHANGELOG.md via git-cliff |
| `make validate SKILL=<name>` | Validate a single skill |
| `make validate-all` | Validate all skills in blueprint/ |
| `make generate-team` | Regenerate TEAM.md from skills |

## Blueprint Contents

| Folder | Contents |
|--------|----------|
| `skills/` | 20 expert agents (backend-go, frontend-nuxt, mcp, etc.) |
| `workflows/` | doc-cleanup, refactor |
| `rules/` | TEAM.md, PIPELINE.md |
| `standards/` | TDD, GIT, TECH_DEBT, TRACEABILITY, DOCUMENT_STRUCTURE |

## Artifact Persistence Rule (Dual-Write)

Antigravity artifacts (`brain/...`) are ephemeral. Documents (`docs/...`) are permanent.

**Rule**: Every skill MUST save its final output (e.g., `discovery-brief.md`) to the `project/docs/` directory and **change its status to "Approved"** BEFORE handing off to the next skill.

## 📖 How It Works

1.  **Design First**: Before creating a skill, answer: What triggers it? What's the decision tree?
2.  **Scaffold**: Use `@skill-creator` to create the standard structure.
3.  **Refine**: Edit `SKILL.md` with your logic and workflows.
4.  **Validate**: Run `make validate SKILL=<name>` to check quality.
5.  **Install**: Run `factory install` in your project.

## 🤝 Contributing

We welcome contributions! Please follow the **Design-First** philosophy:
1.  Keep instructions concise (<500 lines).
2.  Move details to `resources/` or `examples/`.
3.  Always include Team Collaboration and When to Delegate sections.
4.  Customize the checklist for your skill's domain.
5.  Run `make lint && make test` before committing.

## License

MIT
