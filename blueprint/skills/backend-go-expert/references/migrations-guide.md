# Database Migrations Guide

Safe migration patterns for production.

## Workflow

1. **Create migration file** with timestamp
2. **Test locally** on dev database
3. **Review** SQL changes
4. **Apply** to staging
5. **Verify** before production

## File Naming

```
migrations/
├── 001_create_users.up.sql
├── 001_create_users.down.sql
├── 002_add_email_index.up.sql
└── 002_add_email_index.down.sql
```

## Safe Patterns

| Operation | Safe Way |
|-----------|----------|
| Add column | Nullable or with default |
| Remove column | First remove code usage, then drop |
| Rename column | Add new → migrate data → drop old |
| Add index | `CREATE INDEX CONCURRENTLY` |

## Rollback Strategy

Every migration needs a `down.sql`:
```sql
-- up.sql
ALTER TABLE users ADD COLUMN phone VARCHAR(20);

-- down.sql
ALTER TABLE users DROP COLUMN phone;
```

## Breaking Changes

For breaking schema changes:
1. Deploy code that works with both schemas
2. Run migration
3. Deploy code that uses new schema only

## Commands

```bash
# Apply migrations
make db-migrate

# Rollback last
make db-rollback

# Status
make db-status
```

## Quick Reference

| Danger | Safe approach |
|--------|---------------|
| Drop column | Remove code first |
| Rename table | Create new → copy → drop old |
| Add NOT NULL | Add nullable → fill → alter |
