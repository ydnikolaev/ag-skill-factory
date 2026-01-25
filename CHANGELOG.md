# Changelog

All notable changes to this project will be documented in this file.
## [Unreleased]

### Added

- Add versioning, workflows, and validation improvements
- Premium landing page redesign
- Automate changelog with git-cliff
- Add skill catalog website (VitePress)
- Add preset selection with TUI
- Add doctor, link checker, versioning, presets
- Add doc-janitor skill for document cleanup
- Integrate afero for gold-standard testable I/O
- Add refactor-architect skill and /refactor workflow
- Add --force flag to 'skills update' for path rewriting
- Add verbose logging to updateRules
- Track skill-factory-expert and add Factory Skills section to README
- Add skill-updater meta-skill
- Add meta-skills and automation workflows
- Add version command
- Add skills CLI with architecture enforcement
- Add project MCP awareness to all skills
- Add project-bro skill and CONFIG.yaml awareness
- Implement explicit Dual-Write Pattern (Iteration Protocol)
- Implement Artifact Persistence (Dual-Write) pattern
- Add mcp-expert skill and enhance skill-creator workflow
- Split frontend into generic nuxt + specialized tma-expert
- Add delegation to all skills + pipeline in README
- Add delegation guidelines to product/analyst/architect
- Complete TMA Squad with 10 production-ready skills
- Add hello-world example and strict validation
- Add makefile for SSOT installation
- Initial commit of skill-creator

### Changed

- Rename ag-skill-factory to antigravity-factory
- Update factory skills with ARTIFACT_REGISTRY.md
- Inject Pre-Handoff Validation into all skills
- Pipeline hardening rollout
- Add tests + fix MCP examples build
- Split god-file + add tests + enforce lint
- Update /push to merge locally instead of PR
- Separate factory (.agent/skills) from products (squads/)

### Documentation

- Update README and AGENTS
- Add pipeline hardening release notes
- Add coverage badge (95.9%)
- Add afero integration and 95% coverage enforcement info
- Update skill-factory-expert with skill-updater and new standards
- Add skill-updater and Tech Debt Protocol to CHANGELOG
- Update project docs with factory skills and workflows
- Update README with skills CLI documentation
- Enforce 'Draft -> Approved' status change and explicit Dual-Write explanation in all skills
- Clean README from squad-specific content
- Update readme with gold standard and validation features

### Fixed

- Sort changelog entries newest-first
- Strip internal links in catalog generator
- Ignore dead links in VitePress build
- Correct GitHub username yuranikolaev → ydnikolaev
- Use copyDirWithRewrite in showChangesAndApply
- Rewrite _standards/ paths to .agent/rules/ during install/update
- Add rules sync to 'skills update' command
- Embed self-evolve checks in /commit workflow
- Add Dual-Write Pattern sections to feature-fit skill
- Remove global path sync from skills install
- Update global skills path from skills/ to global_skills/
- Use physical copy instead of symlinks for skill installation
- Improve context7 queries in bmad-architect

### Miscellaneous

- Update changelog
- Update changelog
- Add blueprint/private for sensitive skills
- Self-evolve skill-factory-expert (20→21 skills)
- Add doc-janitor to roster (21 skills)
- Self-evolve skill-factory-expert (19→20 skills)
- Un-gitignore squads/ to make skills public
- Gitignore squads/ and reference/ (user-specific content)


