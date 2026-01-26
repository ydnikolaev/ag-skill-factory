---
status: Draft
owner: @ui-implementor
work_unit: {WORK_UNIT}

upstream:
  - doc_type: tokens
    owner: @ux-designer
  - doc_type: design-system
    owner: @ux-designer
downstream:
  - skill: @frontend-nuxt

created: {DATE}
updated: {DATE}
---

# Theming: {WORK_UNIT}

## Color Palette

| Token | Light | Dark |
|-------|-------|------|
| `--color-primary` | #3B82F6 | #60A5FA |
| `--color-bg` | #FFFFFF | #1F2937 |
| `--color-text` | #111827 | #F9FAFB |

---

## Typography

| Token | Value |
|-------|-------|
| `--font-sans` | Inter, system-ui |
| `--text-base` | 16px |

---

## Dark Mode

```css
@media (prefers-color-scheme: dark) {
  :root { /* dark tokens */ }
}
```

