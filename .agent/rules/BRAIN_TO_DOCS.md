---
trigger: always_on
---

# BRAIN_TO_DOCS

> Draft in brain/, persist to project/docs/ only after user approval.

| Phase | Location | Lifetime |
|-------|----------|----------|
| Draft | `brain/` (`~/.gemini/antigravity/brain/<conversation-id>/`) | Per-conversation |
| Persist | `project/docs/` (git-tracked) | Permanent |

**Phase 1: Draft in Brain**  
Create artifacts in brain/. Iterate with user via `notify_user`.  
All drafts stay here: briefs, user stories, tech specs, implementation plans, etc.

**Phase 2: Persist on Approval**  
ONLY after user says "Looks good" → copy to persistent location.  
Git-tracked folders are the only memory that survives between sessions.