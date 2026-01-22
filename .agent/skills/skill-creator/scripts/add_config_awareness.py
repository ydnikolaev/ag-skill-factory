#!/usr/bin/env python3
"""
Mass update all skills with CONFIG.yaml awareness block.
Adds the "First Step" protocol after frontmatter.
"""

import os
import re
from pathlib import Path

SQUADS_DIR = Path("/Users/yuranikolaev/Developer/antigravity/ag-skill-factory/squads")

FIRST_STEP_BLOCK = '''
> [!IMPORTANT]
> ## First Step: Read Project Config
> Before making technical decisions, **always check**:
> ```
> project/CONFIG.yaml
> ```
> This file defines: stack versions, modules, architecture style, features.
> **Never assume defaults — verify against CONFIG.yaml first.**
'''

def update_skill(skill_path: Path) -> bool:
    """Update a single skill's SKILL.md with First Step block."""
    skill_md = skill_path / "SKILL.md"
    
    if not skill_md.exists():
        return False
    
    content = skill_md.read_text()
    
    # Skip if already has First Step
    if "First Step: Read Project Config" in content:
        print(f"  ⏭️  {skill_path.name}: already has First Step")
        return False
    
    # Find the end of frontmatter (second ---)
    # Pattern: starts with ---, then content, then ---
    frontmatter_pattern = r'^(---\n.*?\n---\n)'
    match = re.match(frontmatter_pattern, content, re.DOTALL)
    
    if not match:
        print(f"  ❌ {skill_path.name}: no frontmatter found")
        return False
    
    frontmatter = match.group(1)
    rest = content[len(frontmatter):]
    
    # Find the first heading (# Title)
    heading_match = re.match(r'^(\n*# .+\n+.*?\n)', rest, re.DOTALL)
    
    if heading_match:
        # Insert after first paragraph following heading
        # Find first blank line after heading
        first_section = heading_match.group(1)
        rest_after_heading = rest[len(first_section):]
        
        # Insert First Step block
        new_content = frontmatter + first_section + FIRST_STEP_BLOCK + "\n" + rest_after_heading
    else:
        # Just insert after frontmatter
        new_content = frontmatter + FIRST_STEP_BLOCK + "\n" + rest
    
    skill_md.write_text(new_content)
    print(f"  ✅ {skill_path.name}: updated")
    return True

def main():
    print("🔄 Mass updating skills with CONFIG.yaml awareness...\n")
    
    updated = 0
    skipped = 0
    
    for item in sorted(SQUADS_DIR.iterdir()):
        if item.is_dir() and not item.name.startswith('.'):
            # Skip non-skill directories
            if item.name in ['references']:
                continue
            if (item / "SKILL.md").exists():
                if update_skill(item):
                    updated += 1
                else:
                    skipped += 1
    
    print(f"\n📊 Results: {updated} updated, {skipped} skipped")

if __name__ == "__main__":
    main()
