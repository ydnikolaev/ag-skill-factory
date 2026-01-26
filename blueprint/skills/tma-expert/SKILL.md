---
name: tma-expert
description: Expert in Telegram Mini Apps (TMA) using @tma.js/sdk for native Telegram integration.
version: 1.2.0

phase: implementation
category: technical

receives_from:
  - telegram-mechanic
  - frontend-nuxt

delegates_to:
  - qa-lead

outputs:
  - artifact: tma-config.md
    path: project/docs/active/frontend/
    doc_category: frontend
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

## TDD Protocol (Hard Stop)

> [!CAUTION]
> **NO CODE WITHOUT FAILING TEST.**
> - **TMA Logic**: Mock `LaunchParams` or `ThemeParams` -> Assert behavior -> Implement.
> - **Do not skip tests** just because it is "UI glue code".
>
> **Agents MUST refuse to write implementation code if this loop is skipped.**

## Tech Debt Protocol (Hard Stop)

> [!CAUTION]
> **Follow `../standards/TECH_DEBT_PROTOCOL.md`.**
> When creating workarounds:
> 1. Add `// TODO(TD-XXX): description` in code
> 2. Register in `project/docs/TECH_DEBT.md`
>
> **Forbidden:** Untracked TODOs, undocumented hardcoded values.

## Git Protocol (Hard Stop)

> [!CAUTION]
> **Follow `../standards/GIT_PROTOCOL.md`.**
> 1. **Branch**: Work in `feat/<name>` or `fix/<name>`.
> 2. **Commit**: Use Conventional Commits (`feat:`, `fix:`).
> 3. **Atomic**: One commit = One logical change.
>
> **Reject**: "wip", "update", "fix" as commit messages.

## Language Requirements

> All skill files must be in English. See [LANGUAGE.md](file://blueprint/rules/LANGUAGE.md).

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


## Iteration Protocol (Ephemeral → Persistent)

> [!IMPORTANT]
> **Phase 1: Draft in Brain** — Create TMA Config as artifact. Iterate via `notify_user`.
> **Phase 2: Persist on Approval** — ONLY after "Looks good" → write to `project/docs/active/tma/`

## Document Lifecycle

> **Protocol**: [`DOCUMENT_STRUCTURE_PROTOCOL.md`](../standards/DOCUMENT_STRUCTURE_PROTOCOL.md)

| Operation | Document | Location | Trigger |
|-----------|----------|----------|---------|
| 🔵 Creates | tma-config.md | `active/tma/` | TMA setup complete |
| 📖 Reads | ui-implementation.md | `active/frontend/` | On activation |
| 📖 Reads | webhook-config.md | `active/bot/` | For deep links |
| 📝 Updates | ARTIFACT_REGISTRY.md | `project/docs/` | On create, on complete |
| 🟡 To Review | tma-config.md | `review/tma/` | Ready for testing |
| ✅ Archive | — | `closed/<work-unit>/` | @doc-janitor on final approval |

## Pre-Handoff Validation (Hard Stop)

> [!CAUTION]
> **MANDATORY self-check before `notify_user` or delegation.**

| # | Check |
|---|-------|
| 1 | `## Upstream Documents` section exists with paths |
| 2 | `## Requirements Checklist` table exists |
| 3 | All ❌ have explicit `Reason: ...` |
| 4 | Document in `review/` folder |
| 5 | `ARTIFACT_REGISTRY.md` updated |

**If ANY unchecked → DO NOT PROCEED.**

## Handoff Protocol


> [!CAUTION]
> **BEFORE handoff:**
> 1. Save final document to `project/docs/` path
> 2. Change file status from `Draft` to `Approved` in header/frontmatter
> 3. Update `project/docs/ARTIFACT_REGISTRY.md` status to ✅ Done
> 4. Use `notify_user` for final approval
> 5. THEN delegate to next skill

## Antigravity Best Practices
- Use `task_boundary` when adding TMA features to existing Nuxt app.
- Use `notify_user` to confirm deep link format with user.

> [!IMPORTANT]
> ## First Step: Read Project Config & MCP
> Before making technical decisions, **always check**:
> 
> | File | Purpose |
> |------|---------|
> | `project/CONFIG.yaml` | Stack versions, modules, architecture |
> | `mcp.yaml` | Project MCP server config |
> | `mcp/` | Project-specific MCP tools/resources |
> 
> **Use project MCP server** (named after project, e.g. `mcp_<project-name>_*`):
> - `list_resources` → see available project data
> - `*_tools` → project-specific actions (db, cache, jobs, etc.)
> 
> **Use `mcp_context7`** for library docs:
> - Check `mcp.yaml → context7.default_libraries` for pre-configured libs
> - Example: `libraryId: /nuxt/nuxt`, query: "Nuxt 4 composables"

