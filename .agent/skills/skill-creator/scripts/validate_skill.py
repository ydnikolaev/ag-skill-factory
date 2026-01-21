#!/usr/bin/env python3
import os
import sys

def validate_skill(path):
    print(f"🔍 Validating skill at {path}...")
    
    if not os.path.isdir(path):
        print(f"❌ Error: Path {path} is not a directory.")
        return False
        
    skill_md = os.path.join(path, "SKILL.md")
    if not os.path.exists(skill_md):
        print(f"❌ Error: SKILL.md not found in {path}")
        return False
        
    # Simple check for frontmatter
    has_frontmatter = False
    with open(skill_md, 'r') as f:
        content = f.read()
        if content.startswith("---\n") and "\n---\n" in content:
            has_frontmatter = True
            
    if not has_frontmatter:
        print(f"❌ Error: SKILL.md missing valid YAML frontmatter (--- ... ---)")
        return False
        
    print(f"✅ Skill structure valid!")
    return True

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: validate_skill.py <path-to-skill>")
        sys.exit(1)
        
    success = validate_skill(sys.argv[1])
    sys.exit(0 if success else 1)
