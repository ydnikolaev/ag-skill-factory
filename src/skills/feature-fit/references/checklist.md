# Feature-Fit Quality Checklist

## 🚨 Document Persistence (MANDATORY)

> [!CAUTION]
> **Verify the Dual-Write Pattern was followed:**

### Iteration Protocol
- [ ] **Draft stayed in brain** — no unapproved content written to `docs/`
- [ ] **User approved** — got "Looks good" via `notify_user` before persisting

### Persistence
- [ ] **Feature Brief exists in `docs/features/`** at path defined in Artifact Ownership
- [ ] **File status changed** from `Draft` to `Approved` in header/frontmatter
- [ ] **ARTIFACT_REGISTRY.md updated** with status ✅ Done and Last Updated date

**Why?** Brain artifacts are wiped every session. Without `docs/` file, next skill cannot continue.

---

## 1. Context Loading (Phase 1)
- [ ] **Read `CONFIG.yaml`** — understood current stack and modules
- [ ] **Read `mcp.yaml`** — identified available MCP tools
- [ ] **Read Architecture docs** — know the Bounded Contexts
- [ ] **Read Roadmap** — checked if feature is already planned

## 2. Feature Interview (Phase 2)
- [ ] **Goal defined** — clear what feature does
- [ ] **Data requirements** — new tables or uses existing?
- [ ] **UI requirements** — new screens or components?
- [ ] **Integrations** — any new external APIs needed?

## 3. Gap Analysis (Phase 3)
- [ ] **Backend gaps identified** — new domain logic, constraints
- [ ] **Frontend gaps identified** — new routes, stores, components
- [ ] **MCP gaps identified** — new MCP server needed?

## 4. Impact Assessment (Phase 4)
- [ ] **DB impact documented** — specific tables affected
- [ ] **API impact documented** — breaking contract changes?
- [ ] **Security impact checked** — new permissions needed?

## 5. Feature Brief Quality
- [ ] **All sections filled** — Context, Requirements, Gap Analysis, Risks
- [ ] **No implementation details** — leave design to Architect
- [ ] **Clear scope** — MVP vs V2 split defined
- [ ] **TDD strategy noted** — how will this be tested?

## 6. Handoff Readiness
- [ ] **User approved brief** — confirmed via `notify_user`
- [ ] **Brief saved to `docs/features/<name>.md`**
- [ ] **Ready for `@product-analyst`** — can create specs from this
