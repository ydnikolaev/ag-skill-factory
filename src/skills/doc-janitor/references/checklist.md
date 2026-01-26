# Doc Janitor Checklist

Use this checklist before completing a cleanup operation.

---

## Pre-Run Checks

- [ ] Read `../standards/DOCUMENT_STRUCTURE_PROTOCOL.md`
- [ ] Identify target project directory
- [ ] Confirm `project/docs/` exists

## Dry-Run Phase

- [ ] Scanned entire `project/docs/` structure
- [ ] Detected legacy format signals
- [ ] Listed all orphan files
- [ ] Identified archive candidates
- [ ] Generated change report
- [ ] Presented report via `notify_user`

## Report Quality

- [ ] Statistics section complete (files scanned, issues, actions)
- [ ] All structure fixes listed with current → new paths
- [ ] All archive actions grouped by Work Unit
- [ ] ARTIFACT_REGISTRY.md changes listed
- [ ] Clear "apply" instruction at bottom

## Apply Phase

- [ ] User explicitly said "apply"
- [ ] Created missing folders
- [ ] Moved files to correct lifecycle folders
- [ ] Added missing YAML frontmatter
- [ ] Archived completed Work Units
- [ ] Updated ARTIFACT_REGISTRY.md format

## ARTIFACT_REGISTRY.md Quality

- [ ] Work Units structure used
- [ ] Active units at top
- [ ] Review units in middle
- [ ] Closed units in `<details>` blocks
- [ ] Quick Links table present

## Post-Apply

- [ ] All changes committed with `chore(docs): doc-janitor cleanup`
- [ ] No orphan files in `project/docs/` root
- [ ] No empty folders (cleaned up)
- [ ] Summary report sent via `notify_user`

---

## 🚨 Document Persistence (MANDATORY)

> [!CAUTION]
> **Before completing ANY cleanup:**
> 1. ✅ Changes saved to `project/docs/`
> 2. ✅ ARTIFACT_REGISTRY.md updated
> 3. ✅ Git commit created

**Failure = incomplete cleanup!**
