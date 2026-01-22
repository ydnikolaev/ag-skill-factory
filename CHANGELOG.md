# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added
- **`project-bro` skill**: New utility skill that understands project state, reads docs, analyzes code.
  - Answers "where are we?", "what's done?", "what's next?"
  - Reads `project/CONFIG.yaml`, `docs/AGENTS.md`, and codebase
  - Added to PIPELINE.md as utility available at any phase
- **CONFIG.yaml Awareness**: All skills now check `project/CONFIG.yaml` before technical decisions.
  - Prevents wrong library suggestions (e.g., Chi instead of stdlib)
  - Added "First Step" protocol to skill-creator template
  - Mass-updated all 16 squad skills with awareness block
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
