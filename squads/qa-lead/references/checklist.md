# QA Lead Checklist

## 🚨 Document Persistence (MANDATORY)

> [!CAUTION]
> **BEFORE handing off to another skill, you MUST:**

- [ ] **Final document exists in `docs/`** at the path defined in Artifact Ownership
- [ ] **AGENTS.md updated** with status ✅ Done and Last Updated date
- [ ] **Artifact synced** — if you used an artifact, copy final content to `docs/`

**Why?** Artifacts don't persist between sessions. Without `docs/` file, the next skill cannot continue.

## Test Strategy
- [ ] Test plan documented in `docs/qa/test-plan.md`
- [ ] Test cases written in `docs/qa/test-cases.md`

## Testing
- [ ] API tests passed (Postman/Bruno/curl)
- [ ] E2E tests passed (Playwright/Cypress)
- [ ] Manual testing completed

## TDD Verification (Gatekeeper)
> **Before approving handoff, verify:**

- [ ] New features have unit tests
- [ ] Bug fixes have regression tests  
- [ ] Test coverage not decreased
- [ ] Tests documented: what they prove, how to run

## Bug Reporting
- [ ] All bugs documented with repro steps
- [ ] Severity assigned to each bug
- [ ] Bugs assigned to `@backend-go-expert` or `@frontend-nuxt-tma`

## Release Sign-off
- [ ] All critical bugs fixed
- [ ] Regression tests passed
- [ ] Notified user via `notify_user` for final approval
