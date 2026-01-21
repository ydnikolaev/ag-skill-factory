# TUI Charm Expert Checklist

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
