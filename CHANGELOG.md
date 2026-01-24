# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added
- **`skill-interviewer` meta-skill**: Creative partner for ideating new skills.
  - Interview-first approach with gap analysis
  - Anti-pattern detection (too big, too small, too vague)
  - Writes skill specs for `@skill-creator` to execute
- **`workflow-creator` meta-skill**: Designs and creates `.agent/workflows/`.
  - Interview-based workflow design
  - Duplicate checking against existing workflows
  - Turbo annotation guidance (`// turbo`, `// turbo-all`)
- **`/commit` workflow**: Pre-commit automation for ag-skill-factory.
  - Branch protection with auto-branch creation
  - Self-evolve sync as first step
  - CHANGELOG update (Keep a Changelog format)
  - Conventional Commits message generation
- **`/push` workflow**: Full push pipeline.
  - Runs `/commit` first
  - Runs `make install` before push
  - Pushes to remote

### Changed
- **`project/docs/` convention**: All artifact paths in skill-creator templates now use `project/docs/` prefix.
  - Updated `AGENTS.md` template
  - Updated `checklist.md` template
  - Updated `SKILL.md` template

### Added
- **`project-bro` skill**: New utility skill that understands project state, reads docs, analyzes code.
  - Answers "where are we?", "what's done?", "what's next?"
  - Reads `project/CONFIG.yaml`, `docs/AGENTS.md`, and codebase
  - Added to PIPELINE.md as utility available at any phase
- **CONFIG.yaml Awareness**: All skills now check `project/CONFIG.yaml` before technical decisions.
  - Prevents wrong library suggestions (e.g., Chi instead of stdlib)
  - Added "First Step" protocol to skill-creator template
  - Mass-updated all 16 squad skills with awareness block
- **Project MCP Awareness**: Skills now know about project-level MCP servers.
  - Check `mcp.yaml` for project MCP config
  - Use project MCP tools/resources (named after project)
  - Reference `mcp.yaml → context7.default_libraries` for pre-configured docs
- **Artifact Persistence**: Enforced "Dual-Write Pattern" across the factory.
  - Updated `checklist.md` template with **MANDATORY** persistence check.
  - Updated `SKILL.md` template with `Artifact Ownership` and `Handoff Protocol` sections.
  - Batch updated ALL 15 existing skill checklists in `squads/`.
- **Validation**: `validate_skill.py` now warns if `Artifact Ownership` or `Handoff Protocol` sections are missing.

### Fixed
- **Critical**: Fixed global skills path from `~/.gemini/antigravity/skills/` to `~/.gemini/antigravity/global_skills/` per updated Antigravity documentation
  - Updated `Makefile` GLOBAL_SKILLS_DIR variable
  - Updated `README.md` installation instructions
  - Updated `agent skills.md` local documentation
  - Updated `skill-creator/SKILL.md` references
  - Updated `skill-factory-expert/SKILL.md` references
  - Updated `init_skill.py` output message

### Changed
- `skill-factory-expert`: Enhanced Makefile Commands table with detailed descriptions
- `skill-factory-expert`: Added note about physical copy vs symlinks
