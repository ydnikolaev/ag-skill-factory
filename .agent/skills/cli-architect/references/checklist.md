# CLI Architect Checklist

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
