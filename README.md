# Antigravity Skill Factory 🚀

> **Design High-Quality Autonomous Agent Skills.**
> A powerful framework for creating standardized, effective, and "native" skills for Antigravity and MCP-based agents.

[![Antigravity](https://img.shields.io/badge/Antigravity-Native-purple)](https://antigravity.google)
[![Agent Skills](https://img.shields.io/badge/Agent-Skills-blue)](https://github.com/anthropics/skills)
[![Python](https://img.shields.io/badge/Python-3.10%2B-yellow)](https://python.org)

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
-   **🛠️ Python Scaffolding**: `init_skill.py` automates directory creation, adhering to strict standards.
-   **📚 Design Guide**: The `skill-creator` serves as a textbook for agents on *how* to design good tools.
-   **📦 Physical Install**: Skills are copied (not symlinked) to the global brain for Antigravity compatibility.
-   **📝 Dual-Write Pattern**: Enforces artifact persistence to `docs/` before handoff.

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
├── Makefile                 # Validates and installs skills
└── README.md
```

## 📦 Installation

```bash
# Clone the repository
git clone https://github.com/ydnikolaev/ag-skill-factory.git
cd ag-skill-factory

# Install all skills to global brain (validates first)
make install
```

This **copies** skills to `~/.gemini/antigravity/global_skills/` (physical copy, not symlink).

## 🚀 Usage

Once installed, simply ask your Antigravity agent:

> "Create a new skill called 'docker-manager' that helps me handle containers."

Or run it manually:

```bash
# Create a new skill in squads/
python3 .agent/skills/skill-creator/scripts/init_skill.py my-skill

# Validate your skill
make validate SKILL=my-skill

# Install to global brain
make install-squads
```

## 🔧 Makefile Commands

| Command | Description |
|---------|-------------|
| `make install` | Validate all, generate TEAM.md, install everything |
| `make validate SKILL=<name>` | Validate a single skill |
| `make validate-all` | Validate all skills in squads/ |
| `make generate-team` | Regenerate `squads/TEAM.md` |
| `make install-factory` | Install only factory skills (.agent/skills/) |
| `make install-squads` | Install only squad skills |
| `make uninstall` | Remove all skills from global brain |

## Artifact Persistence Rule (Dual-Write)

Antigravity artifacts (`brain/...`) are ephemeral. Documents (`docs/...`) are permanent.
**Rule**: Every skill MUST save its final output (e.g., `discovery-brief.md`) to the `docs/` directory **BEFORE** handing off to the next skill. This is enforced by mandatory checklist items.

## 📖 How It Works

1.  **Design First**: Before creating a skill, answer: What triggers it? What's the decision tree?
2.  **Scaffold**: Run `init_skill.py` to create the standard structure in `squads/`.
3.  **Refine**: Edit `SKILL.md` with your logic and workflows.
4.  **Validate**: Run `make validate SKILL=<name>` to check quality.
5.  **Install**: Run `make install-squads` to copy to global brain.

## 🤝 Contributing

We welcome contributions! Please follow the **Design-First** philosophy:
1.  Keep instructions concise (<500 lines).
2.  Move details to `resources/` or `examples/`.
3.  Always include Team Collaboration and When to Delegate sections.
4.  Customize the checklist for your skill's domain.

## License

MIT
