# Installer Merge Mode

## Problem
Current `factory install` deletes entire `.agent/` and `project/` folders before copying.
This destroys existing project configs and custom content.

## Required Behavior

### Directory Mapping
| dist/ | Target | Notes |
|-------|--------|-------|
| `_agent/skills/` | `.agent/skills/` | Merge by skill name |
| `_agent/rules/` | `.agent/rules/` | Merge by file |
| `_agent/workflows/` | `.agent/workflows/` | Merge by file |
| `docs/templates/` | `project/docs/templates/` | Merge by file |
| `config/*` | Root or config location | Copy if not exists |

### Merge Logic
1. **Skills**: Replace entire skill folder by name (if backend-go-expert exists, replace it)
2. **Rules/Workflows**: Replace individual .md files by name
3. **Templates**: Replace individual files by name
4. **Config files**: Skip if exists (or use --force flag)

### What NOT to do
- ❌ `RemoveAll(.agent/)` 
- ❌ `RemoveAll(project/)`
- ❌ Delete user's custom workflows
- ❌ Overwrite project-specific configs

## Implementation
1. Update `installer.Install()` to use merge logic
2. Add `copyDirMerge()` that doesn't delete existing
3. Add `--force` flag for config overwrites
4. Update tests

## Priority
Medium — affects production projects like sky-cli
