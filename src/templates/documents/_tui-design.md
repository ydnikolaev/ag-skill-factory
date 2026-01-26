---
status: Draft
owner: @tui-charm-expert
lifecycle: per-feature
work_unit: {WORK_UNIT}

upstream:
  - doc_type: cli-design
    owner: @cli-architect
downstream:
  - skill: @qa-lead

created: {DATE}
updated: {DATE}
---

# Tui Design: {WORK_UNIT}

## Model

```go
type model struct {
    items    []string
    cursor   int
    selected map[int]struct{}
}
```

---

## Key Bindings

| Key | Action |
|-----|--------|
| `↑/k` | Move up |
| `↓/j` | Move down |
| `Enter` | Select |
| `q` | Quit |

---

## Styles

```go
var titleStyle = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("205"))
```

