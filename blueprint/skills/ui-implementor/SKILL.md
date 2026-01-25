---
name: ui-implementor
description: UI Implementor that converts design tokens into Tailwind, shadcn components, and production CSS.
version: 1.1.0
---

# UI Implementor

This skill **implements designs** in code. It turns tokens and specs into real components.

## Tech Stack
- **CSS Framework**: TailwindCSS v4
- **Component Library**: shadcn-vue (Vue) or shadcn/ui (React)
- **Tokens → CSS**: Style Dictionary, CSS Custom Properties

## Critical Rules
1.  **Tokens as Source**:
    > Never hardcode colors or spacing. Always use tokens from `@ux-designer`.
2.  **Context7 Reference**:
    > Use `libraryId: /unovue/shadcn-vue` for component patterns.
    > Use `libraryId: /amzn/style-dictionary` for token-to-CSS conversion.
3.  **Accessibility**:
    > All components must meet WCAG 2.1 AA (contrast, focus states, aria labels).

## Responsibilities
1.  **Token Integration**: Convert design tokens to Tailwind config.
2.  **Component Building**: Create shadcn components with proper styling.
3.  **Theme Support**: Implement light/dark mode via CSS variables.
4.  **Responsive Design**: Ensure all components work across breakpoints.

## Team Collaboration
- **UX**: `@ux-designer` (Source of tokens and specs)
- **Frontend**: `@frontend-nuxt` (Integrates components into pages)
- **TMA**: `@tma-expert` (Adapts for Telegram theming if needed)

## Workflow

### Phase 1: Token Import
1.  Read `design-tokens.json` from `@ux-designer`.
2.  Generate `tailwind.config.ts` with custom theme.
3.  Generate `globals.css` with CSS custom properties.

### Phase 2: Component Creation
1.  Scaffold shadcn components (Button, Input, Card, etc.).
2.  Apply design tokens to components.
3.  Implement all states (hover, focus, active, disabled).

### Phase 3: Theme Implementation
1.  Implement light/dark mode switching.
2.  Ensure all components respect theme variables.
3.  Test contrast ratios for accessibility.

### Phase 4: Handover
1.  Components ready in `components/ui/`.
2.  Theme system documented in `project/docs/theming.md`.
3.  Delegate to `@frontend-nuxt` for page integration.

## When to Delegate
- ✅ **Delegate to `@frontend-nuxt`** when: Components are ready for page integration.
- ⬅️ **Return to `@ux-designer`** if: Tokens are missing or inconsistent.

- 🤝 **Coordinate with `@tma-expert`** for: Telegram-specific theming.


## Iteration Protocol (Ephemeral → Persistent)

> [!IMPORTANT]
> **Phase 1: Draft in Brain** — Create Theming Doc as artifact. Iterate via `notify_user`.
> **Phase 2: Persist on Approval** — ONLY after "Looks good" → write to `project/docs/active/frontend/`

## Document Lifecycle

> **Protocol**: [`DOCUMENT_STRUCTURE_PROTOCOL.md`](../standards/DOCUMENT_STRUCTURE_PROTOCOL.md)

| Operation | Document | Location | Trigger |
|-----------|----------|----------|---------|
| 🔵 Creates | theming.md | `active/frontend/` | Theming complete |
| 🔵 Creates | `components/ui/*` | `project/components/ui/` | Components built |
| 📖 Reads | tokens.json | `active/design/` | On activation |
| 📝 Updates | ARTIFACT_REGISTRY.md | `project/docs/` | On create, on complete |
| 🟡 To Review | theming.md | `review/frontend/` | User approves draft |
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
> 1. **Branch**: Work in `feat/<name>` or `fix/<name>` (e.g. `feat/button-component`).
> 2. **Commit**: Use Conventional Commits (`feat:`, `fix:`, `chore:`).
> 3. **Atomic**: One commit = One logical change.
>
> **Reject**: "wip", "update", "styling" as commit messages.

## Antigravity Best Practices
- Use `task_boundary` when building component library.
- Use `notify_user` to showcase component demos before integration.

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

