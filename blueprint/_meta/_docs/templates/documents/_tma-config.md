---
status: Draft
owner: @tma-expert
work_unit: {WORK_UNIT}

upstream:
  - doc_type: webhook-config
    owner: @telegram-mechanic
downstream:
  - skill: @qa-lead

created: {DATE}
updated: {DATE}
---

# Tma Config: {WORK_UNIT}

## SDK Setup

```typescript
import { init, retrieveLaunchParams } from '@telegram-apps/sdk'

export default defineNuxtPlugin(() => {
  init()
  const lp = retrieveLaunchParams()
  return { provide: { telegram: lp } }
})
```

---

## Main Button

```typescript
import { mainButton } from '@telegram-apps/sdk'

mainButton.setParams({ text: 'Submit', isVisible: true })
mainButton.on('click', () => { /* handle */ })
```

---

## Security

- [ ] initData validation on backend
- [ ] User ID extraction

