# CLI Architect Checklist

## 🚨 Document Persistence (MANDATORY)

> [!CAUTION]
> **BEFORE handing off to another skill, you MUST:**

- [ ] **Final document exists in `docs/`** at the path defined in Artifact Ownership
- [ ] **AGENTS.md updated** with status ✅ Done and Last Updated date
- [ ] **Artifact synced** — if you used an artifact, copy final content to `docs/`

**Why?** Artifacts don't persist between sessions. Without `docs/` file, the next skill cannot continue.

## Structure
- [ ] Root command defined
- [ ] Subcommands organized logically
- [ ] `--help` works for all commands

## Flags
- [ ] Persistent flags on root (verbose, config)
- [ ] Local flags on subcommands
- [ ] Viper config integration

## Signals
- [ ] SIGINT handled (Ctrl+C)
- [ ] SIGTERM handled (graceful shutdown)
- [ ] Context cancellation propagated

## Quality
- [ ] POSIX compliant exit codes
- [ ] No panics in production paths
- [ ] `@tui-charm-expert` integrated for UI
