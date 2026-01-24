# Telegram Mechanic Checklist

## 🚨 Document Persistence (MANDATORY)

> [!CAUTION]
> **BEFORE handing off to another skill, you MUST:**

- [ ] **Final document exists in `docs/`** at the path defined in Artifact Ownership
- [ ] **AGENTS.md updated** with status ✅ Done and Last Updated date
- [ ] **Artifact synced** — if you used an artifact, copy final content to `docs/`

**Why?** Artifacts don't persist between sessions. Without `docs/` file, the next skill cannot continue.

## Bot Setup
- [ ] Bot created in @BotFather
- [ ] Bot Token saved securely (not in code)
- [ ] Menu Button configured (`setChatMenuButton`)

## Webhook
- [ ] Webhook URL configured (HTTPS only)
- [ ] SSL certificate valid
- [ ] `POST /webhook` handler implemented

## Security
- [ ] `initData` validation implemented (HMAC-SHA256)
- [ ] Token used from environment variable
- [ ] Rate limiting considered

## Integration
- [ ] `@backend-go-expert` has Auth middleware
- [ ] `@frontend-nuxt-tma` knows the deep link format
