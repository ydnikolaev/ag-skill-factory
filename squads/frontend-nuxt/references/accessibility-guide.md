# Accessibility (a11y) Guide

Making UI accessible to all users.

## 1. Semantic HTML

```html
<!-- Use proper elements -->
<button>Click</button>  ✅
<div onclick>Click</div>  ❌

<nav>...</nav>  ✅
<div class="nav">...</div>  ❌
```

## 2. ARIA Labels

```vue
<!-- When native semantics aren't enough -->
<button aria-label="Close menu">×</button>

<!-- For icons -->
<Icon name="search" aria-hidden="true" />
<span class="sr-only">Search</span>
```

## 3. Keyboard Navigation

- [ ] All interactive elements focusable via Tab
- [ ] Visible focus indicator (don't remove outline!)
- [ ] Escape closes modals
- [ ] Enter/Space activates buttons

```vue
<button @keydown.enter="submit">Submit</button>
```

## 4. Color Contrast

| Size | Min ratio |
|------|-----------|
| Normal text | 4.5:1 |
| Large text (18px+) | 3:1 |

Tools: 
- Chrome DevTools → Accessibility
- https://webaim.org/resources/contrastchecker/

## 5. Images

```vue
<!-- Informative images -->
<img src="chart.png" alt="Sales grew 20% in Q4" />

<!-- Decorative images -->
<img src="decoration.png" alt="" role="presentation" />
```

## 6. Forms

```vue
<label for="email">Email</label>
<input id="email" type="email" aria-describedby="email-hint" />
<span id="email-hint">We'll never share your email</span>
```

## Quick Checklist

- [ ] All images have alt text
- [ ] Color isn't only indicator
- [ ] Keyboard navigable
- [ ] Focus visible
- [ ] Form labels connected
- [ ] Contrast ratios pass
