#!/usr/bin/env python3
"""
Initialize a new Antigravity skill in the squads/ directory.

Usage:
    python3 init_skill.py <skill-name>

Skills are created in squads/ and then linked to the global brain via:
    make install-squads
"""
import os
import argparse
import sys
import shutil


def find_repo_root():
    """Find the ag-skill-factory repository root."""
    current = os.getcwd()
    while current != "/":
        if os.path.exists(os.path.join(current, "squads")) and \
           os.path.exists(os.path.join(current, ".agent", "skills", "skill-creator")):
            return current
        current = os.path.dirname(current)
    return None


def create_skill(name: str):
    """Create a new skill in the squads/ directory."""
    
    # Find the repository root
    repo_root = find_repo_root()
    if not repo_root:
        print("❌ Error: Could not find ag-skill-factory repository root.")
        print("   Make sure you're running this from within the repository.")
        sys.exit(1)
    
    # Skills are always created in squads/
    squads_dir = os.path.join(repo_root, "squads")
    skill_path = os.path.join(squads_dir, name)
    
    if os.path.exists(skill_path):
        print(f"❌ Error: Skill '{name}' already exists at {skill_path}")
        sys.exit(1)
    
    # Create directories
    os.makedirs(os.path.join(skill_path, "scripts"), exist_ok=True)
    os.makedirs(os.path.join(skill_path, "resources"), exist_ok=True)
    os.makedirs(os.path.join(skill_path, "examples"), exist_ok=True)
    os.makedirs(os.path.join(skill_path, "references"), exist_ok=True)
    
    # Resolve template path (relative to this script)
    script_dir = os.path.dirname(os.path.realpath(__file__))
    template_path = os.path.join(script_dir, "..", "resources", "templates", "SKILL.md")
    checklist_path = os.path.join(script_dir, "..", "resources", "references", "checklist.md")
    
    # Read template
    if os.path.exists(template_path):
        with open(template_path, 'r') as f:
            content = f.read()
    else:
        print(f"⚠️  Warning: Template not found at {template_path}. Using minimal fallback.")
        content = """---
name: {{SKILL_NAME}}
description: {{SKILL_DESCRIPTION}}
---

# {{SKILL_TITLE}}

TODO: Add instructions.
"""
    
    # Copy checklist (helpful for self-checking)
    if os.path.exists(checklist_path):
        shutil.copy(checklist_path, os.path.join(skill_path, "references", "checklist.md"))
    
    # Replace placeholders
    replacements = {
        "{{SKILL_NAME}}": name,
        "{{SKILL_TITLE}}": name.replace("-", " ").title(),
        "{{SKILL_DESCRIPTION}}": f"Description for {name}.",
        "{{SKILL_PURPOSE_SUMMARY}}": f"This skill helps with {name}.",
        "{{SCRIPT_DESC}}": "Description of script.",
        "{{REF_DESC}}": "Description of reference.",
    }
    
    for placeholder, value in replacements.items():
        content = content.replace(placeholder, value)
    
    with open(os.path.join(skill_path, "SKILL.md"), "w") as f:
        f.write(content)
    
    print(f"✅ Skill '{name}' created at {skill_path}")
    print()
    print("👉 Next steps:")
    print(f"   1. Edit {skill_path}/SKILL.md")
    print(f"   2. Add resources/scripts as needed")
    print(f"   3. Run: make install-squads")
    print()
    print("This will link the skill to ~/.gemini/antigravity/skills/")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description="Initialize a new Antigravity skill in squads/.",
        epilog="After creation, run 'make install-squads' to link to global brain."
    )
    parser.add_argument("name", help="Name of the skill (kebab-case recommended)")
    
    args = parser.parse_args()
    create_skill(args.name)
