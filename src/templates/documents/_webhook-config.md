---
status: Draft
owner: @telegram-mechanic
lifecycle: per-feature
work_unit: {WORK_UNIT}

downstream:
  - skill: @backend-go-expert
  - skill: @tma-expert

created: {DATE}
updated: {DATE}
---

# Webhook Config: {WORK_UNIT}

## Webhook Setup

| Field | Value |
|-------|-------|
| **URL** | `https://api.example.com/telegram/webhook` |
| **Secret Token** | `TELEGRAM_WEBHOOK_SECRET` |

---

## Allowed Updates

- `message`
- `callback_query`

---

## Commands

| Command | Description | Handler |
|---------|-------------|---------|
| `/start` | Start bot | `StartHandler` |
| `/help` | Show help | `HelpHandler` |

