# Security Checklist

Apply to all backend code. Check before handoff.

## 1. Input Validation
- [ ] All user input validated and sanitized
- [ ] Use strong typing (no raw `interface{}` from user data)
- [ ] Validate: length, format, allowed characters
- [ ] Reject invalid data early (fail fast)

## 2. Authentication
- [ ] JWT tokens properly validated (signature, expiry, issuer)
- [ ] Telegram initData validated via HMAC-SHA256
- [ ] Session tokens have expiry
- [ ] Logout invalidates tokens

## 3. Authorization
- [ ] Every endpoint checks user permissions
- [ ] No reliance on client-side checks
- [ ] Use middleware for auth guards
- [ ] Principle of least privilege

## 4. Secrets Management
- [ ] **NEVER** log secrets (tokens, passwords, keys)
- [ ] Use environment variables or secret manager
- [ ] No hardcoded credentials in code
- [ ] Rotate secrets periodically

## 5. Database
- [ ] Use parameterized queries (pgx does this by default)
- [ ] Never concatenate SQL strings with user input
- [ ] Limit returned rows (pagination)
- [ ] Mask sensitive data in logs

## 6. API Security
- [ ] HTTPS only (redirect HTTP)
- [ ] CORS configured for allowed origins only
- [ ] Rate limiting enabled
- [ ] Request size limits set

## 7. Error Handling
- [ ] Don't expose stack traces to users
- [ ] Generic error messages externally
- [ ] Detailed errors in logs only
- [ ] Log security events (failed logins, etc.)

## 8. Headers
```go
// Standard security headers
w.Header().Set("X-Content-Type-Options", "nosniff")
w.Header().Set("X-Frame-Options", "DENY")
w.Header().Set("X-XSS-Protection", "1; mode=block")
```

## Quick Reference

| Threat | Mitigation |
|--------|------------|
| SQL Injection | Parameterized queries |
| XSS | Escape output, CSP headers |
| CSRF | Validate origin, use tokens |
| Broken Auth | Validate JWT, check expiry |
| Secrets Leak | Env vars, no logging |
