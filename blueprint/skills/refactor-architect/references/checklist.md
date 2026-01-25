# Refactor Architect Checklist

Quality checklist for refactoring spec creation.

## Pre-Analysis

- [ ] **Scope confirmed** with user (project / module / before-feature)
- [ ] **Depth confirmed** (shallow / deep / custom)
- [ ] **Focus confirmed** (performance / maintainability / coverage / all)

## Context7 Consultation

- [ ] **Library ID resolved** via `mcp_context7_resolve-library-id`
- [ ] **Best practices queried** for relevant frameworks
- [ ] **Gap analysis** documented (current vs ideal)

## Static Analysis

- [ ] **LOC scan** completed — files > 300 LOC flagged
- [ ] **God files** identified — files > 500 LOC listed
- [ ] **Missing tests** detected — uncovered files listed
- [ ] **Complexity** checked — functions > 10 cyclomatic flagged
- [ ] **Layering violations** found — domain → infra leaks documented
- [ ] **Dead code** identified — unused exports listed
- [ ] **Circular deps** checked — architecture smells documented

## Spec Structure

- [ ] `overview.md` created with:
  - [ ] Summary of findings
  - [ ] Priority ranking (Critical / High / Medium / Low)
  - [ ] Risk assessment
  - [ ] Estimated effort per module
- [ ] `modules/*.md` created for each domain:
  - [ ] Clear scope definition
  - [ ] Specific refactoring steps
  - [ ] Designated executor (`@skill-name`)
  - [ ] Success criteria
- [ ] `enforcement/` created with applicable items:
  - [ ] `lint-rules.md` — new golangci-lint rules
  - [ ] `pre-commit-hooks.md` — hook configurations
  - [ ] `ci-additions.md` — pipeline additions
  - [ ] `adrs/*.md` — decision records for major changes

## Enforcement Completeness

> Every problem MUST have a corresponding enforcement

- [ ] God files → `max-lines` lint rule defined
- [ ] Missing tests → CI coverage threshold defined
- [ ] Layering violations → `depguard` rule defined
- [ ] Complexity → `gocyclo` threshold defined
- [ ] Major patterns → ADR written

## 🚨 Document Persistence

> [!CAUTION]
> **BEFORE handoff to executors:**

- [ ] Spec persisted to `project/docs/refactoring/`
- [ ] `overview.md` status changed to `Approved`
- [ ] `project/docs/ARTIFACT_REGISTRY.md` updated with module status
- [ ] User approved via `notify_user`

## Handoff

- [ ] Each module spec assigned to correct executor
- [ ] Handoff instructions clear and actionable
- [ ] Executor skills notified with artifact paths
