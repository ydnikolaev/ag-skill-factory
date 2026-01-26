# Rules Template

Use this format when creating new rules for `.agent/rules/`.

```markdown
# <RULE_FILENAME>

> <One-line description: clear, actionable summary for AI agent>

<Table or structured content with rules>

<Additional context if needed>
```

## Example

```markdown
# BRAIN_TO_DOCS

> Draft in brain/, persist to project/docs/ only after user approval.

| Phase | Location | Lifetime |
|-------|----------|----------|
| Draft | `brain/` | Per-conversation |
| Persist | `project/docs/` | Permanent |
```

## Checklist

- [ ] H1 = exact filename (without .md)
- [ ] One-line description immediately after H1
- [ ] Copy to both `.agent/rules/` and `blueprint/rules/`
- [ ] Create matching section template in `.agent/templates/skill sections/`
- [ ] Update `section-matrix.yaml` with new section
- [ ] Update `validate_skill.py` if enforcement needed
