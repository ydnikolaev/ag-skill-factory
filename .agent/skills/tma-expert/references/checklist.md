# TMA Expert Checklist

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
