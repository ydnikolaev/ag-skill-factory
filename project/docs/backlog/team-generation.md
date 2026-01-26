# Backlog: Team/Pipeline Integration

> Items for @backend-go-expert to implement in factory CLI.

## P1: Factory Install Uses Correct Team/Pipeline Files

**Goal:** `factory install --preset=tma` should copy the correct files.

**Current:** 
- No automatic team/pipeline file selection

**Target:**
- Read `--preset` flag (default: `all`)
- Copy `blueprint/_meta/_teams/TEAM_<preset>.md` → `project/.agent/rules/TEAM.md`
- Copy `blueprint/_meta/_pipelines/PIPELINE_<preset>.md` → `project/.agent/rules/PIPELINE.md`

**Files to modify:**
- `internal/installer/installer.go`
- `cmd/factory/install.go`

---

## P1: Factory Install Protection for project/docs

> [!CAUTION]
> **CRITICAL:** Never overwrite existing documentation!

**Rule:** If `project/docs/` exists, DO NOT copy folder structure.

```go
if exists("project/docs/") {
    log.Warn("project/docs/ exists, skipping document structure")
    return // DO NOT COPY
}
```

---

## P1: Factory Install Copies cliff.toml

**Goal:** Auto-generate CHANGELOG.md via git-cliff.

**Action:** Copy `_meta/_docs/templates/_cliff.toml` → `project/cliff.toml`

**Always copy** (overwrite if exists — it's a config file).

---

## P2: Factory New Work Unit Command

**Goal:** `factory new-work-unit feat-forum-topics` creates full structure.

**Creates:**
```
project/docs/
├── registry/feat-forum-topics.md    ← from _work-unit-registry.md
├── active/
│   ├── discovery/feat-forum-topics.md
│   ├── specs/feat-forum-topics.md
│   └── ...
```

**Template source in factory:** `blueprint/_meta/_docs/templates/`

**Copies to project:** `project/docs/_templates/`

---

## P1: Factory Install Copies Doc Templates

**Goal:** `factory install` creates docs structure with templates.

**If `project/docs/` does NOT exist:**

```
project/docs/
├── _templates/
│   ├── _ARTIFACT_REGISTRY.md
│   ├── _discovery-brief.md
│   ├── _tech-spec.md
│   └── ...
├── ARTIFACT_REGISTRY.md
├── active/{discovery,product,design,specs,backend,frontend,qa}
├── review/{...}
├── closed/{sprints,features,bugs,refactoring}
└── registry/
```

**If `project/docs/` EXISTS:** Skip entirely, DO NOT overwrite.

---

## P2: Doc Matrix Generator

**Goal:** Generate DOC_MATRIX_<preset>.md per preset.

**Like skill-matrix, but for documents:**
- Who creates each doc type
- Who consumes it
- Phase it belongs to

**Source:** `blueprint/_meta/_docs/doc-types.yaml`

---

## P3: Auto-Generate presets.yaml from Skills

**Goal:** 100% SSOT — presets.yaml generated from skill frontmatter.

**Approach:**
1. Add `presets: [backend, fullstack, all]` to each skill frontmatter
2. Generator collects presets tags and builds presets.yaml

**Example:**
```yaml
# In skill frontmatter
presets:
  - backend
  - fullstack
  - all
```

**Trade-offs:**
- More text in each skill
- No `extends` inheritance (must list explicitly)
- But: full SSOT, skill knows its presets

---

## P3: Enrich DOCUMENT_STRUCTURE_PROTOCOL

**Goal:** Make protocol comprehensive enough to not need separate AGENTS.md.

**Add:**
- Quick start section (first-time setup)
- Registry module usage examples
- Project lifecycle examples (from discovery to closed)

---

## P3: Project Doc Templates

**Goal:** Bootstrap project-specific documentation.

**Add to `docs/_project/`:**
```
_project/
├── _CODEBASE_MAP.md     # Project code structure
├── _CONVENTIONS.md      # Code style, naming
├── _DECISION_LOG.md     # ADR log
├── _GLOSSARY.md         # Project terminology
├── _KNOWN_ISSUES.md     # Known bugs, workarounds
├── _CHANGELOG.md        # Release history
└── _BACKLOG.md          # Future work, ideas
```

**Not workflow docs (active/review/closed) — project knowledge docs.**

---

## Notes

- Team files: `blueprint/_meta/_teams/TEAM_*.md`
- Pipeline files: `blueprint/_meta/_pipelines/PIPELINE_*.md`
- Doc templates: `blueprint/_meta/_docs/templates/`
- Generated via `make generate-all`
