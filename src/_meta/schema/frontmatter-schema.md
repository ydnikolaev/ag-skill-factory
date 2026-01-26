# Skill Frontmatter Schema

> Defines the extended frontmatter structure for SKILL.md files.
> Used by generate_skill_matrix.py and generate_pipelines.py.

## Required Fields

```yaml
---
name: skill-name                     # Human-readable name
description: Short description...    # One-line for TEAM.md
version: 1.0.0                       # Semver
---
```

## Pipeline Fields (Optional but Recommended)

```yaml
---
# Pipeline integration
phase: definition                    # discovery | definition | architecture | implementation | delivery | utility
category: analyst                    # factory | technical | analyst | utility

# Handoffs
receives_from:                       # Skills that provide input (list)
  - idea-interview
  - feature-fit
delegates_to:                        # Skills that receive output (list)
  - bmad-architect

# Artifacts produced
outputs:                             # What this skill creates
  - artifact: roadmap.md
    path: project/docs/active/product/
  - artifact: user-stories.md
    path: project/docs/active/specs/
---
```

## Phase Values

| Phase | Description | Typical Skills |
|-------|-------------|----------------|
| `discovery` | Initial requirements gathering | idea-interview, feature-fit |
| `definition` | Product specs and roadmap | product-analyst |
| `architecture` | Technical design | bmad-architect, tech-spec-writer |
| `implementation` | Coding | backend-go-expert, frontend-nuxt |
| `delivery` | Testing and deployment | qa-lead, devops-sre |
| `utility` | Cross-cutting support | debugger, project-bro, doc-janitor |

## Category Values

| Category | Description |
|----------|-------------|
| `factory` | Meta-skills for skill factory |
| `technical` | Implementation skills |
| `analyst` | Planning and design skills |
| `utility` | Support and tooling skills |

## Full Example

```yaml
---
name: product-analyst
description: Defines Vision, Roadmap, User Stories, and translates them into Technical Specs.
version: 1.2.0

phase: definition
category: analyst

receives_from:
  - idea-interview
  - feature-fit

delegates_to:
  - bmad-architect
  - tech-spec-writer

outputs:
  - artifact: roadmap.md
    path: project/docs/active/product/
  - artifact: user-stories.md
    path: project/docs/active/specs/
  - artifact: requirements.md
    path: project/docs/active/specs/
---
```

## Scripts Using This Schema

- `scripts/generate_skill_matrix.py` → `blueprint/_meta/skills/skill-matrix.yaml`
- `scripts/generate_pipelines.py` → `blueprint/_meta/pipelines/PIPELINE_<preset>.md`
- `scripts/generate_teams.py` → `blueprint/_meta/teams/TEAM_<preset>.md`
