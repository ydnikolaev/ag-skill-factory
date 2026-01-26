# Frontend Nuxt Checklist

## 🚨 Document Persistence (MANDATORY)

> [!CAUTION]
> **BEFORE handing off to another skill, you MUST:**

- [ ] **Final document exists in `docs/`** at the path defined in Artifact Ownership
- [ ] **ARTIFACT_REGISTRY.md updated** with status ✅ Done and Last Updated date
- [ ] **Artifact synced** — if you used an artifact, copy final content to `docs/`

**Why?** Artifacts don't persist between sessions. Without `docs/` file, the next skill cannot continue.

## Before Starting
- [ ] Read UI specs from Architect
- [ ] Check Nuxt 4 features via `mcp_context7`

## Code Quality
- [ ] Using `<script setup>` Composition API
- [ ] No inline styles (Tailwind only)
- [ ] Components are atomic and reusable

## UI Quality
- [ ] Dark mode works correctly
- [ ] Responsive on mobile/tablet/desktop
- [ ] Loading states implemented
- [ ] Error states implemented

## Integration
- [ ] API calls to Backend work
- [ ] Data fetching uses `useFetch` or `$fetch`

## Testing
- [ ] Tested on Chrome, Safari, Firefox
- [ ] Notified `@qa-lead`
