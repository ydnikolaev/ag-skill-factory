---
status: Draft
owner: @cli-architect
work_unit: {WORK_UNIT}

downstream:
  - skill: @tui-charm-expert
  - skill: @backend-go-expert

created: {DATE}
updated: {DATE}
---

# Cli Design: {WORK_UNIT}

## Command Structure

```
myapp <command> [subcommand] [flags]
```

---

## Commands

| Command | Description | Example |
|---------|-------------|---------|
| `init` | Initialize project | `myapp init --name foo` |
| `run` | Run application | `myapp run --config ./config.yaml` |

---

## Global Flags

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--config` | `-c` | Config file path | `./config.yaml` |
| `--verbose` | `-v` | Verbose output | `false` |

---

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Invalid arguments |

