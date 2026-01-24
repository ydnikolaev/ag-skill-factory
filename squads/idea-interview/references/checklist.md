# Discovery Readiness Checklist

## 🚨 Document Persistence (MANDATORY)

> [!CAUTION]
> **BEFORE handing off to another skill, you MUST:**

- [ ] **Final document exists in `docs/`** at the path defined in Artifact Ownership
- [ ] **AGENTS.md updated** with status ✅ Done and Last Updated date
- [ ] **Artifact synced** — if you used an artifact, copy final content to `docs/`

**Why?** Artifacts don't persist between sessions. Without `docs/` file, the next skill cannot continue.

Checklist for readiness to hand off to `@product-manager`.

## Information Completeness

### Business Context
- [ ] Problem is clearly formulated
- [ ] Understanding of competitors or their absence
- [ ] Monetization model defined
- [ ] KPIs/success metrics named

### Target Users
- [ ] Target audience described
- [ ] User roles defined (if multiple)
- [ ] Main use case understood

### Platform
- [ ] Platform chosen (TMA/Web/Mobile/CLI)
- [ ] Reason for choice is clear
- [ ] Integrations listed

### Core Functionality
- [ ] 3-5 key features defined
- [ ] Main data entities named
- [ ] Notifications/events described (if needed)

### MVP Scope
- [ ] MVP boundaries clearly defined
- [ ] "NOT in MVP" list exists
- [ ] Timeline/deadline known
- [ ] Constraints (budget, team) stated

## Quality Checks

### Consistency
- [ ] No contradictions between sections
- [ ] Priorities align with business goal
- [ ] Scope is realistic for timeline

### Completeness
- [ ] All required fields in brief are filled
- [ ] Open Questions documented
- [ ] No critical gaps

## Self-Evolve

> After each interview, check:

- [ ] Did questions arise that aren't in `question-bank.md`?
  - If yes → add them
- [ ] Were there poorly worded questions?
  - If yes → rephrase in question-bank
- [ ] Did `discovery-brief-template.md` work well?
  - If not → propose improvements

## Handoff Readiness

- [ ] Discovery Brief created in `docs/discovery-brief.md`
- [ ] User confirmed correctness (via `notify_user`)
- [ ] Ready to hand off to `@product-manager`
