# Skill Maintainer

> Future skill for maintaining and updating Antigravity skills at scale.

## Key Decisions

### Template Standards

1. **Filename convention**: `lowercase-kebab-case.md`
   - All document templates use underscore prefix: `_api-contracts.md`
   - Consistent, easy to parse, cross-platform safe

2. **Frontmatter requirements for per-feature docs**:
   ```yaml
   status: Draft
   owner: @skill-name
   lifecycle: per-feature
   work_unit: {WORK_UNIT}
   upstream:
     - doc_type: parent-doc
       owner: @parent-skill
   downstream:
     - skill: @next-skill
   created: {DATE}
   updated: {DATE}
   ```

3. **Frontmatter for living docs** (no dependencies):
   ```yaml
   status: Draft
   owner: @skill-name
   lifecycle: living
   created: {DATE}
   updated: {DATE}
   ```
   - Omit empty `upstream: []` / `downstream: []`

### Build Architecture

1. **Source of Truth**: `src/` directory
   - `src/skills/` — skill SKILL.md with `{{include:}}` directives
   - `src/partials/` — reusable sections
   - `src/rules/` — protocol markdown files
   - `src/templates/` — document templates
   - `src/workflows/` — workflow definitions
   - `src/_meta/` — generated metadata (skill-matrix, presets, pipelines, teams)

2. **Build Output**: `dist/` directory (gitignored)
   - `dist/_agent/` — ready for → `.agent/` (includes resolved)
   - `dist/project/` — ready for → `project/`

3. **Naming convention in dist**:
   - `_agent` (underscore) to avoid Antigravity treating as system folder
   - Factory install copies `_agent/` → `.agent/`

### Installer Improvements (Backlog)

See: `backlog/installer-merge-mode.md`
- Merge mode instead of destructive overwrite
- Per-folder copy logic
- Config file skip-if-exists

## Future Scope

- Mass rollout of changes across multiple skills
- Validation of skill consistency
- Auto-update of deprecated patterns
- Version tracking per skill
