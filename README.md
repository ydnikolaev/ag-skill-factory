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

## 📂 Repository Structure

This repository separates the **factory** from the **products**:

```
ag-skill-factory/
├── .agent/skills/         # 🏭 Factory (the skill-creator itself)
│   └── skill-creator/     # The meta-skill that creates other skills
│       ├── SKILL.md
│       ├── scripts/
│       └── resources/
│
├── squads/                # 👥 Products (pre-built squad skills)
│   ├── backend-go-expert/
│   ├── frontend-nuxt/
│   ├── tma-expert/
│   ├── cli-architect/
│   └── ...11 skills total
│
├── Makefile               # Installation automation
└── README.md
```

## 📦 Installation

```bash
# Clone the repository
git clone https://github.com/ydnikolaev/ag-skill-factory.git
cd ag-skill-factory

# Install everything (factory + squads)
make install

# Or install separately:
make install-factory   # Only skill-creator
make install-squads    # Only squad skills
```

## 🚀 Usage

Once installed, simply ask your Antigravity agent:

> "Create a new skill called 'docker-manager' that helps me handle containers."

Or run it manually:

```bash
# Create a new skill in squads/
python3 ~/.gemini/antigravity/skills/skill-creator/scripts/init_skill.py docker-manager --output squads/
```

## 👥 Pre-built Development Squad (11 Skills)

A pre-built team of specialized skills for **full-stack development** (Go 1.25 + Nuxt 4 + DDD/BMAD V6).

### Core Pipeline
```
📋 Product Manager  →  📝 Systems Analyst  →  🧠 BMAD Architect
       ↓                      ↓                      ↓
   "Why?"               "What?"                "How?"
                                                   ↓
                    ┌──────────────────────────────┴──────────────────────────────┐
                    ↓                              ↓                              ↓
          ⚙️ Backend Go                   🎨 Frontend Nuxt                 🤖 Telegram Mechanic
                    ↓                              ↓                              ↓
                    └──────────────────────────────┬──────────────────────────────┘
                                                   ↓
                                            🛡️ QA Lead  →  🚀 DevOps SRE
```

### Specialized Add-ons
| Add-on | Purpose |
|---|---|
| `tma-expert` | Telegram Mini Apps (`@tma.js/sdk`) |
| `cli-architect` | CLI tools (Cobra/Viper) |
| `tui-charm-expert` | Terminal UIs (BubbleTea/Lipgloss) |

Each skill knows **when to delegate** and **when to return** to ensure smooth handoffs.

## 🤝 Contributing

We welcome contributions! Please follow the **Design-First** philosophy:
1.  Keep instructions concise.
2.  Move details to `resources/`.
3.  Always include a verification step.

## License

MIT
