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
-   **🏆 Gold Standard Example**: Includes a reference `hello-world` skill to demonstrate best practices.
-   **🛡️ Strict Validation**: `validate_skill.py` enforces the 500-line limit and checks for IDE-aware tool usage.
-   **✅ Auto-Checklists**: Generates `checklist.md` for quality assurance.
-   **🛠️ Python Scaffolding**: `init_skill.py` automates directory creation, adhering to strict standards.
-   **📚 Design Guide**: The `skill-creator` serves as a textbook for agents on *how* to design good tools.

## 📂 Repository Structure

```
ag-skill-factory/
├── .agent/skills/         # 🏭 The Factory
│   └── skill-creator/     # The meta-skill that creates other skills
│       ├── SKILL.md       # Design philosophy
│       ├── scripts/       # init_skill.py, validate_skill.py
│       └── resources/     # Templates, checklists
│
└── squads/                # 👥 Your Skills (gitignored, user-specific)
    └── ...                # Skills you create live here
```

## 📦 Installation

```bash
# Clone the repository
git clone https://github.com/ydnikolaev/ag-skill-factory.git
cd ag-skill-factory

# Install the skill-creator globally
make install
```

## 🚀 Usage

Once installed, simply ask your Antigravity agent:

> "Create a new skill called 'docker-manager' that helps me handle containers."

Or run it manually:

```bash
# Create a new skill in squads/ (your local skills folder)
python3 ~/.gemini/antigravity/skills/skill-creator/scripts/init_skill.py my-skill --output squads/

# Install your custom skills
make install-squads
```

## 📖 How It Works

1.  **Design First**: Before creating a skill, answer: What triggers it? What's the decision tree?
2.  **Scaffold**: Run `init_skill.py` to create the standard structure.
3.  **Refine**: Edit `SKILL.md` with your logic and workflows.
4.  **Verify**: Use the built-in checklist to validate quality.

## 🤝 Contributing

We welcome contributions! Please follow the **Design-First** philosophy:
1.  Keep instructions concise.
2.  Move details to `resources/`.
3.  Always include a verification step.

## License

MIT
