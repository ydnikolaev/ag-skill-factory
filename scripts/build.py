#!/usr/bin/env python3
"""
Build script: src/ → dist/

Processes {{include: path}} directives and assembles final files.

Usage:
    python3 scripts/build.py
    
    # Or via make
    make build
"""

import re
import shutil
from pathlib import Path
from datetime import datetime

# Include patterns:
# New: {{include: partials/pre-handoff-validation.md}}
# Legacy: <!-- INCLUDE: _meta/_skills/sections/pre-handoff-validation.md -->
INCLUDE_PATTERN = re.compile(r'\{\{include:\s*([^\}]+)\}\}')
LEGACY_INCLUDE_PATTERN = re.compile(r'<!--\s*INCLUDE:\s*([^>]+)\s*-->')


def remap_legacy_path(legacy_path: str) -> str:
    """Remap legacy include paths to new structure."""
    legacy_path = legacy_path.strip()
    # _meta/_skills/sections/X.md -> partials/X.md
    if "_meta/_skills/sections/" in legacy_path:
        filename = legacy_path.split("/")[-1]
        return f"partials/{filename}"
    return legacy_path


def process_includes(content: str, src_dir: Path, file_path: Path) -> str:
    """Replace include directives with file contents."""
    
    def replace_new_include(match):
        include_path = match.group(1).strip()
        full_path = src_dir / include_path
        
        if not full_path.exists():
            print(f"  ⚠️  Include not found: {include_path} (in {file_path})")
            return f"<!-- ERROR: Include not found: {include_path} -->"
        
        include_content = full_path.read_text()
        # Recursively process includes in included files
        include_content = process_includes(include_content, src_dir, full_path)
        return include_content
    
    def replace_legacy_include(match):
        legacy_path = match.group(1).strip()
        include_path = remap_legacy_path(legacy_path)
        full_path = src_dir / include_path
        
        if not full_path.exists():
            print(f"  ⚠️  Legacy include not found: {legacy_path} -> {include_path} (in {file_path})")
            return f"<!-- ERROR: Include not found: {legacy_path} -->"
        
        include_content = full_path.read_text()
        include_content = process_includes(include_content, src_dir, full_path)
        return include_content
    
    # Process both patterns
    content = INCLUDE_PATTERN.sub(replace_new_include, content)
    content = LEGACY_INCLUDE_PATTERN.sub(replace_legacy_include, content)
    
    return content


def build_skills(src_dir: Path, dist_dir: Path):
    """Build skills from src/ to dist/.agent/skills/."""
    src_skills = src_dir / "skills"
    dist_skills = dist_dir / ".agent" / "skills"
    
    if not src_skills.exists():
        print("  ⚠️  No src/skills/ directory")
        return 0
    
    dist_skills.mkdir(parents=True, exist_ok=True)
    count = 0
    
    for skill_dir in src_skills.iterdir():
        if not skill_dir.is_dir():
            continue
        
        skill_name = skill_dir.name
        skill_md = skill_dir / "SKILL.md"
        
        if not skill_md.exists():
            continue
        
        # Process includes
        content = skill_md.read_text()
        processed = process_includes(content, src_dir, skill_md)
        
        # Write to dist
        dist_skill_dir = dist_skills / skill_name
        dist_skill_dir.mkdir(parents=True, exist_ok=True)
        (dist_skill_dir / "SKILL.md").write_text(processed)
        
        # Copy examples if exist
        examples_dir = skill_dir / "examples"
        if examples_dir.exists():
            shutil.copytree(examples_dir, dist_skill_dir / "examples", dirs_exist_ok=True)
        
        count += 1
        print(f"  ✅ {skill_name}")
    
    return count


def build_rules(src_dir: Path, dist_dir: Path):
    """Build rules from src/ to dist/.agent/rules/."""
    src_rules = src_dir / "rules"
    dist_rules = dist_dir / ".agent" / "rules"
    
    if not src_rules.exists():
        print("  ⚠️  No src/rules/ directory")
        return 0
    
    dist_rules.mkdir(parents=True, exist_ok=True)
    count = 0
    
    for rule_file in src_rules.glob("*.md"):
        content = rule_file.read_text()
        processed = process_includes(content, src_dir, rule_file)
        (dist_rules / rule_file.name).write_text(processed)
        count += 1
        print(f"  ✅ {rule_file.name}")
    
    return count


def build_workflows(src_dir: Path, dist_dir: Path):
    """Build workflows from src/ to dist/.agent/workflows/."""
    src_workflows = src_dir / "workflows"
    dist_workflows = dist_dir / ".agent" / "workflows"
    
    if not src_workflows.exists():
        print("  ⚠️  No src/workflows/ directory")
        return 0
    
    dist_workflows.mkdir(parents=True, exist_ok=True)
    count = 0
    
    for wf_file in src_workflows.glob("*.md"):
        content = wf_file.read_text()
        processed = process_includes(content, src_dir, wf_file)
        (dist_workflows / wf_file.name).write_text(processed)
        count += 1
        print(f"  ✅ {wf_file.name}")
    
    return count


def build_templates(src_dir: Path, dist_dir: Path):
    """Copy templates to dist/project/docs/templates/."""
    src_templates = src_dir / "templates" / "documents"
    dist_templates = dist_dir / "project" / "docs" / "templates"
    
    if not src_templates.exists():
        print("  ⚠️  No src/templates/documents/ directory")
        return 0
    
    dist_templates.mkdir(parents=True, exist_ok=True)
    count = 0
    
    for template_file in src_templates.iterdir():
        if template_file.is_file():
            content = template_file.read_text()
            processed = process_includes(content, src_dir, template_file)
            (dist_templates / template_file.name).write_text(processed)
            count += 1
    
    print(f"  ✅ {count} templates")
    return count


def clean_dist(dist_dir: Path):
    """Remove dist/ directory."""
    if dist_dir.exists():
        shutil.rmtree(dist_dir)
        print("🗑️  Cleaned dist/")


def main():
    root = Path(__file__).parent.parent
    src_dir = root / "src"
    dist_dir = root / "dist"
    
    print("🔨 Building from src/ to dist/...")
    print(f"   Source: {src_dir}")
    print(f"   Output: {dist_dir}")
    print("")
    
    # Clean and recreate
    clean_dist(dist_dir)
    dist_dir.mkdir(parents=True, exist_ok=True)
    
    # Build each component
    print("📦 Building skills...")
    skills_count = build_skills(src_dir, dist_dir)
    
    print("\n📜 Building rules...")
    rules_count = build_rules(src_dir, dist_dir)
    
    print("\n⚡ Building workflows...")
    workflows_count = build_workflows(src_dir, dist_dir)
    
    print("\n📄 Building templates...")
    templates_count = build_templates(src_dir, dist_dir)
    
    # Summary
    print("\n" + "=" * 40)
    print(f"✅ Build complete!")
    print(f"   Skills:    {skills_count}")
    print(f"   Rules:     {rules_count}")
    print(f"   Workflows: {workflows_count}")
    print(f"   Templates: {templates_count}")
    print(f"\n📁 Output: {dist_dir}")


if __name__ == "__main__":
    main()
