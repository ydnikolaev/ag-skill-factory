#!/usr/bin/env python3
"""
Generate document templates from SSOT (skill-matrix.yaml, doc-types.yaml).

Usage:
    python3 scripts/generate_templates.py

Output:
    blueprint/_meta/_docs/templates/documents/_{doc_type}.md
"""

import yaml
from pathlib import Path
from datetime import date


# Template content per doc_type (body only, frontmatter is generated)
TEMPLATE_BODIES = {
    "api-contracts": """
## Endpoints

| Method | Path | Description | Auth |
|--------|------|-------------|------|
| GET | /api/v1/... | ... | Bearer |

---

## Request/Response Schemas

### GET /api/v1/example

**Request:** `{}`

**Response:**
```json
{"id": "string", "data": {}}
```

---

## Error Codes

| Code | Message | When |
|------|---------|------|
| 400 | Bad Request | Invalid input |
| 401 | Unauthorized | Missing/invalid token |
| 404 | Not Found | Resource doesn't exist |
""",
    "cli-design": """
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
""",
    "debug-report": """
## Issue Summary

| Field | Value |
|-------|-------|
| **Symptom** | ... |
| **Severity** | Critical / High / Medium / Low |
| **Environment** | Production / Staging / Dev |

---

## Reproduction Steps

1. ...
2. ...

---

## Root Cause

<!-- What was the actual cause? -->

---

## Fix Applied

```diff
- old code
+ new code
```

---

## Prevention

- [ ] Test added
- [ ] Monitoring added
- [ ] Documentation updated
""",
    "deployment-guide": """
## Prerequisites

- [ ] Docker installed
- [ ] Access to registry
- [ ] Environment variables configured

---

## Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `DATABASE_URL` | Postgres connection string | ✅ |
| `SECRET_KEY` | Application secret | ✅ |

---

## Build

```bash
docker build -t app:latest .
```

---

## Deploy

```bash
docker compose -f docker-compose.prod.yml up -d
```

---

## Health Checks

| Endpoint | Expected |
|----------|----------|
| `/health` | 200 OK |
""",
    "feature-brief": """
## Request Summary

| Field | Value |
|-------|-------|
| **Requester** | ... |
| **Priority** | High / Medium / Low |

---

## Feature Description

<!-- What does the user want? -->

---

## Gap Analysis

| Capability | Current | Required | Gap |
|------------|---------|----------|-----|
| ... | ❌ | ✅ | New implementation needed |

---

## Estimated Effort

| Phase | Effort |
|-------|--------|
| Design | X days |
| Implementation | X days |
| **Total** | **X days** |
""",
    "ui-implementation": """
## Pages

| Route | Component | Status |
|-------|-----------|--------|
| `/` | `pages/index.vue` | ⬜ |
| `/dashboard` | `pages/dashboard.vue` | ⬜ |

---

## Components

| Component | Location | Props |
|-----------|----------|-------|
| `Button` | `components/ui/Button.vue` | `variant`, `size` |

---

## API Integration

| Endpoint | Composable | Status |
|----------|------------|--------|
| `GET /api/items` | `useItems()` | ⬜ |

---

## Testing

- [ ] Component tests
- [ ] E2E tests
""",
    "server-config": """
## Server Definition

```yaml
servers:
  my-server:
    type: stdio
    command: go run ./cmd/mcp
```

---

## Tools

| Tool | Description | Parameters |
|------|-------------|------------|
| `db_query` | Execute SQL query | `query: string` |

---

## Resources

| URI | Description |
|-----|-------------|
| `db://schema` | Database schema |
""",
    "roadmap": """
## Vision

<!-- High-level vision for this initiative -->

---

## Milestones

| Milestone | Target Date | Status |
|-----------|-------------|--------|
| MVP | ... | ⬜ |
| Beta | ... | ⬜ |
| GA | ... | ⬜ |

---

## Phases

### Phase 1: Foundation
- [ ] Core infrastructure
- [ ] Basic features

### Phase 2: Enhancement
- [ ] Advanced features
- [ ] Integrations
""",
    "requirements": """
## Functional Requirements

| ID | Requirement | Priority | Status |
|----|-------------|----------|--------|
| FR-01 | ... | Must | ⬜ |
| FR-02 | ... | Should | ⬜ |

---

## Non-Functional Requirements

| ID | Category | Requirement |
|----|----------|-------------|
| NFR-01 | Performance | Response time < 200ms |
| NFR-02 | Security | All endpoints require auth |

---

## Constraints

- ...

---

## Out of Scope

- ...
""",
    "test-cases": """
## Test Summary

| Type | Count | Pass | Fail |
|------|-------|------|------|
| Unit | ... | ... | ... |
| Integration | ... | ... | ... |
| E2E | ... | ... | ... |

---

## Test Cases

### TC-01: ...

| Field | Value |
|-------|-------|
| **Preconditions** | ... |
| **Steps** | 1. ... 2. ... |
| **Expected** | ... |
| **Status** | ⬜ |
""",
    "refactoring-overview": """
## Current State

<!-- Describe the current problematic state -->

---

## Goals

- [ ] Improve maintainability
- [ ] Reduce complexity
- [ ] Increase test coverage

---

## Scope

| File/Module | Change Type | Priority |
|-------------|-------------|----------|
| ... | Restructure | High |

---

## Metrics

| Metric | Before | Target | After |
|--------|--------|--------|-------|
| LOC | ... | ... | ... |
| Coverage | ...% | ...% | ...% |
""",
    "webhook-config": """
## Webhook Setup

| Field | Value |
|-------|-------|
| **URL** | `https://api.example.com/telegram/webhook` |
| **Secret Token** | `TELEGRAM_WEBHOOK_SECRET` |

---

## Allowed Updates

- `message`
- `callback_query`

---

## Commands

| Command | Description | Handler |
|---------|-------------|---------|
| `/start` | Start bot | `StartHandler` |
| `/help` | Show help | `HelpHandler` |
""",
    "tma-config": """
## SDK Setup

```typescript
import { init, retrieveLaunchParams } from '@telegram-apps/sdk'

export default defineNuxtPlugin(() => {
  init()
  const lp = retrieveLaunchParams()
  return { provide: { telegram: lp } }
})
```

---

## Main Button

```typescript
import { mainButton } from '@telegram-apps/sdk'

mainButton.setParams({ text: 'Submit', isVisible: true })
mainButton.on('click', () => { /* handle */ })
```

---

## Security

- [ ] initData validation on backend
- [ ] User ID extraction
""",
    "tui-design": """
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
""",
    "theming": """
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
""",
    "tokens": """
{
  "colors": {
    "primary": {"500": "#3b82f6", "600": "#2563eb"},
    "gray": {"50": "#f9fafb", "900": "#111827"}
  },
  "typography": {
    "fontFamily": {"sans": ["Inter", "system-ui"]},
    "fontSize": {"base": "1rem", "lg": "1.125rem"}
  },
  "spacing": {"1": "0.25rem", "2": "0.5rem", "4": "1rem"}
}
""",
    "design-system": """
## Tokens

See [tokens.json](./tokens.json) for raw values.

---

## Colors

| Name | Usage | Token |
|------|-------|-------|
| Primary | CTAs, links | `colors.primary.500` |
| Background | Page background | `colors.gray.50` |

---

## Typography

| Style | Size | Weight |
|-------|------|--------|
| H1 | 2.25rem | 700 |
| Body | 1rem | 400 |

---

## Components

### Button

| Variant | Background | Text |
|---------|------------|------|
| Primary | primary.500 | white |
| Secondary | gray.100 | gray.900 |
""",
}


