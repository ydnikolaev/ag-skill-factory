# Frontend Nuxt TMA Checklist

## Before Starting
- [ ] Read UI specs from `@bmad-architect`
- [ ] Check Nuxt 4 features via `mcp_context7` query: "Nuxt 4 features"

## TMA Specifics
- [ ] Uses `@tma.js/sdk` or `telegram-web-app.js`
- [ ] Colors from `Telegram.WebApp.themeParams`
- [ ] Respects `viewportHeight` for layouts
- [ ] Back Button handled via SDK (not browser history)

## UI Quality
- [ ] Dark mode works correctly
- [ ] No horizontal scroll on mobile
- [ ] Loading states implemented
- [ ] Error states implemented

## Integration
- [ ] API calls to Backend work
- [ ] `initData` passed in Authorization header

## Testing
- [ ] Tested on Telegram Desktop
- [ ] Tested on Telegram Mobile (iOS/Android)
- [ ] Notified `@qa-lead`
