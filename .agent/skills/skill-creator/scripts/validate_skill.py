#!/usr/bin/env python3
"""
Validate a skill against quality standards.

Usage:
    python3 validate_skill.py <path-to-skill>
    python3 validate_skill.py blueprint/skills/mcp-expert
"""
import os
import sys
import re


def check_links(skill_path: str, content: str) -> list:
    """Check that all internal links in SKILL.md are valid.
    
    Returns list of (link_type, path) tuples for broken links.
    """
    broken = []
    
    # Pattern: See examples/foo.py, See references/bar.md
    see_pattern = re.compile(r'See\s+([a-zA-Z_]+/[^\s\)]+)')
    for match in see_pattern.finditer(content):
        ref_path = match.group(1)
        full_path = os.path.join(skill_path, ref_path)
        if not os.path.exists(full_path):
            broken.append(("file", ref_path))
    
    # Pattern: @skill-name (not @user or email)
    skill_ref_pattern = re.compile(r'(?<![a-zA-Z0-9])@([a-z][a-z0-9-]+)(?![a-zA-Z0-9@.])')
    known_skills = set()
    
    # Get blueprint skills path
    factory_root = os.path.dirname(os.path.dirname(os.path.dirname(os.path.dirname(skill_path))))
    blueprint_skills = os.path.join(factory_root, "blueprint", "skills")
    
    if os.path.isdir(blueprint_skills):
        known_skills = set(os.listdir(blueprint_skills))
    
    # Add factory skills
    factory_skills = {"skill-creator", "skill-factory-expert", "skill-interviewer", 
                      "skill-updater", "workflow-creator"}
    known_skills.update(factory_skills)
    
    for match in skill_ref_pattern.finditer(content):
        ref_skill = match.group(1)
        if ref_skill not in known_skills:
            broken.append(("skill", ref_skill))
    
    return broken


