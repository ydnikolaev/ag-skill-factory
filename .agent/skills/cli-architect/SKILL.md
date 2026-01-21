---
name: cli-architect
description: Expert in Go CLI Architecture (Cobra, Viper, POSIX).
---

# CLI Architect

This skill designs the Command Line Interface using Cobra, Viper, and POSIX standards.

## Tech Stack
- **Framework**: `spf13/cobra`.
- **Config**: `spf13/viper`.
- **Standards**: POSIX compliance, 12-factor CLI.

## Collaboration
- **TUI**: `@tui-charm-expert` (Handle the fancy UI)
- **Backend**: `@backend-go-expert` (Reuse business logic)

## Workflow
1.  Define Command Structure (`root -> sub -> leaf`).
2.  Define Flags (Persistent vs Local).
3.  Handle OS Signals (Graceful Shutdown).

## Antigravity Best Practices
- Use `task_boundary` when adding new command groups.
- Use `notify_user` if breaking changes to CLI interface are needed.

