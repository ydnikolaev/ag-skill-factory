# BMAD Architect Checklist

## 🚨 Document Persistence (MANDATORY)

> [!CAUTION]
> **BEFORE handing off to another skill, you MUST:**

- [ ] **Final document exists in `docs/`** at the path defined in Artifact Ownership
- [ ] **AGENTS.md updated** with status ✅ Done and Last Updated date
- [ ] **Artifact synced** — if you used an artifact, copy final content to `docs/`

**Why?** Artifacts don't persist between sessions. Without `docs/` file, the next skill cannot continue.

## Event Storming
- [ ] All Domain Events identified (orange)
- [ ] Aggregates defined (yellow)
- [ ] Bounded Contexts mapped

## Context Mapping
- [ ] Relationships defined (Partnership, Customer-Supplier, etc.)
- [ ] Context Map diagram created (Mermaid/PlantUML)

## Handover Documents
- [ ] `specs/backend-api.yaml` created for `@backend-go-expert`
- [ ] `specs/ui-mockups.md` created for `@frontend-nuxt-tma`
- [ ] Database schema drafted

## Review
- [ ] Used `mcp_context7` for latest DDD patterns
- [ ] Notified `@product-manager` about feasibility
- [ ] Design approved by user via `notify_user`
