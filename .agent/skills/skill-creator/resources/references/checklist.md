# Skill Quality Checklist

Use this checklist to verify your skill before finishing.

## 1. Structure & Paths
- [ ] **Folder Structure**: Does the skill follow `.agent/skills/<skill-name>/`?
- [ ] **Resources**: Are scripts in `scripts/`, templates in `resources/`?
- [ ] **Absolute Paths**: Does the skill use absolute paths or relative paths correctly (relative to the workspace root)? *Note: Scripts usually run from workspace root.*

## 2. SKILL.md Content
- [ ] **Frontmatter**: Is the YAML valid? (`name`, `description` only).
- [ ] **Description**: Is it in the third person? Does it clearly explain *when* to trigger the skill?
- [ ] **Conciseness**: Is the file under 500 lines?
- [ ] **No Fluff**: Did you remove "Usage" sections that just repeat standard agent behavior?
- [ ] **Decision Tree**: Does it guide the agent on *how* to choose the right tools/scripts?

## 3. Best Practices (Antigravity Specific)
- [ ] **Task Boundaries**: If the skill involves a complex workflow, does it instruct the agent to use `task_boundary`?
- [ ] **User Interaction**: Does it instruct when to use `notify_user`?
- [ ] **Tools**: Does it reference standard MCP tools (e.g., `run_command`, `write_to_file`) correctly?

## 4. Scripts
- [ ] **Help**: Do scripts support a `--help` flag?
- [ ] **Consistency**: Do scripts output predictable results?
- [ ] **Permissions**: Are scripts executable (`chmod +x`)?

## 5. IDE Considerations
- [ ] **Context**: Does the skill assume the agent knows the file structure? (It often does not, so `ls -R` or `find` instructions are helpful).
