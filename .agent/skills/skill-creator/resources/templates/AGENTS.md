# Artifact Registry

> Single source of truth for project documentation. Each skill updates this on artifact create/update.

| Artifact | Owner | Path | Status | Last Updated |
|----------|-------|------|--------|--------------|
| Discovery Brief | @idea-interview | `project/docs/discovery/discovery-brief.md` | ⏳ Pending | - |
| Roadmap | @product-analyst | `project/docs/product/roadmap.md` | ⏳ Pending | - |
| User Stories | @product-analyst | `project/docs/product/user-stories.md` | ⏳ Pending | - |
| Requirements | @product-analyst | `project/docs/specs/requirements.md` | ⏳ Pending | - |
| Context Map | @bmad-architect | `project/docs/architecture/context-map.md` | ⏳ Pending | - |
| API Contracts | @bmad-architect | `project/docs/architecture/api-contracts.yaml` | ⏳ Pending | - |
| Test Cases | @qa-lead | `project/docs/qa/test-cases.md` | ⏳ Pending | - |
| Design Tokens | @ux-designer | `project/docs/design/tokens.json` | ⏳ Pending | - |

## Status Legend

- ⏳ Pending — not started
- 📝 Draft — work in progress
- 🔄 Review — awaiting user approval
- ✅ Done — approved, ready for downstream

## Versioning

Git-based. Use `git log project/docs/<path>` to see history.

## Rules

- **Create**: Owner creates artifact → updates this table
- **Update**: Only owner modifies → updates "Last Updated"
- **Handoff**: Downstream reads from path → if missing, returns to owner
