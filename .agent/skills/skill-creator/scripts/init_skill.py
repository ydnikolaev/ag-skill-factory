#!/usr/bin/env python3
import os
import argparse
import sys
import shutil

def create_skill(name, is_global=False):
    # Determine base path
    if is_global:
        base_path = os.path.expanduser("~/.gemini/antigravity/skills")
    else:
        # Default to .agent/skills in current directory
        if os.path.exists(".agent/skills"):
            base_path = ".agent/skills"
        elif os.path.isdir(".agent"):
             base_path = ".agent/skills"
        else:
            base_path = ".agent/skills"

    skill_path = os.path.join(base_path, name)

    if os.path.exists(skill_path):
        print(f"Error: Skill '{name}' already exists at {skill_path}")
        sys.exit(1)

    # Create directories
    os.makedirs(os.path.join(skill_path, "scripts"), exist_ok=True)
    os.makedirs(os.path.join(skill_path, "resources"), exist_ok=True)
    os.makedirs(os.path.join(skill_path, "examples"), exist_ok=True)
    os.makedirs(os.path.join(skill_path, "references"), exist_ok=True)

    # Resolve template path
    # 1. Try to find the template relative to this script
    script_dir = os.path.dirname(os.path.realpath(__file__))
    template_path = os.path.join(script_dir, "..", "resources", "templates", "SKILL.md")
    checklist_path = os.path.join(script_dir, "..", "resources", "references", "checklist.md")

    # Read template
    if os.path.exists(template_path):
        with open(template_path, 'r') as f:
            content = f.read()
    else:
        # Fallback if template is missing
        print(f"⚠️ Warning: Template not found at {template_path}. Using minimal fallback.")
        content = "---\nname: {{SKILL_NAME}}\ndescription: {{SKILL_DESCRIPTION}}\n---\n\n# {{SKILL_TITLE}}\n\nTODO: Add instructions."

    # Copy checklist (helpful for the agent to self-check)
    if os.path.exists(checklist_path):
        shutil.copy(checklist_path, os.path.join(skill_path, "references", "checklist.md"))

    # Replace placeholders
    content = content.replace("{{SKILL_NAME}}", name)
    content = content.replace("{{SKILL_TITLE}}", name.replace("-", " ").title())
    content = content.replace("{{SKILL_DESCRIPTION}}", f"Description for {name}.")
    content = content.replace("{{SKILL_PURPOSE_SUMMARY}}", f"This skill helps with {name}.")
    content = content.replace("{{SCRIPT_DESC}}", "Description of script.")
    content = content.replace("{{REF_DESC}}", "Description of reference.")

    with open(os.path.join(skill_path, "SKILL.md"), "w") as f:
        f.write(content)

    print(f"✅ Skill '{name}' created successfully at {skill_path}")
    print(f"👉 Next steps:")
    print(f"1. Read {skill_path}/references/checklist.md")
    print(f"2. Edit {skill_path}/SKILL.md")
    print(f"3. Populate resources/ and scripts/ as needed")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Initialize a new Antigravity skill.")
    parser.add_argument("name", help="Name of the skill (kebab-case recommended)")
    parser.add_argument("--global", action="store_true", dest="is_global", help="Install skill globally")

    args = parser.parse_args()
    create_skill(args.name, args.is_global)
