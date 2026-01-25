# DevOps SRE Checklist

## 🚨 Document Persistence (MANDATORY)

> [!CAUTION]
> **BEFORE handing off to another skill, you MUST:**

- [ ] **Final document exists in `docs/`** at the path defined in Artifact Ownership
- [ ] **ARTIFACT_REGISTRY.md updated** with status ✅ Done and Last Updated date
- [ ] **Artifact synced** — if you used an artifact, copy final content to `docs/`

**Why?** Artifacts don't persist between sessions. Without `docs/` file, the next skill cannot continue.

## Docker
- [ ] `Dockerfile.backend` uses multi-stage build
- [ ] `Dockerfile.frontend` optimized for Nuxt
- [ ] Images are minimal (distroless/scratch)

## Compose
- [ ] `docker-compose.yml` works for local dev
- [ ] Database service included
- [ ] Volumes configured correctly

## CI/CD
- [ ] GitHub Actions workflow created
- [ ] Tests run in pipeline
- [ ] Build artifacts stored

## Deployment
- [ ] Production server accessible via SSH
- [ ] SSL configured (Certbot)
- [ ] Nginx reverse proxy configured
- [ ] User notified before production deploy
