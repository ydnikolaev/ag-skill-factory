---
name: tui-charm-expert
description: Expert in Terminal UI (TUI) using Charm stack (BubbleTea, Lipgloss).
---

# TUI Charm Expert

This skill makes the terminal beautiful using BubbleTea and Lipgloss.

## Tech Stack
- **Framework**: `bubbletea` (The Elm Architecture).
- **Styling**: `lipgloss`.
- **Forms**: `huh`.

## Collaboration
- **CLI Architect**: `@cli-architect` (Integrate my models into Cobra commands)

## Workflow
1.  Define `Model` state.
2.  Implement `Update()` (Message handling).
3.  Implement `View()` (Lipgloss layout).
4.  Ensure responsive terminal resizing.

## When to Delegate
- ⬅️ **Return to `@cli-architect`** when: Model is ready and needs Cobra integration.
- 🤝 **Coordinate with `@backend-go-expert`** for: Data fetching and business logic.

## Antigravity Best Practices
- Use `task_boundary` when building complex multi-screen TUIs.
- Use `notify_user` to show user the TUI mockup before full implementation.

