# Example: Feature Brief for "User Notifications"

---
status: Approved
author: feature-fit
date: 2026-01-23
---

# Feature: User Push Notifications

## Context

**Project Stack** (from CONFIG.yaml):
- Backend: Go 1.25, PostgreSQL 16
- Frontend: Nuxt 4, TailwindCSS v4
- Integrations: Telegram Bot API, Firebase (not yet integrated)

**Where it fits**:
- Bounded Context: User Management
- Existing Components: Uses User table, adds NotificationPreferences

## Requirements

**User Goal**: Users want to receive timely updates about their subscriptions via push notifications.

**Key Features**:
1. Push notification opt-in/opt-out settings
2. Real-time delivery via Firebase Cloud Messaging
3. Notification history in user profile

**Scope**:
- ✅ MVP: Basic push for subscription renewals
- ❌ V2: Rich notifications with actions, notification center

## Gap Analysis

### Backend
- [ ] Add `notification_preferences` table (user_id, push_enabled, topics)
- [ ] Create `/api/notifications/preferences` endpoint
- [ ] Integrate Firebase Admin SDK

### Frontend
- [ ] Add Notification Settings page in user profile
- [ ] Implement service worker for push reception
- [ ] Add notification permission request flow

### MCP / Integrations
- [ ] Add `firebase-mcp` server for FCM operations

## Impact / Risks

| Area | Impact | Risk Level |
|------|--------|------------|
| Database | New table, FK to users | 🟢 Low |
| API | New endpoints, no breaking changes | 🟢 Low |
| Security | Firebase credentials management | 🟡 Medium |

## TDD Strategy

- **Unit Tests**: Mock Firebase SDK, test preference CRUD
- **Integration Tests**: /api/notifications/* endpoints
- **E2E Tests**: User enables notifications → receives test push

## Next Steps

Delegate to `@product-analyst` for User Stories and Specs.
