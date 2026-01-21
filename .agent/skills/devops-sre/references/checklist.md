# DevOps SRE Checklist

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
