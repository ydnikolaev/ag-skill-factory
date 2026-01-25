# TMA Expert Checklist

## 🚨 Document Persistence (MANDATORY)

> [!CAUTION]
> **BEFORE handing off to another skill, you MUST:**

- [ ] **Final document exists in `docs/`** at the path defined in Artifact Ownership
- [ ] **ARTIFACT_REGISTRY.md updated** with status ✅ Done and Last Updated date
- [ ] **Artifact synced** — if you used an artifact, copy final content to `docs/`

**Why?** Artifacts don't persist between sessions. Without `docs/` file, the next skill cannot continue.

## Before Starting
- [ ] Base Nuxt project ready via `@frontend-nuxt`
- [ ] Checked latest TMA SDK via `mcp_context7`

## SDK Setup
- [ ] `@tma.js/sdk` and `@tma.js/sdk-vue` installed
- [ ] SDK initialized in app entry point

## Theming
- [ ] `bindThemeParamsCSSVars()` called
- [ ] `bindViewportCSSVars()` called
- [ ] No hardcoded colors (using CSS variables)

## Native Features
- [ ] Back Button handled via SDK (not browser)
- [ ] MainButton configured (if needed)
- [ ] Haptic feedback used appropriately

## Security
- [ ] `initData` passed to Backend for validation
- [ ] No sensitive data stored client-side

## Testing
- [ ] Tested on Telegram Desktop (Web)
- [ ] Tested on Telegram Mobile (iOS)
- [ ] Tested on Telegram Mobile (Android)
- [ ] Notified `@qa-lead`
