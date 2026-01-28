# Structure Schema

> Entry point for understanding project structure.

## Factory Source Structure

```
src/
├── skills/              # Skill sources
│   └── {name}/SKILL.md
├── rules/               # Rule sources
│   └── SCREAMING_SNAKE.md
├── workflows/           # Workflow sources
│   └── lowercase-hyphen.md
├── templates/
│   ├── documents/       # Doc templates
│   │   └── _doc-type.md
│   └── projects/        # Project scaffolds
├── sections/            # Shared includes
│   ├── protocols/
│   └── triggers/
└── _meta/               # Metadata
    ├── schema/          # ALL SCHEMAS HERE
    │   ├── skills/
    │   ├── documents/
    │   ├── rules/
    │   ├── workflows/
    │   └── structure/
    ├── pipelines/
    └── teams/
```

## Installed Project Structure

After `factory install`:

```
project/
├── .agent/
│   ├── rules/           # Compiled rules
│   ├── skills/          # Compiled skills
│   └── workflows/       # Compiled workflows
└── project/
    ├── docs/
    │   ├── active/      # Working docs
    │   ├── review/      # Pending review
    │   └── closed/      # Archived
    ├── planning/
    └── userdir/
```

## Key Files

| File | Purpose |
|------|---------|
| `structure-schema.yaml` | Full structure definition |

## Document Categories

Active docs organized by category:
- `architecture/` — System design
- `backend/` — Backend docs
- `frontend/` — Frontend docs
- `design/` — UI/UX design
- `product/` — Product requirements
- `specs/` — Technical specs
