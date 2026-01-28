# Skill Schema V3

> Entry point for AI agents creating Antigravity skills.

## Quick Start

1. Copy `examples/_template.yaml` as base
2. Fill `identity` — name, description, phase, category
3. Fill `capabilities` — mcp_servers, dependencies (or use defaults)
4. Define `workflow` — presets, receives_from, delegates_to
5. Define `documents` — what you create/require
6. Define `validation` — protocols and checks (or use defaults)
7. Run `make skill-validate` to check

## Commands

| Command | Purpose |
|---------|---------|
| `make skill-validate` | Validate all skills |
| `make skill-validate SKILL=name` | Validate single skill |
| `make presets-rebuild` | Regenerate presets.yaml |
| `make generate-all` | Regenerate all artifacts |

## Key Files

| Purpose | Path |
|---------|------|
| Core schema | `skill-schema.yaml` |
| Factory enums | `enums/factory.yaml` |
| Runtime enums | `enums/runtime.yaml` |
| Section definitions | `definitions/section-definitions.yaml` |
| Protocol definitions | `definitions/protocol-definitions.yaml` |
| Category defaults | `defaults/category-defaults.yaml` |
| Template | `examples/_template.yaml` |
| Example | `examples/backend-go-expert.yaml` |

## When to Read Each File

| Task | File |
|------|------|
| Creating new skill | `examples/_template.yaml`, `skill-schema.yaml` |
| Understanding workflow | `src/_meta/skill-matrix.yaml` |
| Checking valid enums | `enums/factory.yaml`, `enums/runtime.yaml` |
| Understanding sections | `definitions/section-definitions.yaml` |
| Understanding protocols | `definitions/protocol-definitions.yaml` |
| Modifying schema | `schema-governance.md` |
| Best practices | `tips.md` |
| Step-by-step guide | `skill-authoring-guide.md` |

## Related Docs

- [DOCUMENT_FLOW.md](../../../project/docs/DOCUMENT_FLOW.md) — Document lifecycle
- [preset-hierarchy.yaml](../../preset-hierarchy.yaml) — Preset inheritance
- [skill-matrix.yaml](../../skill-matrix.yaml) — Skill handoffs
- [schema-governance.md](schema-governance.md) — How to change schema
- [tips.md](tips.md) — Best practices
- [skill-authoring-guide.md](skill-authoring-guide.md) — Step-by-step

## Validation Checklist

- [ ] `identity.description` explains WHAT and WHEN
- [ ] `identity.name` is lowercase-with-hyphens
- [ ] `identity.phase` is valid (from factory.yaml)
- [ ] `identity.category` is valid (from factory.yaml)
- [ ] `workflow.presets` are leaf presets only
- [ ] All handoffs share at least one preset
- [ ] `documents.creates.doc_type` exists in doc-types.yaml
- [ ] Run `make skill-validate` passes
