## Document Workflow

> **Protocol**: See @DOCUMENT_STRUCTURE_PROTOCOL

<!-- TODO: доработать, сделать заглушку с автогенерацией правильных путей из пайплайнов для каждого скилла -->

<!-- ===== idea-interview =====
| Operation | Document | Location | Trigger |
|-----------|----------|----------|---------|
| 🔵 Creates | discovery-brief.md | `active/discovery/` | Interview complete |
| 📖 Reads | CONFIG.yaml | `project/` | On activation (if exists) |
| 📝 Updates | ARTIFACT_REGISTRY.md | `project/docs/` | On create, on complete |
| 🟡 To Review | discovery-brief.md | `review/discovery/` | User approves draft |
| ✅ Archive | — | `closed/<work-unit>/` | @doc-janitor on final approval |
===== /idea-interview ===== -->

<!-- ===== product-analyst =====
| Operation | Document | Location | Trigger |
|-----------|----------|----------|---------|
| 🔵 Creates | roadmap.md | `active/product/` | Product definition complete |
| 🔵 Creates | user-stories.md | `active/product/` | User stories written |
| 🔵 Creates | requirements.md | `active/specs/` | Requirements finalized |
| 📖 Reads | discovery-brief.md | `active/discovery/` | On activation |
| 📝 Updates | ARTIFACT_REGISTRY.md | `project/docs/` | On create, on complete |
| 🟡 To Review | roadmap.md, user-stories.md, requirements.md | `review/product/`, `review/specs/` | User approves drafts |
| ✅ Archive | — | `closed/<work-unit>/` | @doc-janitor on final approval |
===== /product-analyst ===== -->

<!-- ===== backend-go-expert =====
| Operation | Document | Location | Trigger |
|-----------|----------|----------|---------|
| 🔵 Creates | service-implementation.md | `active/backend/` | Implementation complete |
| 📖 Reads | `<feature>-tech-spec.md` | `active/specs/` | On activation |
| 📖 Reads | api-contracts.yaml | `active/architecture/` | On activation |
| 📝 Updates | ARTIFACT_REGISTRY.md | `project/docs/` | On create, on complete |
| 🟡 To Review | service-implementation.md | `review/backend/` | Ready for QA |
| ✅ Archive | — | `closed/<work-unit>/` | @doc-janitor on final approval |
===== /backend-go-expert ===== -->
