# Antigravity Skill Factory 🚀

> **Design High-Quality Autonomous Agent Skills.**
> A powerful framework for creating standardized, effective, and "native" skills for Antigravity and MCP-based agents.

[![Antigravity](https://img.shields.io/badge/Antigravity-Native-purple)](https://antigravity.google)
[![Go](https://img.shields.io/badge/Go-1.25-00ADD8)](https://go.dev)
[![Agent Skills](https://img.shields.io/badge/Agent-Skills-blue)](https://github.com/anthropics/skills)

## What is this?

**Skill Factory** is the home of the `skill-creator` — a meta-skill that empowers AI agents to create *other* high-quality skills.

Unlike simple scaffolding scripts, this tool enforces a **Design-First Philosophy**:
1.  **Context-Optimized**: Enforces concise `SKILL.md` (<500 lines) to respect context windows.
2.  **IDE-Aware**: Generates skills that understand absolute paths, `task_boundary`, and local environments.
3.  **Self-Verifying**: Includes built-in QA checklists for agents to validate their own work.

## ✨ Features

-   **🧠 Smart Templates**: Starts every skill with a "Decision Tree" and "Phased Workflow" structure.
-   **🛡️ Strict Validation**: `validate_skill.py` enforces the 500-line limit and checks for IDE-aware tool usage.
-   **✅ Auto-Checklists**: Generates `checklist.md` for quality assurance.
-   **🛠️ Skills CLI**: Go-based CLI for install, update, backport, and list operations.
-   **📚 Design Guide**: The `skill-creator` serves as a textbook for agents on *how* to design good tools.
-   **📦 Physical Install**: Skills are copied (not symlinked) to the global brain for Antigravity compatibility.
-   **📝 Dual-Write Pattern**: Enforces artifact persistence to `docs/` before handoff.
-   **⚙️ CONFIG.yaml Awareness**: All skills read `project/CONFIG.yaml` to understand stack and versions.
-   **🏗️ Architecture Tests**: Enforces Go Modern standards via AST analysis.

## 📂 Repository Structure

```
ag-skill-factory/
├── .agent/skills/           # 🏭 The Factory (internal tooling)
│   ├── skill-creator/       # Meta-skill that creates other skills
│   └── skill-factory-expert/# Project expert (gitignored, local)
│
├── squads/                  # 👥 Your Skills (gitignored, user-specific)
│   ├── backend-go-expert/
│   ├── frontend-nuxt/
│   ├── mcp-expert/
│   └── ...
│
├── cmd/skills/              # 🔧 CLI source code
├── internal/                # 📦 Core packages
│   ├── installer/           # Install/update/backport logic
│   ├── diff/                # Directory comparison
│   └── coverage/            # Meta-test coverage
│
├── Makefile                 # Build, test, install
└── README.md
```

## 🔧 Skills CLI

The `skills` CLI manages skill installation and synchronization across workspaces.

### Installation

```bash
# Clone and build
git clone https://github.com/ydnikolaev/ag-skill-factory.git
cd ag-skill-factory
make install
```

This builds the CLI to `bin/skills` and installs it to `/usr/local/bin/skills`.

### Commands

```bash
skills install    # Bootstrap .agent/ in current project
skills list       # Show skill inventory with sync status
skills update     # Pull latest from factory (with diff preview)
skills backport   # Push local changes back to factory
```

### Example Workflow

```bash
# 1. Go to your project
cd my-project

# 2. Install skills
skills install
# ✅ Installed 12 skills, 5 rules

# 3. Check status
skills list
# product-manager  ✓  ✓  synced
# my-custom-skill  ✓  -  local only

# 4. Update from factory
skills update
# Shows diff, asks for confirmation

# 5. Push improvements back
skills backport product-manager
# ✅ Backported 'product-manager' to factory
```

### Configuration

Config file: `~/.config/ag-skills/config.yaml`

```yaml
source: ~/Developer/antigravity/ag-skill-factory/squads
global_path: ~/.gemini/antigravity/global_skills
```

## 🧪 Development

```bash
# Run linter (FASCIST MODE)
make lint

# Run all tests
make test

# Build CLI
make build-skills

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
| `make build-skills` | Build CLI binary to `bin/skills` |
| `make install-skills` | Symlink binary to `/usr/local/bin` |
| `make install-completions` | Add shell completions for zsh/bash |
| `make lint` | Run golangci-lint (FASCIST MODE) |
| `make test` | Run all tests |
| `make clean` | Remove build artifacts |
| `make validate SKILL=<name>` | Validate a single skill |
| `make validate-all` | Validate all skills in squads/ |
| `make install-squads` | Copy squad skills to global brain |

## Artifact Persistence Rule (Dual-Write)

Antigravity artifacts (`brain/...`) are ephemeral. Documents (`docs/...`) are permanent.

**Rule**: Every skill MUST save its final output (e.g., `discovery-brief.md`) to the `docs/` directory and **change its status to "Approved"** BEFORE handing off to the next skill.

## 📖 How It Works

1.  **Design First**: Before creating a skill, answer: What triggers it? What's the decision tree?
2.  **Scaffold**: Run `init_skill.py` to create the standard structure in `squads/`.
3.  **Refine**: Edit `SKILL.md` with your logic and workflows.
4.  **Validate**: Run `make validate SKILL=<name>` to check quality.
5.  **Install**: Run `skills install` in your project.

## 🤝 Contributing

We welcome contributions! Please follow the **Design-First** philosophy:
1.  Keep instructions concise (<500 lines).
2.  Move details to `resources/` or `examples/`.
3.  Always include Team Collaboration and When to Delegate sections.
4.  Customize the checklist for your skill's domain.
5.  Run `make lint && make test` before committing.

## License

MIT
