# Backlog: Team/Pipeline Integration

> Items for @backend-go-expert to implement in factory CLI.

## P1: Factory Install Uses Correct Team/Pipeline Files

**Goal:** `factory install --preset=tma` should copy the correct files.

**Current:** 
- No automatic team/pipeline file selection

**Target:**
- Read `--preset` flag (default: `all`)
- Copy `blueprint/_meta/teams/TEAM_<preset>.md` → `project/.agent/rules/TEAM.md`
- Copy `blueprint/_meta/pipelines/PIPELINE_<preset>.md` → `project/.agent/rules/PIPELINE.md`

**Files to modify:**
- `internal/installer/installer.go`
- `cmd/factory/install.go`

---

## Notes

- Team files: `blueprint/_meta/teams/TEAM_*.md`
- Pipeline files: `blueprint/_meta/pipelines/PIPELINE_*.md`
- Generated via `make generate-all`
