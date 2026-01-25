# UX Designer Checklist

## 🚨 Document Persistence (MANDATORY)

> [!CAUTION]
> **BEFORE handing off to another skill, you MUST:**

- [ ] **Final document exists in `docs/`** at the path defined in Artifact Ownership
- [ ] **ARTIFACT_REGISTRY.md updated** with status ✅ Done and Last Updated date
- [ ] **Artifact synced** — if you used an artifact, copy final content to `docs/`

**Why?** Artifacts don't persist between sessions. Without `docs/` file, the next skill cannot continue.

## Before Starting
- [ ] Brand guidelines gathered (or defined)
- [ ] Target platforms identified (web, mobile, TMA)

## Token Definition
- [ ] Color palette defined (primitives + semantic)
- [ ] Typography scale defined
- [ ] Spacing scale defined (4px/8px grid)
- [ ] Shadow/elevation tokens defined
- [ ] Border radius tokens defined

## Documentation
- [ ] `design-tokens.json` or `tokens.yaml` created
- [ ] `docs/design-system.md` written
- [ ] Component states documented (hover, focus, disabled)

## Handover
- [ ] Tokens provided to `@ui-implementor`
- [ ] Specs provided to `@frontend-nuxt`
- [ ] Design approved by user via `notify_user`
