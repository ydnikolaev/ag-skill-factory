# UI Implementor Checklist

## 🚨 Document Persistence (MANDATORY)

> [!CAUTION]
> **BEFORE handing off to another skill, you MUST:**

- [ ] **Final document exists in `docs/`** at the path defined in Artifact Ownership
- [ ] **ARTIFACT_REGISTRY.md updated** with status ✅ Done and Last Updated date
- [ ] **Artifact synced** — if you used an artifact, copy final content to `docs/`

**Why?** Artifacts don't persist between sessions. Without `docs/` file, the next skill cannot continue.

## Before Starting
- [ ] Design tokens received from `@ux-designer`
- [ ] Component specs reviewed

## Tailwind Setup
- [ ] `tailwind.config.ts` extended with custom theme
- [ ] CSS custom properties in `globals.css`
- [ ] Color palette uses token values

## Components
- [ ] Core components created (Button, Input, Card)
- [ ] All states implemented (hover, focus, active, disabled)
- [ ] Components use tokens (no hardcoded values)

## Theming
- [ ] Light mode works correctly
- [ ] Dark mode works correctly
- [ ] Theme switching implemented

## Accessibility
- [ ] Color contrast meets WCAG AA (4.5:1 text, 3:1 UI)
- [ ] Focus states visible
- [ ] ARIA labels where needed

## Handover
- [ ] Components in `components/ui/`
- [ ] `docs/theming.md` written
- [ ] Notified `@frontend-nuxt` for integration
