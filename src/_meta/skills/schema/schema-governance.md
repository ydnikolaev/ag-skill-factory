# Schema Governance

> How to safely modify the Skill Schema V3.

## Change Categories

| Category | Impact | Approval |
|----------|--------|----------|
| **Patch** | Add optional field | None |
| **Minor** | Add enum value | Team review |
| **Major** | Change required field, remove field | User approval |

## Before Making Changes

1. Check if change is really needed
2. Identify impact category
3. For Major: document migration path

## Change Workflow

### Adding Optional Field

```yaml
# 1. Add to skill-schema.yaml
capabilities:
  new_field:
    type: string
    required: false
    description: What it does
```

No validation changes needed.

### Adding Enum Value

```yaml
# 1. Add to appropriate enum file
# enums/factory.yaml or enums/runtime.yaml
phases:
  - existing
  - new_phase  # Added
```

```bash
# 2. Regenerate artifacts
make generate-all
```

### Changing Required Field

1. Document in CHANGELOG
2. Update all existing skills
3. Update validators
4. Get user approval

## Impact Analysis

| File Changed | Affected |
|--------------|----------|
| `skill-schema.yaml` | All skills, validators |
| `enums/factory.yaml` | Skill creation, pipelines |
| `enums/runtime.yaml` | Pre-handoff validation |
| `definitions/` | Skill generation, validation |
| `defaults/` | New skills only |

## Rollback

Keep backup:
```bash
cp -r src/_meta/skills/schema tmp_schema_backup
```

## Validation After Change

```bash
make skill-validate
make presets-rebuild
python3 scripts/validate_handoffs.py
```
