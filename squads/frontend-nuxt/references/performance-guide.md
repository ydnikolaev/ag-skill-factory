# Frontend Performance Guide

Optimization patterns for Nuxt apps.

## 1. Lazy Loading

```vue
<!-- Lazy load components -->
<LazyUserProfile v-if="showProfile" />

<!-- Lazy load images -->
<NuxtImg loading="lazy" src="/image.jpg" />
```

## 2. Code Splitting

Nuxt auto-splits by route. Additional:
```ts
// Dynamic import for heavy libs
const Chart = defineAsyncComponent(() => import('chart.js'))
```

## 3. Core Web Vitals

| Metric | Target | How to improve |
|--------|--------|----------------|
| LCP | < 2.5s | Optimize images, SSR |
| FID | < 100ms | Defer non-critical JS |
| CLS | < 0.1 | Reserve space for images |

## 4. Image Optimization

```vue
<NuxtImg
  src="/photo.jpg"
  width="400"
  height="300"
  format="webp"
  quality="80"
/>
```

## 5. Data Fetching

```ts
// Prefetch on hover
<NuxtLink prefetch to="/users" />

// Cache API responses
const { data } = useFetch('/api/users', {
  key: 'users',
  transform: (data) => data.users
})
```

## 6. Bundle Analysis

```bash
npx nuxi analyze
```

## Quick Reference

| Problem | Solution |
|---------|----------|
| Slow load | SSR, lazy load |
| Large bundle | Code split, tree shake |
| Slow images | NuxtImg, webp |
| API latency | Prefetch, cache |
