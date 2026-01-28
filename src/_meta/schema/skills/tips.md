# Skill Authoring Tips

> Best practices and anti-patterns for skill authors.

## DOs ✅

### Identity

- **DO** write description as two sentences: WHAT + WHEN
  ```yaml
  description: Expert Go developer (1.25+) specializing in Clean Architecture. Activate when implementing Go services or APIs.
  ```

- **DO** use lowercase-with-hyphens for names
  ```yaml
  name: backend-go-expert  # ✅
  name: BackendGoExpert    # ❌
  ```

- **DO** use semver for versions
  ```yaml
  version: 1.3.0  # ✅
  version: v1.3   # ❌
  ```

### Workflow

- **DO** use leaf presets only
  ```yaml
  presets: [backend]     # ✅
  presets: [core]        # ❌ (core is parent, not leaf)
  ```

- **DO** ensure handoffs share presets
  ```yaml
  # If backend-go-expert has presets: [backend]
  # Then tech-spec-writer must also have [backend] or parent
  ```

### Documents

- **DO** use doc_types from doc-types.yaml
- **DO** follow path pattern: `project/docs/active/{category}/`

## DON'Ts ❌

### Identity

- **DON'T** include activation instructions in description
  ```yaml
  # ❌ Verbose activation
  description: "Use when user says 'implement backend' or asks about Go..."
  
  # ✅ Simple WHAT + WHEN
  description: "Expert Go developer. Activate for Go implementation tasks."
  ```

### Workflow

- **DON'T** create circular delegations
  ```yaml
  # ❌ A → B → A
  skill-a:
    delegates_to: [skill-b]
  skill-b:
    delegates_to: [skill-a]
  ```

- **DON'T** skip pipeline phases
  ```yaml
  # ❌ Discovery → Implementation (skips Definition, Architecture)
  # Follow natural pipeline order
  ```

### Documents

- **DON'T** create docs for other skills
  ```yaml
  # ❌ backend-go-expert creating user-stories
  # Only create docs you own
  ```

## Anti-Patterns

| Anti-Pattern | Problem | Solution |
|--------------|---------|----------|
| God Skill | Does everything | Split by phase/domain |
| Orphan Skill | No handoffs | Add receives_from/delegates_to |
| Preset Mismatch | Handoff skills in different presets | Align presets |
| Duplicate Docs | Multiple skills create same doc | Single owner |

## Validation Commands

```bash
# Validate single skill
make skill-validate SKILL=my-skill

# Validate all skills
make skill-validate

# Check handoffs
python3 scripts/validate_handoffs.py
```
