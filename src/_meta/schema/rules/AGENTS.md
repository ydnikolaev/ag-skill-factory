# Rule Schema

> Entry point for AI agents creating rules.

## Quick Start

1. Create new file in `src/rules/RULE_NAME.md`
2. Add YAML frontmatter with `trigger` and `description`
3. Write rule content with H1 title and H2 sections
4. Run `make build` to compile to `dist/`

## Rule Structure

```markdown
---
trigger: model_decision
description: Brief explanation for AI to decide when to apply
---

# Rule Name 🔧

Introduction explaining the rule purpose.

## Section 1

Rule content...

## Enforcement

What agents MUST do and MUST NOT do.
```

## Key Files

| File | Purpose |
|------|---------|
| `rule-schema.yaml` | Core schema |
| `enums.yaml` | Triggers, categories |

## Trigger Types

| Trigger | When Applied |
|---------|--------------|
| `always_on` | Every conversation |
| `glob` | When file patterns match |
| `model_decision` | AI decides based on context |
| `manual` | User explicitly invokes |

## Naming Convention

Use `SCREAMING_SNAKE_CASE.md`:
- `GIT_PROTOCOL.md`
- `TDD_PROTOCOL.md`
- `LANGUAGE_REQUIREMENTS.md`
