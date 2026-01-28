# SSOT Architecture Policy

> **Version**: 1.0
> **Status**: Approved
> **Scope**: All Factory schema and validation systems

## Core Principle

```
REAL FILES are the Single Source of Truth (SSOT).
Schemas DESCRIBE expected format, they do not DEFINE truth.
```

## SSOT Hierarchy

| Entity | SSOT Location | Schema Role |
|--------|---------------|-------------|
| Skills | `src/skills/*/SKILL.md` | Describes frontmatter format |
| Rules | `src/rules/*.md` | Describes frontmatter format |
| Workflows | `src/workflows/*.md` | Describes frontmatter format |
| Documents | `src/templates/documents/_*.md` | Describes frontmatter format |
| Folder Structure | `src/templates/folder-structure/` | Describes expected structure |
| **Enums** | `src/_meta/schema/*/enums/*.yaml` | **SSOT for allowed values** |

## Hybrid Approach

### Implementation-First (Real Files = SSOT)

Used for **content** that is manually created and frequently edited:

```
SKILL.md files → validators check → report errors
```

**Why**: Skills, rules, workflows are authored by humans. Schema follows implementation.

### Schema-First (Schema = SSOT)

Used for **constraints** that define allowed values:

```
enums/*.yaml → validators check → reject invalid values
```

**Enums are Schema-First because**:
- They define allowed `phases`, `categories`, `presets`
- Changes require deliberate decision
- New values must be added to schema FIRST

## Validation Flow

```
┌─────────────────────────────────────────────────────┐
│  ENUMS (Schema-First)                               │
│  → Allowed values defined here                      │
│  → Add new enum BEFORE using in files               │
└──────────────────────────┬──────────────────────────┘
                           │ validates against
┌──────────────────────────▼──────────────────────────┐
│  REAL FILES (Implementation-First)                  │
│  → Content created here                             │
│  → Must use values from enums                       │
└──────────────────────────┬──────────────────────────┘
                           │ validated by
┌──────────────────────────▼──────────────────────────┐
│  VALIDATORS                                         │
│  → Check real files against schema + enums          │
│  → Report errors if mismatch                        │
└─────────────────────────────────────────────────────┘
```

## Rules for Contributors

### ✅ DO

1. **Add new enum values first** before using them in files
2. **Run `make schema-validate-all`** after changes
3. **Treat folder-structure/ as SSOT** for document categories
4. **Update validators** when adding new validation rules

### ❌ DON'T

1. **DON'T change schema to match broken files** — fix the files
2. **DON'T add categories to documents** without adding folder in `folder-structure/`
3. **DON'T claim schema is SSOT** in comments (use "Describes format")
4. **DON'T bypass validators** — all changes must pass validation

## Cross-Validation

Structure validator ensures consistency:

```
folder-structure/active/    ←→    document enums/categories
         ↑                                  ↑
         └────── skill creates.path ────────┘
```

All three must align. Validator enforces this.

## Adding New Entities

### New Category

1. Create folder in `src/templates/folder-structure/active/{category}/`
2. Add to `src/_meta/schema/documents/enums/enums.yaml`
3. Run `make schema-validate-all`

### New Enum Value

1. Add to appropriate `enums/*.yaml`
2. Use in files
3. Run `make schema-validate-all`

### New Schema Field

1. Add to `*-schema.yaml`
2. Regenerate JSON: `make json-schema`
3. Update validators if needed
4. Run `make schema-validate-all`

## Enforcement

- **CI/CD**: `make schema-validate-all` runs on every PR
- **Pre-commit**: Recommended to run locally
- **Code Review**: Reviewers check SSOT compliance
