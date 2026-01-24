# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added
- **Pre-Handoff Validation (Hard Stop)**: Injected into all 20 skills.
  - 5-point checklist before `notify_user` or delegation
  - Downstream skills reject invalid handoffs
- **TDD Task Creation (Hard Stop)**: Added to `backend-go-expert` and `frontend-nuxt`.
  - Phase 1 MUST be RED (Tests First)
  - `make check` verification loop
  - Git commit order verification (`test:` → `feat:` → `refactor:`)
- **Test Skeleton section**: Mandatory in `tech-spec-writer` output.
  - Unit/Integration test tables
  - TDD verification via git commit order
- **Forced Handoff Path**: `bmad-architect` → `tech-spec-writer` (direct skip forbidden).
- **`scripts/inject_pre_handoff.py`**: Batch injection tool for skill updates.

### Changed
- **`AGENTS.md` → `ARTIFACT_REGISTRY.md`**: Renamed across all 75+ references.
- **`artifact-registry.md` template**: Feature-centric Work Unit structure.
  - Lifecycle folders: `active/` → `review/` → `closed/`
  - Collapsible `<details>` for archived work
  - Quick Links navigation table
- **Factory skills updated**: `skill-creator`, `skill-factory-expert`, `skill-interviewer` templates now use `ARTIFACT_REGISTRY.md`.
- **`qa-lead` Gatekeeper**: Now checks git commit order for TDD compliance.

### Added
- **Afero filesystem abstraction**: Integrated `spf13/afero` for testable file I/O.
  - `Installer.Fs` field abstracts all file operations
  - `NewWithFs()` constructor for test injection
  - Production uses `OsFs`, tests use `MemMapFs`
- **Prompter interface for DI**: Enables stdin mocking in tests.
  - `StdinPrompter` for production
  - `MockPrompter` for tests
- **Coverage enforcement test**: `TestInstallerCoverage100` fails CI if coverage < 95%
- **`refactor-architect` skill**: Analyzes codebase and designs modular refactoring specs.
  - Runs static analysis (LOC, complexity, coverage gaps)
  - Queries Context7 for current best practices (mandatory)
  - Creates enforcement mechanisms (lint rules, pre-commit, CI, ADRs)
  - Delegates module specs to domain executors
- **`/refactor` workflow**: Full refactoring pipeline.
  - Activates `@refactor-architect` for analysis
  - Module-based execution by domain executors
  - Enforcement application by `@devops-sre`
  - Validation by `@qa-lead`

### Changed
- **Installer test architecture**: Consolidated 4 test files → 1 unified file.
  - Old: `operations_test.go`, `converter_test.go`, `rewriter_test.go`, `prompter_test.go`
  - New: `installer_test.go` (970 lines, clean afero-based tests)
  - Coverage: 41% → 95.9% (+134% increase)
- **Error testing strategy**: Replaced `chmod` hacks with `afero.ReadOnlyFs`
- **Refactored `internal/installer/`**: Split 518 LOC god-file into 5 focused files.
  - `installer.go` (154 LOC) — core operations
  - `operations.go` (199 LOC) — entry processing

### Fixed
- **MCP examples build failure**: Added `//go:build ignore` to `go-server-mcp-go.go` and `go-server-official.go` to prevent duplicate `main` declarations and undefined `mcp.NewError` errors

### Changed
- **Refactored `internal/installer/`**: Split 518 LOC god-file into 5 focused files.
  - `installer.go` (154 LOC) — core operations
  - `operations.go` (199 LOC) — entry processing
  - `converter.go` (83 LOC) — standards → rules
  - `rewriter.go` (54 LOC) — path rewriting
  - `utils.go` (62 LOC) — file utilities
- **Tightened lint rules**: funlen (80/50), cyclop (12), added gocognit/lll
- **Added `make check-loc`**: Enforce max 300 LOC per file

### Fixed
- **Module path**: Corrected GitHub username `yuranikolaev` → `ydnikolaev` in go.mod and all imports

### Added
- **Command tests**: Added 13+ test cases for `cmd/skills/` commands
  - `list_test.go`, `install_test.go`, `update_test.go`, `backport_test.go`
- **`skill-updater` meta-skill**: Maintains and updates existing skills.
- **`skill-updater` meta-skill**: Maintains and updates existing skills.
  - Mass rollout of new patterns and standards
  - Preview → Approve → Apply workflow
  - Complements skill-creator (create) and skill-interviewer (design)
- **Tech Debt Protocol**: New standard for tracking TODOs and workarounds.
  - `_standards/TECH_DEBT_PROTOCOL.md` — central protocol document
  - `project/docs/TECH_DEBT.md` — registry template
  - Added to 9 developer skills: backend, frontend, debugger, mcp, tma, tui, cli, ui, devops
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
