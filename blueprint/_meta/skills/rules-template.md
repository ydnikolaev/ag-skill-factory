# Rules Template

Use this format when creating new rules for `.agent/rules/` or `blueprint/rules/`.

## Frontmatter (Required)

```yaml
---
trigger: always_on | manual | model_decision | glob
description: Clear, actionable summary for AI agent decision
glob: "*.go"  # Only for trigger: glob
---
```

### Trigger Values

| Trigger | When Applied | Use Case |
|---------|--------------|----------|
| `always_on` | Every conversation | Core rules (BRAIN_TO_DOCS, LANGUAGE) |
| `model_decision` | Model decides by description | Context-specific protocols |
| `manual` | Only via @mention | Rarely used rules |
| `glob` | When working with matching files | File-type rules (*.go, *.ts) |

## Full Example

```markdown
---
trigger: model_decision
description: Test-first development with Red-Green-Refactor. Apply when writing code or fixing bugs.
---

# TDD_PROTOCOL

> No implementation code without a failing test.

## The Golden Rule

1. **RED**: Write failing test
2. **GREEN**: Write minimal code
3. **REFACTOR**: Clean up
```


<!-- TODO:  Подумать над чеклистами в rules, здесь или отдельно вынести и какие должны быть -->


## Checklist

- [ ] YAML frontmatter with `trigger` and `description`
- [ ] H1 = exact filename (without .md)
- [ ] One-line quote summary after H1
- [ ] Add to `blueprint/rules/`
- [ ] Update `_meta/rules/rules-matrix.md`
