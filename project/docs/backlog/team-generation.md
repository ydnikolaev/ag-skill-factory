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

## Notes

- Team files: `blueprint/_meta/_teams/TEAM_*.md`
- Pipeline files: `blueprint/_meta/_pipelines/PIPELINE_*.md`
- Doc templates: `blueprint/_meta/_docs/templates/`
- Generated via `make generate-all`

