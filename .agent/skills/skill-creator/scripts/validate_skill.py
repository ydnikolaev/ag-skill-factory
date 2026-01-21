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
        
    # Read content
    with open(skill_md, 'r') as f:
        content = f.read()
        lines = content.splitlines()

    # 1. Frontmatter Check
    if not (content.startswith("---\n") and "\n---\n" in content):
        print(f"❌ Error: SKILL.md missing valid YAML frontmatter (--- ... ---)")
        return False

    # 2. Line Count Check (Strict for Antigravity)
    line_count = len(lines)
    if line_count > 500:
        print(f"❌ Error: SKILL.md is too long ({line_count} lines). Limit is 500.")
        print("   👉 Tip: Move context to references/ or resources/")
        return False
    else:
        print(f"✅ Length check passed ({line_count}/500 lines)")

    # 3. Best Practices Check
    required_keywords = ["task_boundary", "notify_user"]
    missing_keywords = [k for k in required_keywords if k not in content]
    
    if missing_keywords:
        print(f"⚠️  Warning: SKILL.md seems to miss key Antigravity tools: {missing_keywords}")
        print("   👉 It is HIGHLY recommended to mention these for effective agent control.")
    
    # 4. Resource validation (if mentioned)
    if "scripts/" in content and not os.path.exists(os.path.join(path, "scripts")):
        print(f"⚠️  Warning: 'scripts/' mentioned but directory not found.")

    print(f"✅ Skill structure valid!")
    return True

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: validate_skill.py <path-to-skill>")
        sys.exit(1)
        
    success = validate_skill(sys.argv[1])
    sys.exit(0 if success else 1)
