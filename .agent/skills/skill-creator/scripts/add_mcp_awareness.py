#!/usr/bin/env python3
"""
Update all skills with expanded CONFIG.yaml + MCP awareness block.
Replaces the old "First Step" block with the new expanded version.
"""

import os
import re
from pathlib import Path

SQUADS_DIR = Path("/Users/yuranikolaev/Developer/antigravity/ag-skill-factory/squads")

OLD_BLOCK_PATTERN = r'> \[!IMPORTANT\]\n> ## First Step: Read Project Config\n> Before making technical decisions, \*\*always check\*\*:\n> ```\n> project/CONFIG\.yaml\n> ```\n> This file defines: stack versions, modules, architecture style, features\.\n> \*\*Never assume defaults — verify against CONFIG\.yaml first\.\*\*'

NEW_BLOCK = '''> [!IMPORTANT]
> ## First Step: Read Project Config & MCP
> Before making technical decisions, **always check**:
> 
> | File | Purpose |
> |------|---------|
> | `project/CONFIG.yaml` | Stack versions, modules, architecture |
> | `mcp.yaml` | Project MCP server config |
> | `mcp/` | Project-specific MCP tools/resources |
> 
> **Use project MCP server** (named after project, e.g. `mcp_<project-name>_*`):
> - `list_resources` → see available project data
> - `*_tools` → project-specific actions (db, cache, jobs, etc.)
> 
> **Use `mcp_context7`** for library docs:
> - Check `mcp.yaml → context7.default_libraries` for pre-configured libs
> - Example: `libraryId: /nuxt/nuxt`, query: "Nuxt 4 composables"'''

def update_skill(skill_path: Path) -> bool:
    """Update a single skill's SKILL.md with expanded MCP awareness block."""
    skill_md = skill_path / "SKILL.md"
    
    if not skill_md.exists():
        return False
    
    content = skill_md.read_text()
    
    # Skip if already has new block
    if "First Step: Read Project Config & MCP" in content:
        print(f"  ⏭️  {skill_path.name}: already has MCP awareness")
        return False
    
    # Check if has old block
    if "First Step: Read Project Config" not in content:
        print(f"  ⚠️  {skill_path.name}: no First Step block found")
        return False
    
    # Replace old block with new
    new_content = re.sub(OLD_BLOCK_PATTERN, NEW_BLOCK, content)
    
    if new_content == content:
        print(f"  ⚠️  {skill_path.name}: pattern not matched exactly")
        return False
    
    skill_md.write_text(new_content)
    print(f"  ✅ {skill_path.name}: updated with MCP awareness")
    return True

def main():
    print("🔄 Updating skills with MCP awareness...\n")
    
    updated = 0
    skipped = 0
    failed = 0
    
    for item in sorted(SQUADS_DIR.iterdir()):
        if item.is_dir() and not item.name.startswith('.'):
            if item.name in ['references']:
                continue
            if (item / "SKILL.md").exists():
                result = update_skill(item)
                if result:
                    updated += 1
                elif "already has" in str(result) if result else False:
                    skipped += 1
                else:
                    failed += 1
    
    print(f"\n📊 Results: {updated} updated, {skipped} skipped, {failed} need manual fix")

if __name__ == "__main__":
    main()
