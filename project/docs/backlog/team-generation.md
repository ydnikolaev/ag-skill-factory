# Backlog: Team Generation Improvements

> Items for @backend-go-expert to implement in factory CLI.

## P1: Factory Install Preset Selection ✅ DONE (generate_teams.py)

Team files are now generated per preset in `blueprint/_meta/teams/`.

---

## P2: Factory Install Uses Correct Team File

**Goal:** `factory install --preset=tma` should copy the correct team file.

**Current:** 
- No automatic team file selection

**Target:**
- Read `--preset` flag
- Copy `blueprint/_meta/teams/TEAM_<preset>.md` → `project/.agent/rules/TEAM.md`

**Files to modify:**
- `internal/installer/installer.go`
- `cmd/factory/install.go`

---

## Notes

- `make generate-teams` generates all preset files via Python script
- Team files are in `blueprint/_meta/teams/`
