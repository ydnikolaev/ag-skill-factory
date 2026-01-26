# Rules Matrix

> Mapping of rules to activation triggers and target skills.

## Legend

| Trigger | Meaning |
|---------|---------|
| `always_on` | Applied to every conversation |
| `model_decision` | Model decides based on description |
| `manual` | Only via @mention |
| `glob` | When working with matching files |

## Core Rules

| Rule | Trigger | Skills | Description |
|------|---------|--------|-------------|
| BRAIN_TO_DOCS | always_on | all | Draft in brain/, persist after approval |
| LANGUAGE_REQUIREMENTS | always_on | all | User language for drafts, English for persist |

## Protocol Rules

| Rule | Trigger | Skills | Description |
|------|---------|--------|-------------|
| GIT_PROTOCOL | model_decision | all | Branching, commits, atomic changes |
| DOCUMENT_STRUCTURE_PROTOCOL | model_decision | all | docs/ lifecycle: active/review/closed |
| TRACEABILITY_PROTOCOL | model_decision | @product-analyst, @bmad-architect, @tech-spec-writer, @qa-lead | Upstream docs, requirements checklist |
| TDD_PROTOCOL | model_decision | @backend-go-expert, @frontend-nuxt, @qa-lead, @debugger | Red-Green-Refactor, reproduction first |
| TECH_DEBT_PROTOCOL | model_decision | @backend-go-expert, @frontend-nuxt, @debugger | Track TODOs in TECH_DEBT.md |

## Team Rules

| Rule | Trigger | Skills | Description |
|------|---------|--------|-------------|
| PIPELINE | model_decision | all | Workflow from discovery to deployment |
| TEAM | model_decision | all | Skills roster for collaboration |

---

## TODO: Skill-Specific Targeting

> Отложено: нужно валидировать гипотезу что @skill-name в description активирует rule.

**Открытые вопросы:**
1. Добавлять ли @skill-name в description?
2. Как поддерживать синхронизацию при переименовании?
3. Нужен ли валидатор для проверки существования скиллов?
