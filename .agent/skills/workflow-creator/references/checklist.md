# Workflow Creator Quality Checklist

## 🚨 Document Persistence (MANDATORY)

> [!CAUTION]
> **Verify the Dual-Write Pattern was followed:**

### Iteration Protocol
- [ ] **Proposal in brain** — draft stayed in brain for user review
- [ ] **User approved** — got "Looks good" before persisting

### Persistence
- [ ] **Workflow file exists** at `.agent/workflows/<name>.md`
- [ ] **Frontmatter correct** with `description:` field

---

## 1. Interview Complete
- [ ] **Trigger defined** — slash command name confirmed
- [ ] **Goal clear** — end artifact or action identified
- [ ] **Mode decided** — turbo-all or step-by-step?
- [ ] **Skills identified** — which skills are involved?

## 2. Overlap Check
- [ ] **Existing workflows reviewed** — `ls .agent/workflows/`
- [ ] **No duplicate** — this doesn't replicate existing workflow
- [ ] **Complementary** — integrates with existing workflows if related

## 3. Workflow Structure
- [ ] **Frontmatter present** — description in YAML header
- [ ] **Steps numbered** — clear 1, 2, 3... progression
- [ ] **Bash blocks used** — for auto-runnable commands
- [ ] **Turbo annotation** — `// turbo` or `// turbo-all` if appropriate

## 4. Skill Integration
- [ ] **Skills matched** — steps reference appropriate skills
- [ ] **Team.md consulted** — used correct skill names
- [ ] **Pipeline aware** — follows logical flow

## 5. Quality
- [ ] **Atomic steps** — one action per step
- [ ] **Report included** — final step summarizes results
- [ ] **Tested** — workflow was run at least once
