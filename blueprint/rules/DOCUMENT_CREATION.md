---
trigger: model_decision
description: Creating new markdown documents not covered by existing templates. Use when no template exists in _meta/_docs/templates/documents/.
---

# Document Creation (No Template)

> Use this ONLY when creating a document that has no template in `_meta/_docs/templates/documents/`.

---

## When to Use

Check if template exists:
```bash
ls blueprint/_meta/_docs/templates/documents/_*.md
```

- **Template exists?** → Copy template, fill placeholders. Done.
- **No template?** → Follow this protocol.

---

## Step-by-Step

### 1. Get Work Unit

```bash
work_unit=$(git branch --show-current | sed 's|/|-|g')
# Example: feat/forum → feat-forum
```

### 2. Determine Path

```
project/docs/active/{category}/{work_unit}.md
```

Where `{category}` is one of:
- `discovery/` — briefs, research
- `product/` — user stories, roadmap
- `design/` — tokens, theming
- `architecture/` — context maps, contracts
- `specs/` — technical specifications
- `backend/` — service implementation
- `frontend/` — UI implementation
- `qa/` — test cases, reports
- `infrastructure/` — deployment
- `bugs/` — debug reports
- `refactoring/` — refactor plans
- `project/` — decision logs, issues

### 3. Create Frontmatter

```yaml
---
status: Draft
owner: @skill-name
work_unit: {WORK_UNIT}
created: YYYY-MM-DD
updated: YYYY-MM-DD
---
```

### 4. Create Content

```markdown
# Title: {WORK_UNIT}

## Overview

<!-- Brief description -->

---

## Details

<!-- Main content -->

---

## Checklist

- [ ] Item 1
- [ ] Item 2
```

---

## Rules

1. **NEVER** use generic filenames (e.g., `notes.md`, `draft.md`)
2. **ALWAYS** use `{work_unit}.md` as filename
3. **ALWAYS** include frontmatter with `status`, `owner`, `work_unit`
4. **ALWAYS** place in correct category folder