def load_yaml(path: Path) -> dict:
    """Load YAML file."""
    with open(path) as f:
        return yaml.safe_load(f) or {}


def build_upstream_downstream(handoffs: list, doc_type: str) -> tuple[list, list]:
    """Build upstream and downstream lists for a doc_type based on handoffs."""
    upstream = []
    downstream = []
    
    for h in handoffs:
        if h.get("doc_type") == doc_type:
            # This doc_type is handed off from h["from"] to h["to"]
            # So h["from"] is the creator (upstream is whoever sends TO the creator)
            # And h["to"] is the consumer (downstream)
            downstream.append({
                "skill": h["to"],
            })
    
    # Find what doc_types feed into the creator of this doc_type
    # For now, use the doc-types.yaml to find the creator
    return upstream, downstream


def generate_frontmatter(doc_type: str, creator: str, lifecycle: str, handoffs: list) -> str:
    """Generate YAML frontmatter for a document template."""
    upstream_docs = []
    downstream_docs = []
    
    # Handoffs where this doc_type is passed FROM creator TO someone
    for h in handoffs:
        if h.get("doc_type") == doc_type:
            downstream_docs.append({"skill": h["to"]})
    
    # Handoffs where creator receives something (those are upstream)
    for h in handoffs:
        if h.get("to") == creator and h.get("doc_type") != doc_type:
            upstream_docs.append({"doc_type": h["doc_type"], "owner": h["from"]})
    
    lines = [
        "---",
        "status: Draft",
        f"owner: @{creator}",
        f"lifecycle: {lifecycle}",
    ]
    
    # For per-feature docs, work_unit is a placeholder
    # For living docs, work_unit is omitted (they use fixed names)
    if lifecycle == "per-feature":
        lines.append("work_unit: {WORK_UNIT}")
    
    lines.append("")
    
    if upstream_docs:
        lines.append("upstream:")
        seen = set()
        for u in upstream_docs:
            key = (u["doc_type"], u["owner"])
            if key not in seen:
                seen.add(key)
                lines.append(f"  - doc_type: {u['doc_type']}")
                lines.append(f"    owner: @{u['owner']}")
    
    if downstream_docs:
        lines.append("downstream:")
        seen = set()
        for d in downstream_docs:
            if d["skill"] not in seen:
                seen.add(d["skill"])
                lines.append(f"  - skill: @{d['skill']}")
    
    lines.extend(["", "created: {DATE}", "updated: {DATE}", "---"])
    
    return "\n".join(lines)


