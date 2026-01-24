# Skill Interviewer Quality Checklist

## 🚨 Document Persistence (MANDATORY)

> [!CAUTION]
> **Verify the Dual-Write Pattern was followed:**

### Iteration Protocol
- [ ] **Drafts in brain** — all exploration stayed in brain
- [ ] **User approved spec** — got explicit approval before persisting

### Persistence
- [ ] **Spec exists** at `project/docs/specs/skill-<name>-spec.md`
- [ ] **All sections filled** — no empty placeholders

---

## 1. Discovery Phase Complete
- [ ] **Problem identified** — what problem does this skill solve?
- [ ] **User type clear** — who activates this skill?
- [ ] **Trigger phrases defined** — what phrases activate it?

## 2. Boundary Definition
- [ ] **DOES list complete** — clear responsibilities
- [ ] **DOES NOT list complete** — explicit exclusions
- [ ] **Single responsibility** — one primary output
- [ ] **No overlap** — doesn't duplicate existing skill

## 3. Team Fit Analysis
- [ ] **Team reviewed** — read `TEAM.md`
- [ ] **Pipeline understood** — read `PIPELINE.md`
- [ ] **Handoffs defined** — receives from / passes to
- [ ] **Gap confirmed** — this skill fills a real gap

## 4. Technical Shape
- [ ] **Artifact ownership** — what files in `project/docs/`?
- [ ] **Resources needed** — scripts, templates, references?
- [ ] **Mode decided** — interactive or autonomous?

## 5. Naming & Identity
- [ ] **Name follows conventions** — `domain-role` or `action-target`
- [ ] **One-liner clear** — can explain in one sentence
- [ ] **Emoji chosen** — visual identity

## 6. Anti-Patterns Checked
- [ ] **Not too big** — single responsibility
- [ ] **Not too small** — worthy of standalone skill
- [ ] **Not vague** — specific, not "helps with..."
- [ ] **Not overlap** — doesn't duplicate existing

## 7. Spec Quality
- [ ] **All sections present** — Identity, Trigger, Workflow, Boundaries, Team, Artifacts
- [ ] **Open questions resolved** — no blockers for skill-creator
- [ ] **Ready for handoff** — skill-creator can execute from this
