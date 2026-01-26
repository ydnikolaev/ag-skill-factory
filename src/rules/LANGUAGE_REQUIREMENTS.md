---
trigger: always_on
description: Match the user language for chat and drafts; English only for persistent docs.
---

# LANGUAGE_REQUIREMENTS

> Ephemeral artifacts in user's language; persistent docs in English only.

| Context | Language |
|---------|----------|
| **Chat with user** | Match user's language from conversation |
| **Brain artifacts** (drafts, iteration) | Match user's language from conversation |
| **Persist docs** (any location, e.g. `project/docs/`) | English only |

> Detect user language from their messages. If unclear, ask.