def generate_template(doc_type: str, creator: str, lifecycle: str, handoffs: list) -> str:
    """Generate complete template content."""
    frontmatter = generate_frontmatter(doc_type, creator, lifecycle, handoffs)
    
    body = TEMPLATE_BODIES.get(doc_type, """
## Overview

<!-- Brief description -->

---

## Details

<!-- Main content -->

---

## Checklist

- [ ] Item 1
- [ ] Item 2
""")
    
    title = doc_type.replace("-", " ").title()
    
    # Living docs have fixed title, per-feature docs have work_unit placeholder
    if lifecycle == "living":
        return f"""{frontmatter}

# {title}
{body}
"""
    else:
        return f"""{frontmatter}

# {title}: {{WORK_UNIT}}
{body}
"""


def main():
    root = Path(__file__).parent.parent
    matrix_file = root / "blueprint" / "_meta" / "_skills" / "skill-matrix.yaml"
    doc_types_file = root / "blueprint" / "_meta" / "_docs" / "doc-types.yaml"
    output_dir = root / "blueprint" / "_meta" / "_docs" / "templates" / "documents"
    
    matrix = load_yaml(matrix_file)
    doc_types = load_yaml(doc_types_file)
    
    handoffs = matrix.get("handoffs", [])
    types = doc_types.get("types", {})
    
    output_dir.mkdir(parents=True, exist_ok=True)
    
    print("📄 Generating document templates...")
    
    for doc_type, info in types.items():
        creator = info.get("creator", "unknown")
        lifecycle = info.get("lifecycle", "per-feature")
        
        content = generate_template(doc_type, creator, lifecycle, handoffs)
        
        # Determine extension
        ext = ".json" if doc_type == "tokens" else ".md"
        output_file = output_dir / f"_{doc_type}{ext}"
        
        output_file.write_text(content)
        print(f"  ✅ _{doc_type}{ext}")
    
    print(f"\n📁 Generated {len(types)} templates in {output_dir.relative_to(root)}")


if __name__ == "__main__":
    main()
