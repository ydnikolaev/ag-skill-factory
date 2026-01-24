# Skill Quality Checklist

## 🚨 Document Persistence (MANDATORY)

## 0. TDD Compliance (HARD STOP)
> [!CAUTION]
> **Does this skill enforce Red-Green-Refactor?**
- [ ] **Hard Block**: Does it say "Refuse to write code without a failing test"? (Ref: `squads/TDD_PROTOCOL.md`)
- [ ] **Debugging**: Does it require a reproduction case before fixing bugs?
- [ ] **Planning**: Does it require a Verification Strategy in plans?

> [!CAUTION]
> **Verify the Dual-Write Pattern was followed:**

### Iteration Protocol
- [ ] **Drafts stayed in brain** — no unapproved content written to `project/docs/`
- [ ] **User approved** — got "Looks good" via `notify_user` before persisting

### Persistence
- [ ] **Final document exists in `project/docs/`** at path defined in Artifact Ownership
- [ ] **File status changed** from `Draft` to `Approved` in header/frontmatter
- [ ] **ARTIFACT_REGISTRY.md updated** with status ✅ Done and Last Updated date

**Why?** Brain artifacts are wiped every session. Without `project/docs/` file, next skill cannot continue.

Use this checklist to verify your skill before finishing.

## 1. Structure & Paths
- [ ] **Folder Structure**: Does the skill follow `squads/<skill-name>/`?
- [ ] **Resources**: Are scripts in `scripts/`, templates in `resources/`?
- [ ] **Examples**: Are code examples in `examples/`, NOT embedded in SKILL.md?
- [ ] **References**: Are documentation/cheatsheets in `references/`?
- [ ] **Absolute Paths**: Does the skill use absolute paths where needed?

## 2. SKILL.md Content
- [ ] **Frontmatter**: Is the YAML valid? (`name`, `description` only).
- [ ] **Description**: Is it in third person? Does it explain *when* to trigger?
- [ ] **Conciseness**: Is the file under 500 lines?
- [ ] **No Embedded Code**: Are large code examples moved to `examples/`?
- [ ] **Decision Tree**: Does it guide the agent on *how* to choose actions?
- [ ] **References to Examples**: Does SKILL.md reference `examples/` files instead of embedding code?

## 3. Best Practices (Antigravity Specific)
- [ ] **Task Boundaries**: For complex workflows, does it instruct to use `task_boundary`?
- [ ] **User Interaction**: Does it instruct when to use `notify_user`?
- [ ] **Tools**: Does it reference standard MCP tools correctly?

## 4. Scripts
- [ ] **Help**: Do scripts support a `--help` flag?
- [ ] **Consistency**: Do scripts output predictable results?
- [ ] **Permissions**: Are scripts executable (`chmod +x`)?

## 5. IDE Considerations
- [ ] **Context**: Does the skill assume the agent knows file structure? (Often not, so `ls` instructions help).

## 6. Content Organization Rules

**SKILL.md should contain:**
- Decisions, workflows, and logic (the "brain")
- Brief inline examples (max 10 lines)
- References to detailed examples: `See examples/python-server.py`

**examples/ should contain:**
- Full working code examples
- Configuration samples
- Complete templates

**references/ should contain:**
- Cheatsheets
- External documentation links
- Troubleshooting guides
