# Frontend Security Checklist

Apply to all frontend code. Check before handoff.

## 1. XSS Prevention
- [ ] Never use `v-html` with user input
- [ ] Sanitize any HTML from API responses
- [ ] Use template literals, not string concatenation

## 2. CSRF Protection
- [ ] Validate `Origin` header on API
- [ ] Use SameSite cookies
- [ ] Include CSRF token in forms if needed

## 3. Content Security Policy
- [ ] Configure CSP headers via Nuxt Security module
- [ ] Restrict script sources
- [ ] Report violations

## 4. Sensitive Data
- [ ] Never store tokens in localStorage (use httpOnly cookies)
- [ ] Clear auth state on logout
- [ ] Don't log sensitive data to console

## 5. API Calls
- [ ] Always use HTTPS
- [ ] Validate responses before rendering
- [ ] Handle 401/403 properly (redirect to login)

## 6. Dependencies
- [ ] Keep packages updated
- [ ] Run `npm audit` regularly
- [ ] Don't trust unvetted packages

## Quick Reference

| Threat | Mitigation |
|--------|------------|
| XSS | No v-html with user data |
| CSRF | SameSite cookies, validate origin |
| Token theft | httpOnly cookies, not localStorage |
| Data exposure | Don't log sensitive info |
