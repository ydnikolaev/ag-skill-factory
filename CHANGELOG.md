# Changelog

All notable changes to this project will be documented in this file.
## [Unreleased]

### Added

- Initial commit of skill-creator
- Add makefile for SSOT installation
- Add hello-world example and strict validation
- Complete TMA Squad with 10 production-ready skills
- Add delegation guidelines to product/analyst/architect
- Add delegation to all skills + pipeline in README
- Split frontend into generic nuxt + specialized tma-expert
- Add mcp-expert skill and enhance skill-creator workflow
- Implement Artifact Persistence (Dual-Write) pattern
- Implement explicit Dual-Write Pattern (Iteration Protocol)
- Add project-bro skill and CONFIG.yaml awareness
- Add project MCP awareness to all skills
- Add skills CLI with architecture enforcement
- Add version command
- Add meta-skills and automation workflows
- Add skill-updater meta-skill
- Track skill-factory-expert and add Factory Skills section to README
- Add verbose logging to updateRules
- Add --force flag to 'skills update' for path rewriting
- Add refactor-architect skill and /refactor workflow
- Integrate afero for gold-standard testable I/O
- Add doc-janitor skill for document cleanup
- Add doctor, link checker, versioning, presets
- Add preset selection with TUI
- Add skill catalog website (VitePress)
- Automate changelog with git-cliff

### Changed

- Separate factory (.agent/skills) from products (squads/)
- Update /push to merge locally instead of PR
- Split god-file + add tests + enforce lint
- Add tests + fix MCP examples build
- Pipeline hardening rollout
- Inject Pre-Handoff Validation into all skills
- Update factory skills with ARTIFACT_REGISTRY.md
- Rename ag-skill-factory to antigravity-factory

### Documentation

- Update readme with gold standard and validation features
- Clean README from squad-specific content
- Enforce 'Draft -> Approved' status change and explicit Dual-Write explanation in all skills
- Update README with skills CLI documentation
- Update project docs with factory skills and workflows
- Add skill-updater and Tech Debt Protocol to CHANGELOG
- Update skill-factory-expert with skill-updater and new standards
- Add afero integration and 95% coverage enforcement info
- Add coverage badge (95.9%)
- Add pipeline hardening release notes
- Update README and AGENTS

### Fixed

- Improve context7 queries in bmad-architect
- Use physical copy instead of symlinks for skill installation
- Update global skills path from skills/ to global_skills/
- Remove global path sync from skills install
- Add Dual-Write Pattern sections to feature-fit skill
- Embed self-evolve checks in /commit workflow
- Add rules sync to 'skills update' command
- Rewrite _standards/ paths to .agent/rules/ during install/update
- Use copyDirWithRewrite in showChangesAndApply
- Correct GitHub username yuranikolaev → ydnikolaev
- Ignore dead links in VitePress build
- Strip internal links in catalog generator

### Miscellaneous

- Gitignore squads/ and reference/ (user-specific content)
- Un-gitignore squads/ to make skills public
- Self-evolve skill-factory-expert (19→20 skills)
- Add doc-janitor to roster (21 skills)
- Self-evolve skill-factory-expert (20→21 skills)
- Add blueprint/private for sensitive skills


