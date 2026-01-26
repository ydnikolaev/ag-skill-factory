# Debugger Checklist

## Before Starting
- [ ] Read error message and stack trace fully
- [ ] Check `project/CONFIG.yaml` for stack versions
- [ ] Identify which component is affected (backend/frontend/infra)

## 7-Step Workflow
- [ ] **1. Reproduce**: Capture exact conditions to trigger issue
- [ ] **2. Minimize**: Reduce to smallest repro case
- [ ] **3. Hypothesize**: Form 2-5 ranked hypotheses
- [ ] **4. Instrument**: Add logging/assertions
- [ ] **5. Fix**: Apply smallest correct change
- [ ] **6. Prevent**: Add regression test or guard
- [ ] **7. Verify**: Run tests, confirm fix

## Report
- [ ] Documented symptom, root cause, and fix
- [ ] Added regression test
- [ ] Created bug report in `docs/bugs/`

## Handoff Protocol

> [!CAUTION]
> **BEFORE handoff:**
> 1. Save debug report to `docs/bugs/`
> 2. Change file status from `Draft` to `Approved`
> 3. Update `docs/ARTIFACT_REGISTRY.md` status to ✅ Done
> 4. Use `notify_user` for approval