def validate_skill(path: str) -> bool:
    """Validate a skill directory against quality standards."""
    print(f"🔍 Validating skill at {path}...")
    errors = []
    warnings = []
    
    # Check directory exists
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
    
    # ========================================
    # 1. Frontmatter Check
    # ========================================
    if not (content.startswith("---\n") and "\n---\n" in content):
        errors.append("SKILL.md missing valid YAML frontmatter (--- ... ---)")
    else:
        # Check required frontmatter fields
        frontmatter_end = content.find("\n---\n", 4)
        frontmatter = content[4:frontmatter_end]
        
        if "name:" not in frontmatter:
            errors.append("Frontmatter missing 'name:' field")
        if "description:" not in frontmatter:
            errors.append("Frontmatter missing 'description:' field")
    
    # ========================================
    # 2. Line Count Check (max 500)
    # ========================================
    line_count = len(lines)
    if line_count > 500:
        errors.append(f"SKILL.md is too long ({line_count} lines). Limit is 500.")
    else:
        print(f"✅ Length: {line_count}/500 lines")
    
    # ========================================
    # 3. Best Practices Check
    # ========================================
    recommended_keywords = ["task_boundary", "notify_user"]
    missing_keywords = [k for k in recommended_keywords if k not in content]
    if missing_keywords:
        warnings.append(f"Missing recommended Antigravity tools: {missing_keywords}")
    
    # ========================================
    # 4. Team Collaboration Check
    # ========================================
    if "## Team Collaboration" not in content:
        warnings.append("Missing '## Team Collaboration' section")
    
    if "## When to Delegate" not in content:
        warnings.append("Missing '## When to Delegate' section")
    
    if "## Document Lifecycle" not in content:
        warnings.append("Missing '## Document Lifecycle' section")
        
    if "## Handoff Protocol" not in content:
        warnings.append("Missing '## Handoff Protocol' section (New Requirement)")
    
    if "## Iteration Protocol" not in content and "Iteration Protocol" not in content:
        warnings.append("Missing '## Iteration Protocol' section explaining brain→docs flow")
    
    # ========================================
    # 5. Link Checker (NEW)
    # ========================================
    broken_links = check_links(path, content)
    for link_type, link_path in broken_links:
        if link_type == "file":
            errors.append(f"Broken link: '{link_path}' not found")
        elif link_type == "skill":
            warnings.append(f"Unknown skill reference: @{link_path}")
    
    if not broken_links:
        print("✅ Links: all valid")
    
    # ========================================
    # 6. Examples Check (no large code blocks in SKILL.md)
    # ========================================
    # Count lines in code blocks
    in_code_block = False
    code_block_lines = 0
    max_code_block = 0
    current_block = 0
    
    for line in lines:
        if line.startswith("```"):
            if in_code_block:
                max_code_block = max(max_code_block, current_block)
                current_block = 0
            in_code_block = not in_code_block
        elif in_code_block:
            current_block += 1
            code_block_lines += 1
    
    if max_code_block > 15:
        warnings.append(f"Large code block found ({max_code_block} lines). Consider moving to examples/")
    
    # ========================================
    # 6. Language Check (English only)
    # ========================================
    # Check for Cyrillic characters (Russian, Ukrainian, etc.)
    cyrillic_pattern = re.compile(r'[\u0400-\u04FF]')
    cyrillic_lines = []
    for i, line in enumerate(lines, 1):
        if cyrillic_pattern.search(line):
            # Skip if it's just a comment or quote
            if not line.strip().startswith('#') and not line.strip().startswith('>'):
                cyrillic_lines.append(i)
    
    if cyrillic_lines:
        if len(cyrillic_lines) > 5:
            errors.append(f"SKILL.md contains Cyrillic text (lines: {cyrillic_lines[:5]}... and {len(cyrillic_lines) - 5} more). Skills must be in English.")
        else:
            errors.append(f"SKILL.md contains Cyrillic text (lines: {cyrillic_lines}). Skills must be in English.")
    else:
        print("✅ Language: English")
    
    # ========================================
    # 7. Directory Structure Check
    # ========================================
    expected_dirs = ["examples", "references", "resources", "scripts"]
    for dir_name in expected_dirs:
        dir_path = os.path.join(path, dir_name)
        if not os.path.exists(dir_path):
            # Only warn if mentioned in content
            if f"{dir_name}/" in content:
                warnings.append(f"'{dir_name}/' mentioned but directory not found")
    
    # Check examples/ has content if skill has code examples
    examples_dir = os.path.join(path, "examples")
    if os.path.isdir(examples_dir):
        examples_files = [f for f in os.listdir(examples_dir) if not f.startswith('.')]
        if examples_files:
            print(f"✅ Examples: {len(examples_files)} files")
        else:
            warnings.append("examples/ directory is empty")
    
    # Check references/checklist.md exists and is customized
    checklist_path = os.path.join(path, "references", "checklist.md")
    if os.path.exists(checklist_path):
        with open(checklist_path, 'r') as f:
            checklist_content = f.read()
        if "Use this checklist to verify your skill" in checklist_content:
            warnings.append("references/checklist.md appears to be the generic template. Customize it!")
        else:
            print("✅ Checklist: customized")
    else:
        warnings.append("references/checklist.md not found")
    
    # ========================================
    # Print Results
    # ========================================
    if errors:
        print("\n❌ ERRORS (must fix):")
        for e in errors:
            print(f"   • {e}")
    
    if warnings:
        print("\n⚠️  WARNINGS (should fix):")
        for w in warnings:
            print(f"   • {w}")
    
    if not errors and not warnings:
        print("\n🎉 Skill is perfect!")
    elif not errors:
        print("\n✅ Skill structure valid (with warnings)")
    
    return len(errors) == 0


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: validate_skill.py <path-to-skill>")
        print("Example: validate_skill.py blueprint/skills/mcp-expert")
        sys.exit(1)
    
    success = validate_skill(sys.argv[1])
    sys.exit(0 if success else 1)
