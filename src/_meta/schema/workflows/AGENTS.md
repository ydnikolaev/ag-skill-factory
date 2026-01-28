# Workflow Schema

> Entry point for AI agents creating workflows.

## Quick Start

1. Create new file in `src/workflows/workflow-name.md`
2. Add YAML frontmatter with `description`
3. Write workflow steps with H2 "Steps" section
4. Add `// turbo` annotations for auto-run commands
5. Run `make build` to compile to `dist/`

## Workflow Structure

```markdown
---
description: Short description for help text
---

# /workflow-name Workflow

Brief explanation of what this workflow does.

## Steps

### 1. First Step

Explanation of this step.

// turbo
\`\`\`bash
command-to-auto-run
\`\`\`

### 2. Second Step

Next step...
```

## Key Files

| File | Purpose |
|------------|--------|
| `workflow-schema.yaml` | Core schema |
| `enums/factory.yaml` | workflow_categories |
| `enums/runtime.yaml` | turbo_modes |

## Turbo Annotations

| Annotation | Effect |
|------------|--------|
| `// turbo` | Auto-run the NEXT command block |
| `// turbo-all` | Auto-run ALL commands in workflow |

## Naming Convention

Use `lowercase-with-hyphens.md`:
- `refactor.md` → `/refactor`
- `new-feature.md` → `/new-feature`
- `doc-cleanup.md` → `/doc-cleanup`

## Invoking Workflows

Users invoke via slash command matching filename:
```
/refactor      → runs refactor.md
/new-feature   → runs new-feature.md
```
