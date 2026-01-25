# Technical Specification: ag-skills CLI

**Status**: Approved
**Date**: 2026-01-23
**Source**: `docs/discovery/ag-skills-cli.md`

---

## Overview

Go CLI утилита для двусторонней синхронизации Antigravity скиллов между фабрикой (`ag-skill-factory/squads/`) и проектами (`.agent/skills/`).

---

## Commands

### 1. `skills install`

**Behavior**: 
> Копирует все скиллы из source в текущий воркспейс, создаёт структуру `.agent/`, конвертирует standards в rules.

### 2. `skills update`

**Behavior**:
> Обновляет скиллы в проекте из source. Показывает diff для изменённых файлов, спрашивает подтверждение.

### 3. `skills backport <name>`

**Behavior**:
> Копирует изменённый скилл из проекта обратно в фабрику. Показывает diff, спрашивает подтверждение.

### 4. `skills list`

**Behavior**:
> Показывает таблицу скиллов: установленные в проекте, доступные в source, статус синхронизации.

---

## Config File

**Location**: `~/.config/ag-skills/config.yaml`

```yaml
source: /path/to/ag-skill-factory/squads
global_path: ~/.gemini/antigravity/global_skills
```

---

## File Structure After Install

```text
project/.agent/
├── skills/           # Full copies of skills
├── rules/            # Converted standards
└── workflows/        # Empty (for future)
```

---

## Implementation

See full spec in brain artifacts for pseudocode and edge cases.
Implemented in Go with Cobra CLI framework.
