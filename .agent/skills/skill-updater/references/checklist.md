# Skill Updater Quality Checklist

## 🚨 Document Persistence (MANDATORY)

> [!CAUTION]
> **Verify the Dual-Write Pattern was followed:**

### Iteration Protocol
- [ ] **Preview created in brain** — change plan as artifact
- [ ] **User approved preview** — got "Looks good" via `notify_user`

### Execution
- [ ] **Feature branch created** — not updating on main
- [ ] **All affected skills updated** — no partial updates
- [ ] **Validation passed** — `make validate-all` succeeds

---

## 1. Context Loading
- [ ] **Read TEAM.md** — know current skill roster
- [ ] **Read _standards/** — understand current protocols
- [ ] **Identified all affected skills** — no missed skills

## 2. Preview Quality
- [ ] **Listed all affected skills** — with specific change description
- [ ] **Sample diff shown** — user can see exactly what changes
- [ ] **Approval obtained** — via `notify_user`

## 3. Execution Quality
- [ ] **Batch processing** — efficient, not one-by-one manually
- [ ] **Consistent changes** — same pattern applied everywhere
- [ ] **No regressions** — validate-all passes

## 4. Git Discipline
- [ ] **Feature branch** — `refactor/skill-update-<desc>`
- [ ] **Conventional commit** — `refactor(skills): <description>`
- [ ] **No orphan changes** — all changes in one commit

## 5. Cleanup
- [ ] **No leftover debug code** — no commented-out sections
- [ ] **No TODO in changes** — complete implementation
