# Skill Specification Template

Use this template when writing skill specifications.

---

```markdown
# Skill Specification: [name]

## Identity
- **Name**: `[name]`
- **Emoji**: [emoji]
- **One-liner**: [One sentence description]

## Trigger Phrases
- "[phrase 1]"
- "[phrase 2]"
- "[phrase 3]"

## Workflow

### Phase 1: [Name]
[Description of what happens]

### Phase 2: [Name]
[Description of what happens]

### Phase 3: [Name]
[Description of what happens]

## Boundaries

### DOES ✅
- [Responsibility 1]
- [Responsibility 2]
- [Responsibility 3]

### DOES NOT ❌
- [Exclusion 1] → Delegate to `@[skill]`
- [Exclusion 2] → Delegate to `@[skill]`

## Team Collaboration
- **Receives from**: `@[skill-name]` — [what it receives]
- **Passes to**: `@[skill-name]` — [what it passes]

## Artifacts
- **Creates**: `project/docs/[path]`
- **Reads**: `[files or paths]`
- **Updates**: `project/docs/ARTIFACT_REGISTRY.md`

## Technical Requirements
- **Scripts needed**: [Yes/No, describe if yes]
- **Templates needed**: [Yes/No, describe if yes]
- **References needed**: [Yes/No, describe if yes]

## Open Questions
- [Question 1, if any]
- [Question 2, if any]

---

**Status**: Ready for `@skill-creator`
```

---

## Notes

- Fill ALL sections before handing off to skill-creator
- Open Questions should be empty or resolved before handoff
- Be specific in boundaries — vague boundaries cause confusion
