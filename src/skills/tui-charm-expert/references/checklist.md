# TUI Charm Expert Checklist

## 🚨 Document Persistence (MANDATORY)

> [!CAUTION]
> **BEFORE handing off to another skill, you MUST:**

- [ ] **Final document exists in `docs/`** at the path defined in Artifact Ownership
- [ ] **ARTIFACT_REGISTRY.md updated** with status ✅ Done and Last Updated date
- [ ] **Artifact synced** — if you used an artifact, copy final content to `docs/`

**Why?** Artifacts don't persist between sessions. Without `docs/` file, the next skill cannot continue.

## Model
- [ ] State struct defined
- [ ] Initial state set correctly

## Update
- [ ] All key events handled
- [ ] Window resize handled (`tea.WindowSizeMsg`)
- [ ] Quit gracefully on `ctrl+c`

## View
- [ ] Lipgloss styles defined
- [ ] Layout responsive to terminal size
- [ ] Colors work in light/dark terminals

## Integration
- [ ] Model integrates with Cobra command
- [ ] Loading states shown for async ops
- [ ] Error states displayed clearly
