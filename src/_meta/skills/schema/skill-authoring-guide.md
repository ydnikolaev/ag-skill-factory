# Skill Authoring Guide

> Step-by-step instructions for creating and updating Antigravity skills.
> SSOT for @skill-creator and @skill-updater.

## Key Files (from project root)

| Purpose | Path |
|---------|------|
| Entry point | `src/_meta/skills/schema/AGENTS.md` |
| Core schema | `src/_meta/skills/schema/skill-schema.yaml` |
| Factory enums | `src/_meta/skills/schema/enums/factory.yaml` |
| Runtime enums | `src/_meta/skills/schema/enums/runtime.yaml` |
| Template | `src/_meta/skills/schema/examples/_template.yaml` |
| Preset hierarchy | `src/_meta/preset-hierarchy.yaml` |
| Skill matrix | `src/_meta/skill-matrix.yaml` |
| Doc types | `src/_meta/doc-types.yaml` |
| Document flow | `project/docs/DOCUMENT_FLOW.md` |
| Pipelines | `src/_meta/pipelines/PIPELINE_{preset}.md` |

---

## Creating a New Skill

### Step 1: Preparation

1. Read `src/_meta/skills/schema/AGENTS.md`
2. Check `src/_meta/skill-matrix.yaml` — existing skills and handoffs
3. Check `src/_meta/pipelines/PIPELINE_{preset}.md` — which phase?
4. Verify skill doesn't exist in `src/skills/`

### Step 2: Identity

1. Copy `src/_meta/skills/schema/examples/_template.yaml` to `src/skills/{name}/skill.yaml`
2. Set `name` — lowercase-with-hyphens, unique
3. Write `description`:
   - First sentence: WHAT it does
   - Second sentence: WHEN to use it
4. Set `phase` — from `src/_meta/skills/schema/enums/factory.yaml#phases`
5. Set `category` — from `src/_meta/skills/schema/enums/factory.yaml#categories`

### Step 3: Capabilities

1. Set `mcp_servers` — context7, sky-cli, github
2. Set `dependencies` — go1.25, python3, docker
3. Define `context.required` — files that MUST exist
4. Define `context.optional` — nice to have

### Step 4: Workflow

1. Set `presets` — leaf presets only (check `src/_meta/preset-hierarchy.yaml`)
2. Define `receives_from` — check `src/_meta/skill-matrix.yaml#handoffs`
3. Define `delegates_to` — check `src/_meta/skill-matrix.yaml#handoffs`
4. Define `return_paths` if applicable

### Step 5: Documents

1. Define `requires` — check `src/_meta/doc-types.yaml`
2. Define `creates`:
   - doc_type from `src/_meta/doc-types.yaml`
   - path: `project/docs/active/{category}/`
   - lifecycle: `per-feature` | `living` | `per-work-unit`
3. Define `updates` if applicable
4. Define `archives` — usually only @doc-janitor

### Step 6: Validation

1. Set `protocols` — from `src/_meta/skills/schema/enums/runtime.yaml#protocols`
2. Set `checks` — from `src/_meta/skills/schema/enums/runtime.yaml#checks`

### Step 7: Generate & Validate

```bash
make skill-generate SKILL=skill-name
make skill-validate SKILL=skill-name
make presets-rebuild
git add src/skills/skill-name/
git commit -m "feat(skills): add skill-name"
```

---

## Updating an Existing Skill

1. Read `src/skills/{name}/SKILL.md`
2. Update YAML frontmatter
3. Run `make skill-validate SKILL={name}`
4. If workflow changed → `make presets-rebuild`
5. Commit: `fix(skills): description`

---

## Commands

| Command | Purpose |
|---------|---------|
| `make skill-validate` | Validate all skills |
| `make skill-validate SKILL=name` | Validate single skill |
| `make presets-rebuild` | Regenerate presets.yaml |
| `make generate-all` | Regenerate all SKILL.md |
