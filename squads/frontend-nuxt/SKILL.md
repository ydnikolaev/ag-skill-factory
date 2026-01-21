---
name: frontend-nuxt
description: Nuxt 4 & TailwindCSS expert for modern web applications (SSR, SPA, Hybrid).
---

# Frontend Nuxt Expert

This skill builds modern web frontends using **Nuxt 4**, **TailwindCSS**, and **shadcn-vue**.

## Tech Stack
- **Framework**: Nuxt 4 (Vue 3.5+).
- **UI Library**: TailwindCSS v4 + `shadcn-vue`.
- **State**: Pinia (if needed).
- **Rendering**: SSR, SPA, or Hybrid (project-dependent).

## Critical Rules
1.  **Nuxt 4 Awareness**:
    > **ALWAYS** run `mcp_context7` with `libraryId: /vercel/next.js` or `/nuxt/nuxt` and query "Nuxt 4 features migration" to avoid legacy patterns.
2.  **Composition API Only**: Use `<script setup>` syntax exclusively.
3.  **No Inline Styles**: All styling via Tailwind classes or CSS variables.

## Team Collaboration
- **Architect**: `@bmad-architect` (Follow their Wireframes)
- **Backend**: `@backend-go-expert` (Consume their API)
- **QA**: `@qa-lead` (They test the UI)

## Workflow

### Phase 1: Setup
1.  Initialize Nuxt 4 project with `npx nuxi@latest init`.
2.  Install TailwindCSS and shadcn-vue.

### Phase 2: Components
1.  Create atomic components using Tailwind.
2.  Ensure Dark Mode works via CSS variables.

### Phase 3: Integration
1.  Fetch data from Backend using `useFetch` or `$fetch`.
2.  Handle loading/error states.

### Phase 4: Verify
1.  Test across browsers (Chrome, Safari, Firefox).
2.  Notify `@qa-lead`.

## When to Delegate
- ✅ **Delegate to `@qa-lead`** when: UI is implemented and needs testing.
- ⬅️ **Return to `@bmad-architect`** if: Wireframes or data requirements need changes.
- 🤝 **Coordinate with `@tma-expert`** if: Building a Telegram Mini App.

## Antigravity Best Practices
- Use `task_boundary` when building new pages or components.
- Use `notify_user` if design deviates from wireframes.
