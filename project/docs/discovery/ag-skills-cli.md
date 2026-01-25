# Discovery Brief: ag-skills CLI

**Status**: Approved
**Date**: 2026-01-23
**Author**: @idea-interview

---

## Problem Statement

1. **Global skills не читаются**: После апдейта агенты перестали активировать скиллы из `~/.gemini/antigravity/global_skills/`.
2. **Нет проектной установки**: Текущий `make install` копирует только в глобальный brain.
3. **Неудобный DX**: Нужно вручную управлять скиллами в каждом воркспейсе.

---

## Solution

Go CLI утилита `skills` для двусторонней синхронизации скиллов между фабрикой и проектами.

---

## Commands (MVP)

| Команда | Описание | Направление |
|---------|----------|-------------|
| `skills install` | Установить скиллы + создать `.agent/` | Factory → Project |
| `skills update` | Обновить скиллы из source (с diff) | Factory → Project |
| `skills backport <name>` | Вернуть изменения в фабрику (с diff) | Project → Factory |
| `skills list` | Показать установленные + доступные | — |

### Поведение с Diff

При `update` и `backport`:
1. Показать diff изменений
2. Спросить подтверждение
3. Применить или отменить

---

## Структура после `skills install`

```text
/some-project/
├── .agent/
│   ├── skills/           # ← Полные копии скиллов
│   │   ├── backend-go-expert/
│   │   ├── frontend-nuxt/
│   │   └── ...
│   ├── rules/            # ← Преобразованные standards + team + pipeline
│   │   ├── tdd_protocol.md
│   │   ├── git_protocol.md
│   │   ├── team.md
│   │   └── pipeline.md
│   └── workflows/        # ← Пустая папка (на будущее)
└── ...
```

---

## Technical Decisions

| Аспект | Решение |
|--------|---------|
| **Symlink vs Copy** | Полная копия (симлинки не работают) |
| **Rules conversion** | Автоконвертация в YAML frontmatter формат |
| **Source path** | Конфиг `~/.config/ag-skills/config.yaml` |
| **Binary** | Через `make install` в Makefile |
| **Global skills** | Продолжаем дублировать |
| **Versioning** | Убрать скиллы из .gitignore |

---

## Config File

```yaml
# ~/.config/ag-skills/config.yaml
source: /Users/yuranikolaev/Developer/antigravity/ag-skill-factory/squads
```

---

## Not in MVP

- ❌ `skills remove`
- ❌ `skills validate` (валидация через `make install` в фабрике)
- ❌ Git auto-commit

---

## Next Steps

1. ✅ Discovery complete
2. ✅ Persisted to `docs/discovery/`
3. → `@tech-spec-writer`: Детальный спек CLI
4. → `@cli-architect` + `@backend-go-expert`: Разработка
