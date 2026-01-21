---
name: tma-expert
description: Expert in Telegram Mini Apps (TMA) using @tma.js/sdk for native Telegram integration.
---

# TMA Expert (Telegram Mini Apps)

This skill specializes in building **Telegram Mini Apps** that feel native to the Telegram ecosystem.

## Tech Stack
- **SDK**: `@tma.js/sdk` (v3.x) — the modern TypeScript SDK.
- **Vue Integration**: `@tma.js/sdk-vue` for Vue/Nuxt projects.
- **Alternative**: `telegram-web-app.js` (legacy, avoid if possible).

## Critical Rules
1.  **TMA SDK Awareness**:
    > **ALWAYS** run `mcp_context7` with:
    > `libraryId: /telegram-mini-apps/tma.js`, query: "SDK initialization Vue launch params theme"
2.  **Native Feel**:
    - Use `bindThemeParamsCSSVars()` for Telegram theme colors.
    - Use `bindViewportCSSVars()` for responsive height.
    - Handle Back Button via SDK, NOT browser history.
3.  **initData Security**: Pass `initData` to backend for HMAC-SHA256 validation.

## Key Concepts

### Launch Parameters
```typescript
import { useLaunchParams } from '@tma.js/sdk-vue';
const lp = useLaunchParams();
console.log(lp.startParam); // Deep link parameter
```

### Theme Colors (CSS Variables)
```typescript
import { bindThemeParamsCSSVars, initThemeParams } from '@telegram-apps/sdk';
const tp = initThemeParams();
bindThemeParamsCSSVars(tp);
// Now use: var(--tg-theme-bg-color), var(--tg-theme-text-color)
```

### Viewport Height
```typescript
import { bindViewportCSSVars, initViewport } from '@telegram-apps/sdk';
const vp = await initViewport();
bindViewportCSSVars(vp);
// Now use: var(--tg-viewport-height), var(--tg-viewport-stable-height)
```

## Team Collaboration
- **Frontend**: `@frontend-nuxt` (Base Nuxt setup, you add TMA layer)
- **Telegram**: `@telegram-mechanic` (Bot setup, deep links)
- **Backend**: `@backend-go-expert` (initData validation)

## Workflow

### Phase 1: Setup
1.  Start with `@frontend-nuxt` for base Nuxt 4 project.
2.  Install: `npm install @tma.js/sdk @tma.js/sdk-vue`.
3.  Initialize SDK in app entry point.

### Phase 2: Theming
1.  Bind CSS variables for theme and viewport.
2.  Replace hardcoded colors with `var(--tg-theme-*)`.

### Phase 3: Native Features
1.  Implement MainButton, BackButton if needed.
2.  Handle haptic feedback (`HapticFeedback.impactOccurred`).

### Phase 4: Verify
1.  Test on Telegram Desktop (Web).
2.  Test on Telegram Mobile (iOS + Android).

## When to Delegate
- ⬅️ **Return to `@frontend-nuxt`** if: Generic UI work without TMA specifics.
- 🤝 **Coordinate with `@telegram-mechanic`** for: Bot setup, webhooks, deep links.
- ✅ **Delegate to `@qa-lead`** when: Ready for testing.

## Antigravity Best Practices
- Use `task_boundary` when adding TMA features to existing Nuxt app.
- Use `notify_user` to confirm deep link format with user.
