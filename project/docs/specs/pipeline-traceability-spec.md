# Spec: Pipeline Traceability Protocol

## Problem Statement

Requirements get lost at each pipeline transition:
- `@bmad-architect` doesn't reference User Stories
- `@tech-spec-writer` doesn't verify coverage
- `@qa-lead` tests against tech-spec, not original requirements

**Result**: E2E reveals missing features that were in original requirements.

---

## Antigravity Context Constraints

| Constraint | Impact |
|------------|--------|
| **One agent, multiple skills** | Agent can't "remember" what previous skill wrote |
| **Small context window** | Can't load all upstream docs at once |
| **Skill separation** | Each skill sees only its input artifact |

**Root Cause**: No explicit traceability mechanism. Each skill trusts previous output completely.

---

## Proposed Solution: Requirements Checklist Protocol

### 1. Mandatory Section in Every Artifact

Every pipeline artifact MUST have **Requirements Checklist** header:

```markdown
## Requirements Checklist

> Source: `project/docs/features/<name>.md` or `user-stories-<name>.md`

| Req ID | Description | Status | Notes |
|--------|-------------|--------|-------|
| REQ-001 | Typing indicator (2-3 sec) | ✅ Covered | Section 1.1 |
| REQ-002 | MAX_INTERVIEW_MESSAGES | ✅ Covered | Section 1.3 |
| REQ-007 | Prompt Guardrails (5 техник) | ⚠️ Partial | Only 2 covered |
| ... | ... | ... | ... |
```

### 2. Status Values

| Status | Meaning |
|--------|---------|
| ✅ Covered | Fully specified in this document |
| ⚠️ Partial | Needs more detail |
| ❌ Deferred | Explicitly out of scope (with reason) |
| ➡️ Upstream | Already handled by previous skill |

### 3. Upstream Reference (Mandatory)

Every artifact MUST start with:

```markdown
## Upstream Documents
- **Input**: `project/docs/...` (the artifact this skill received)
- **Original**: `project/docs/features/...` (source of truth)
```

This allows ANY skill to trace back to original requirements.

---

## Skills to Update

### 1. `@bmad-architect`

**Add to Workflow:**
```markdown
### Phase 0: Load Requirements
1. Read upstream `user-stories-*.md` or `requirements.md`
2. Extract all US-XXX / REQ-XXX items
3. Create Requirements Checklist template in output

### Implementation Phases
Each phase MUST reference:
- `Covers: US-001, US-003`
- Or `Covers: None (infrastructure only)`
```

**Add to Artifact Template:**
- Upstream Documents section
- Requirements Checklist section

---

### 2. `@tech-spec-writer`

**Add to Workflow:**
```markdown
### Phase 0: Traceability Setup
1. Read original `feature-*.md` or `user-stories-*.md`
2. Create Requirements Checklist
3. For EACH AC in User Stories → identify which tech-spec section covers it

### Before Handoff
1. Verify checklist: no ❌ without explicit reason
2. Call out any gaps to user via `notify_user`
```

**Add to Artifact Template:**
- Requirements Checklist (MANDATORY)
- Explicitly map each section to requirements

---

### 3. `@qa-lead`

**Add to Workflow:**
```markdown
### Test Source of Truth
> [!CAUTION]
> Test against **User Stories Acceptance Criteria**, NOT tech-spec!
> Tech-spec may have gaps.

1. Load `user-stories-*.md`
2. Each AC → at least one test case
3. Report shows: `US-001.AC-1: ✅ Passed`
```

---

### 4. `@product-analyst`

**Add to Workflow:**
```markdown
### Output Completeness Check
Before handing off User Stories:
1. Each US has numbered ACs (AC-1, AC-2...)
2. Each AC is testable (has clear pass/fail criteria)
3. Assign REQ-XXX ID to each major requirement for downstream tracing
```

---

## New: Gap Analysis Step (Optional Skill or Checklist)

Consider adding `@gap-analyst` or checklist step at key transitions:

**Option A: Checklist in existing skills**
- `@tech-spec-writer` runs gap analysis before handoff
- Part of existing workflow, no new skill

**Option B: Dedicated `@gap-analyst` skill (overkill?)**
- Only for critical features
- Runs AFTER tech-spec, BEFORE implementation
- Outputs: gap-analysis.md

**Recommendation**: Option A — embed in `@tech-spec-writer`

---

## Implementation Summary

| Skill | Changes |
|-------|---------|
| `@product-analyst` | Add REQ-XXX IDs, numbered ACs |
| `@bmad-architect` | Add Requirements Checklist, phase→US mapping |
| `@tech-spec-writer` | Add Upstream Documents, Requirements Checklist, gap check |
| `@qa-lead` | Test against User Stories, not tech-spec |

## Artifact Template Changes

Add to all templates:
```markdown
## Upstream Documents
- **Input**: [path to input artifact]
- **Original**: [path to feature-fit or user-stories]

## Requirements Checklist
| Req ID | Description | Status | Notes |
|--------|-------------|--------|-------|
```

---

## Decisions (Approved)

| Question | Decision |
|----------|----------|
| **Granularity** | Per User Story (US-XXX), with AC references inside cells |
| **Who creates REQ list** | `@product-analyst` — US numbering is the REQ list |
| **Enforcement** | Checklist in each skill + visual obviousness |

### Enforcement Details

> [!CAUTION]
> **Checklist often not read!** Add Hard Stop section directly in SKILL.md body.

**1. Add to SKILL.md (not just checklist!):**

Each affected skill gets this section in body:
```markdown
## Traceability Protocol (Hard Stop)

> [!CAUTION]
> **Follow `_standards/TRACEABILITY_PROTOCOL.md`.**
> Your output artifact MUST include:
> 1. **Upstream Documents** section — input + original source paths
> 2. **Requirements Checklist** — all US-XXX with status (✅/⚠️/❌)
>
> **BEFORE handoff:**
> - No ❌ without explicit reason
> - Any ⚠️ must be called out to user via `notify_user`
```

**2. Also update `references/checklist.md`** (backup enforcement):
```markdown
## 🚨 Traceability Check (MANDATORY)
- [ ] **Upstream Documents section** — filled with source paths
- [ ] **Requirements Checklist** — all US-XXX statused
- [ ] **No unmarked gaps** — any ⚠️ or ❌ has explanation
```

**3. Visual enforcement**: Empty Requirements Checklist = obvious problem to user
