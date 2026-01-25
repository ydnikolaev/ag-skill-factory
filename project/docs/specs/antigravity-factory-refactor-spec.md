---
status: Approved
owner: "@skill-factory-expert"
created: 2026-01-25
updated: 2026-01-25
---

# Antigravity Factory Refactor Spec

> Эволюция от "Skill Factory" к полноценной **"Antigravity Factory"** — фабрика всего агентского контекста (skills, workflows, rules).

## Upstream Documents

- Brain artifact: `brain/<conversation-id>/implementation_plan.md`

## Решения

| Вопрос | Решение |
|--------|---------|
| Папка шаблона | `blueprint/` |
| Разделение | `standards/` = protocols, `rules/` = TEAM + PIPELINE |
| CLI имя | `factory` (без ag) |
| Команды | `install` + `list` (удаляем update, backport) |
| Репо | `antigravity-factory` |
| Backwards compatibility | Нет, сжигаем мосты 🔥 |

---

## Новая архитектура

```
antigravity-factory/
├── .agent/                      # 🔧 Factory-internal (НЕ копируется)
│   ├── skills/
│   │   ├── skill-creator/
│   │   ├── skill-factory-expert/
│   │   ├── skill-interviewer/
│   │   ├── skill-updater/
│   │   └── workflow-creator/
│   └── workflows/
│       ├── commit.md            # Factory-specific
│       ├── push.md              # Factory-specific
│       └── self-evolve.md       # Factory-specific
│
├── blueprint/                   # 📦 Копируется целиком → .agent/
│   ├── skills/
│   │   ├── backend-go-expert/
│   │   ├── frontend-nuxt/
│   │   ├── mcp-expert/
│   │   └── ... (21 skills)
│   ├── workflows/
│   │   ├── doc-cleanup.md
│   │   └── refactor.md
│   ├── rules/
│   │   ├── TEAM.md
│   │   └── PIPELINE.md
│   └── standards/
│       ├── TDD_PROTOCOL.md
│       ├── GIT_PROTOCOL.md
│       ├── TECH_DEBT_PROTOCOL.md
│       ├── TRACEABILITY_PROTOCOL.md
│       └── DOCUMENT_STRUCTURE_PROTOCOL.md
│
├── cmd/factory/                 # 🔧 CLI
├── internal/installer/          # Installer logic (упрощённый)
├── Makefile
└── README.md
```

---

## Requirements Checklist

### Core Changes

| # | Requirement | Status |
|---|------------|--------|
| 1 | Remove global_skills installation from Makefile | ✅ |
| 2 | Create blueprint/ folder structure | ✅ |
| 3 | Move skills from squads/ to blueprint/skills/ | ✅ |
| 4 | Move _standards/ to blueprint/standards/ | ✅ |
| 5 | Create blueprint/rules/ with TEAM.md, PIPELINE.md | ✅ |
| 6 | Create blueprint/workflows/ with doc-cleanup.md, refactor.md | ✅ |
| 7 | Update `_standards/` refs → `../standards/` in 23 skills | ✅ |
| 8 | Delete squads/ folder | ✅ |

### CLI Changes

| # | Requirement | Status |
|---|------------|--------|
| 9 | Rename cmd/skills/ → cmd/factory/ | ✅ |
| 10 | Remove update.go, backport.go | ✅ |
| 11 | Simplify installer (remove rewriter.go, converter.go logic) | ✅ |
| 12 | Update install.go for simple copy | ✅ |
| 13 | Update list.go for categories | ✅ |
| 14 | Update config default source to blueprint/ | ✅ |

### Go Module

| # | Requirement | Status |
|---|------------|--------|
| 15 | Rename module to github.com/ydnikolaev/antigravity-factory | ✅ |
| 16 | Update all import paths | ✅ |

### Documentation

| # | Requirement | Status |
|---|------------|--------|
| 17 | Update README.md (squads/ → blueprint/) | ✅ |
| 18 | Update AGENTS.md | ✅ |
| 19 | Update skill-factory-expert SKILL.md (self-evolve) | ✅ |
| 20 | Update validate_skill.py paths | ✅ |

### Testing

| # | Requirement | Status |
|---|------------|--------|
| 21 | Add testscript dependency | ❌ |
| 22 | Write E2E test (factory install in tmp) | ✅ (manual) |
| 23 | Verify no _standards/ refs in installed skills | ✅ |
| 24 | Verify no squads/ refs anywhere | ❌ |
| 25 | All Go tests pass | ✅ |
| 26 | Lint passes | ❌ |

---

## Installer Simplification

### Remove

- `rewriter.go` — path rewriting no longer needed
- `converter.go` — YAML frontmatter conversion for standards no longer needed
- `Update()` method
- `Backport()` method  
- `ForceRefresh()` method

### Simplify

```go
type InstallResult struct {
    SkillCount    int
    WorkflowCount int
    RuleCount     int
    StandardCount int
}

func (i *Installer) Install() (*InstallResult, error) {
    // 1. Remove existing .agent/ (if exists)
    // 2. Copy blueprint/ → .agent/ as-is
    // 3. Count each category
    return result, nil
}
```

---

## Path Updates Required

### In Skills (23 files)

Replace `_standards/X.md` → `../standards/X.md`:

```
_standards/DOCUMENT_STRUCTURE_PROTOCOL.md → ../standards/DOCUMENT_STRUCTURE_PROTOCOL.md
_standards/TDD_PROTOCOL.md → ../standards/TDD_PROTOCOL.md
_standards/GIT_PROTOCOL.md → ../standards/GIT_PROTOCOL.md
_standards/TECH_DEBT_PROTOCOL.md → ../standards/TECH_DEBT_PROTOCOL.md
_standards/TRACEABILITY_PROTOCOL.md → ../standards/TRACEABILITY_PROTOCOL.md
```

### In Documentation (25+ files)

Replace `squads/` references with `blueprint/`:

- `squads/` → `blueprint/skills/`
- `squads/_standards/` → `blueprint/standards/`
- `squads/TEAM.md` → `blueprint/rules/TEAM.md`

---

## E2E Test (testscript)

```txtar
# Test factory install creates correct structure

exec factory install

# Verify directories
exists .agent/skills
exists .agent/workflows
exists .agent/rules
exists .agent/standards

# Verify files
exists .agent/skills/backend-go-expert/SKILL.md
exists .agent/workflows/doc-cleanup.md
exists .agent/rules/TEAM.md
exists .agent/standards/TDD_PROTOCOL.md

# Verify NO old path references
! grep '_standards/' .agent/skills/backend-go-expert/SKILL.md
! grep 'squads/' .agent/
```

---

## Handoff

**Next skill:** `@backend-go-expert`

**Scope:** Execute this spec — all Go code changes in cmd/ and internal/
