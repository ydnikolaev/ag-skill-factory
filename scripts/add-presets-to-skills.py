#!/usr/bin/env python3
"""Add presets field to all skill frontmatters"""
import re
from pathlib import Path

SKILLS_DIR = Path("blueprint/skills")

# Mapping: skill -> leaf preset(s)
SKILL_PRESETS = {
    # Core skills
    "idea-interview": ["core"],
    "product-analyst": ["core"],
    "bmad-architect": ["core"],
    "tech-spec-writer": ["core"],
    "qa-lead": ["core"],
    "doc-janitor": ["core"],
    "refactor-architect": ["core"],
    
    # Backend skills
    "backend-go-expert": ["backend"],
    "devops-sre": ["backend"],
    "debugger": ["backend", "minimal"],
    "project-bro": ["backend", "minimal"],
    
    # Frontend skills
    "frontend-nuxt": ["frontend", "tma"],
    "ux-designer": ["frontend", "tma"],
    "ui-implementor": ["frontend", "tma"],
    
    # TMA skills
    "tma-expert": ["tma"],
    "telegram-mechanic": ["tma"],
    
    # CLI skills
    "cli-architect": ["cli"],
    "tui-charm-expert": ["cli"],
    
    # Special skills
    "feature-fit": ["core"],
    "mcp-expert": ["core"],
}

def add_presets_to_skill(skill_dir: Path, presets: list):
    skill_file = skill_dir / "SKILL.md"
    if not skill_file.exists():
        return False
    
    content = skill_file.read_text()
    
    # Check if already has presets
    if re.search(r'^presets:', content, re.MULTILINE):
        print(f"  Skip {skill_dir.name} (already has presets)")
        return False
    
    # Find category: line and add presets after it
    presets_yaml = "presets:\n" + "\n".join(f"  - {p}" for p in presets) + "\n"
    
    # Try to insert after category:
    new_content = re.sub(
        r'(^category:.*\n)',
        r'\1\n' + presets_yaml,
        content,
        count=1,
        flags=re.MULTILINE
    )
    
    if new_content == content:
        # Try after phase:
        new_content = re.sub(
            r'(^phase:.*\n)',
            r'\1\n' + presets_yaml,
            content,
            count=1,
            flags=re.MULTILINE
        )
    
    if new_content != content:
        skill_file.write_text(new_content)
        print(f"  ✅ Added {presets} to {skill_dir.name}")
        return True
    else:
        print(f"  ⚠️  Could not add to {skill_dir.name}")
        return False

def main():
    print("Adding presets to skills...")
    added = 0
    
    for skill_dir in sorted(SKILLS_DIR.iterdir()):
        if not skill_dir.is_dir():
            continue
        
        skill_name = skill_dir.name
        presets = SKILL_PRESETS.get(skill_name, ["core"])  # default to core
        
        if add_presets_to_skill(skill_dir, presets):
            added += 1
    
    print(f"\n✅ Added presets to {added} skills")

if __name__ == "__main__":
    main()
