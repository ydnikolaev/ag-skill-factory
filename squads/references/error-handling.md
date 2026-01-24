# Error Handling Patterns

Shared guide for backend and frontend.

---

## Go Error Patterns

### 1. Custom Error Types
```go
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
```

### 2. Error Wrapping
```go
if err != nil {
    return fmt.Errorf("failed to create user: %w", err)
}
```

### 3. Error Response Format
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid email format",
    "details": { "field": "email" }
  }
}
```

### 4. Don't Panic
- Use `panic` only for unrecoverable state
- Return errors, don't panic
- Recover in middleware for safety

---

## Vue/Nuxt Error Patterns

### 1. Error Boundaries
```vue
<template>
  <NuxtErrorBoundary>
    <MyComponent />
    <template #error="{ error }">
      <p>Error: {{ error.message }}</p>
    </template>
  </NuxtErrorBoundary>
</template>
```

### 2. Loading/Error States
```vue
<script setup>
const { data, error, status } = useFetch('/api/users')
</script>

<template>
  <div v-if="status === 'pending'">Loading...</div>
  <div v-else-if="error">Error: {{ error.message }}</div>
  <div v-else>{{ data }}</div>
</template>
```

### 3. Toast Notifications
```ts
const { $toast } = useNuxtApp()
try {
  await saveUser()
  $toast.success('Saved!')
} catch (e) {
  $toast.error(e.message)
}
```

### 4. Global Error Handler
```ts
// plugins/error-handler.ts
export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.config.errorHandler = (error, instance, info) => {
    console.error('Global error:', error)
    // Report to monitoring
  }
})
```

---

## Quick Reference

| Situation | Go | Vue |
|-----------|-----|-----|
| Validation fail | Return ValidationError | Show inline error |
| API error | Wrap with context | Toast + log |
| Unexpected | Log + generic response | Error boundary |
| Not found | Return 404 | Redirect or message |
