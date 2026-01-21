# Telegram Mechanic Checklist

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
