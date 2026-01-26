#!/usr/bin/env python3
"""
Build script: src/ → dist/

Processes {{include: path}} directives and assembles final files.

Usage:
    python3 scripts/build.py                     # Full build (all skills)
    python3 scripts/build.py --team backend      # Only backend team skills
    python3 scripts/build.py --team tma,core     # Multiple teams
    python3 scripts/build.py --list-teams        # Show available teams
    
    # Or via make
    make build                    # Full build
    make build-team TEAM=backend  # Team-specific build
"""

import re
import sys
import shutil
import argparse
from pathlib import Path
from datetime import datetime

import yaml

# Include patterns:
# New: {{include: partials/pre-handoff-validation.md}}
# Legacy: <!-- INCLUDE: _meta/_skills/sections/pre-handoff-validation.md -->
INCLUDE_PATTERN = re.compile(r'\{\{include:\s*([^\}]+)\}\}')
LEGACY_INCLUDE_PATTERN = re.compile(r'<!--\s*INCLUDE:\s*([^>]+)\s*-->')

# YAML frontmatter pattern
FRONTMATTER_PATTERN = re.compile(r'^---\s*\n(.*?)\n---', re.DOTALL)


def parse_frontmatter(content: str) -> dict:
    """Extract YAML frontmatter from markdown content."""
    match = FRONTMATTER_PATTERN.match(content)
    if not match:
        return {}
    try:
        return yaml.safe_load(match.group(1)) or {}
    except yaml.YAMLError:
        return {}


def load_preset_hierarchy(src_dir: Path) -> dict:
    """Load preset hierarchy to resolve inheritance."""
    hierarchy_file = src_dir / "_meta" / "preset-hierarchy.yaml"
    if not hierarchy_file.exists():
        return {}
    
    with open(hierarchy_file) as f:
        return yaml.safe_load(f) or {}


def resolve_team_presets(team: str, hierarchy: dict) -> set:
    """
    Resolve all presets that should be included for a team.
    E.g., 'backend' → {'backend', 'core'} (because backend inherits core)
    """
    if team == "all":
        return {"all"}
    
    resolved = {team}
    
    # Get direct inheritance
    team_config = hierarchy.get(team, {})
    inherits = team_config.get("inherits", [])
    
    for parent in inherits:
        resolved.update(resolve_team_presets(parent, hierarchy))
    
    return resolved


def get_teams_for_preset(preset: str, hierarchy: dict) -> set:
    """
    Get all teams that include this preset.
    E.g., 'core' → {'core', 'backend', 'frontend', 'fullstack', 'tma', 'cli'}
    """
    teams = {preset}
    
    for team_name, team_config in hierarchy.items():
        inherits = team_config.get("inherits", [])
        if preset in inherits:
            teams.update(get_teams_for_preset(team_name, hierarchy))
    
    return teams


def skill_matches_teams(skill_presets: list, target_teams: set, hierarchy: dict) -> bool:
    """Check if a skill belongs to any of the target teams."""
    if "all" in target_teams:
        return True
    
    for preset in skill_presets:
        # Direct match
        if preset in target_teams:
            return True
        
        # Check if any target team inherits from this preset
        teams_for_preset = get_teams_for_preset(preset, hierarchy)
        if teams_for_preset & target_teams:
            return True
        
        # Check if this skill's preset is a parent of target teams
        for target in target_teams:
            resolved = resolve_team_presets(target, hierarchy)
            if preset in resolved:
                return True
    
    return False


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


def build_skills(src_dir: Path, dist_dir: Path, target_teams: set = None, hierarchy: dict = None):
    """Build skills from src/ to dist/skills/."""
    src_skills = src_dir / "skills"
    dist_skills = dist_dir / "skills"
    
    if not src_skills.exists():
        print("  ⚠️  No src/skills/ directory")
        return 0, 0
    
    dist_skills.mkdir(parents=True, exist_ok=True)
    count = 0
    skipped = 0
    
    # Process all skill directories (including private/*)
    skill_dirs = []
    
    for item in src_skills.iterdir():
        if not item.is_dir():
            continue
        if item.name == "private":
            # Add private skills as top-level
            for private_skill in item.iterdir():
                if private_skill.is_dir():
                    skill_dirs.append((private_skill, True))
        else:
            skill_dirs.append((item, False))
    
    for skill_dir, is_private in skill_dirs:
        skill_name = skill_dir.name
        skill_md = skill_dir / "SKILL.md"
        
        if not skill_md.exists():
            continue
        
        # Read and parse frontmatter
        content = skill_md.read_text()
        frontmatter = parse_frontmatter(content)
        skill_presets = frontmatter.get("presets", [])
        
        # Filter by team if specified
        if target_teams and hierarchy:
            if not skill_matches_teams(skill_presets, target_teams, hierarchy):
                skipped += 1
                continue
        
        # Process includes
        processed = process_includes(content, src_dir, skill_md)
        
        # Write to dist (private skills go to top-level, not private/)
        dist_skill_dir = dist_skills / skill_name
        dist_skill_dir.mkdir(parents=True, exist_ok=True)
        (dist_skill_dir / "SKILL.md").write_text(processed)
        
        # Copy examples if exist
        examples_dir = skill_dir / "examples"
        if examples_dir.exists():
            shutil.copytree(examples_dir, dist_skill_dir / "examples", dirs_exist_ok=True)
        
        # Copy references if exist
        references_dir = skill_dir / "references"
        if references_dir.exists():
            shutil.copytree(references_dir, dist_skill_dir / "references", dirs_exist_ok=True)
        
        # Copy resources if exist
        resources_dir = skill_dir / "resources"
        if resources_dir.exists():
            shutil.copytree(resources_dir, dist_skill_dir / "resources", dirs_exist_ok=True)
        
        count += 1
        marker = "🔒" if is_private else "✅"
        presets_str = f" [{', '.join(skill_presets)}]" if skill_presets else ""
        print(f"  {marker} {skill_name}{presets_str}")
    
    return count, skipped

def build_rules(src_dir: Path, dist_dir: Path, target_team: str = None):
    """Build rules from src/ to dist/rules/."""
    src_rules = src_dir / "rules"
    dist_rules = dist_dir / "rules"
    
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
    
    # Copy team-specific TEAM and PIPELINE files
    team_name = target_team if target_team else "all"
    team_count = build_team_rules(src_dir, dist_rules, team_name)
    
    return count + team_count


def build_team_rules(src_dir: Path, dist_rules: Path, team_name: str):
    """Copy TEAM_*.md and PIPELINE_*.md for the target team to rules/."""
    count = 0
    
    # TEAM file
    team_file = src_dir / "_meta" / "teams" / f"TEAM_{team_name}.md"
    if team_file.exists():
        content = team_file.read_text()
        (dist_rules / "TEAM.md").write_text(content)
        count += 1
        print(f"  📋 TEAM.md (from {team_name})")
    else:
        print(f"  ⚠️  TEAM_{team_name}.md not found")
    
    # PIPELINE file
    pipeline_file = src_dir / "_meta" / "pipelines" / f"PIPELINE_{team_name}.md"
    if pipeline_file.exists():
        content = pipeline_file.read_text()
        (dist_rules / "PIPELINE.md").write_text(content)
        count += 1
        print(f"  🔀 PIPELINE.md (from {team_name})")
    else:
        print(f"  ⚠️  PIPELINE_{team_name}.md not found")
    
    return count

def build_workflows(src_dir: Path, dist_dir: Path):
    """Build workflows from src/ to dist/workflows/."""
    src_workflows = src_dir / "workflows"
    dist_workflows = dist_dir / "workflows"
    
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
    """Copy templates to dist/docs/templates/."""
    src_templates = src_dir / "templates" / "documents"
    dist_templates = dist_dir / "docs" / "templates"
    
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
    
    print(f"  ✅ {count} document templates")
    
    # Copy folder-structure template
    src_folder_struct = src_dir / "templates" / "folder-structure"
    dist_folder_struct = dist_dir / "docs" / "folder-structure"
    
    if src_folder_struct.exists():
        shutil.copytree(src_folder_struct, dist_folder_struct, dirs_exist_ok=True)
        print(f"  ✅ folder-structure template")
    
    return count

def build_configs(src_dir: Path, dist_dir: Path):
    """Copy configs from src/ to dist/configs/."""
    src_configs = src_dir / "configs"
    dist_configs = dist_dir / "configs"
    
    if not src_configs.exists():
        return 0
    
    dist_configs.mkdir(parents=True, exist_ok=True)
    count = 0
    
    for config_file in src_configs.iterdir():
        if config_file.is_file():
            shutil.copy2(config_file, dist_configs / config_file.name)
            count += 1
            print(f"  ✅ {config_file.name}")
    
    return count


def clean_dist(dist_dir: Path):
    """Remove dist/ directory."""
    if dist_dir.exists():
        shutil.rmtree(dist_dir)
        print("🗑️  Cleaned dist/")


def list_available_teams(src_dir: Path):
    """List all available teams from preset-hierarchy.yaml."""
    hierarchy = load_preset_hierarchy(src_dir)
    
    print("\n📋 Available Teams:\n")
    for team_name, config in sorted(hierarchy.items()):
        desc = config.get("description", "")
        inherits = config.get("inherits", [])
        inherits_str = f" (extends: {', '.join(inherits)})" if inherits else ""
        print(f"  • {team_name}: {desc}{inherits_str}")
    
    print("\n💡 Usage: python3 scripts/build.py --team <team-name>")
    print("   Example: python3 scripts/build.py --team backend")


def main():
    parser = argparse.ArgumentParser(description="Build src/ to dist/")
    parser.add_argument(
        "--team", "-t",
        help="Build only skills for specific team(s). Comma-separated for multiple.",
        default=None
    )
    parser.add_argument(
        "--list-teams",
        action="store_true",
        help="List available teams and exit"
    )
    args = parser.parse_args()
    
    root = Path(__file__).parent.parent
    src_dir = root / "src"
    dist_dir = root / "dist"
    
    # List teams mode
    if args.list_teams:
        list_available_teams(src_dir)
        return
    
    # Load hierarchy for team filtering
    hierarchy = load_preset_hierarchy(src_dir)
    
    # Parse target teams
    target_teams = None
    if args.team:
        target_teams = set(t.strip() for t in args.team.split(","))
        # Validate teams
        for team in target_teams:
            if team not in hierarchy and team != "all":
                print(f"❌ Unknown team: {team}")
                print(f"   Available: {', '.join(sorted(hierarchy.keys()))}")
                sys.exit(1)
    
    # Header
    if target_teams:
        teams_str = ", ".join(sorted(target_teams))
        print(f"🔨 Building for team(s): {teams_str}")
    else:
        print("🔨 Building from src/ to dist/...")
    print(f"   Source: {src_dir}")
    print(f"   Output: {dist_dir}")
    print("")
    
    # Clean and recreate
    clean_dist(dist_dir)
    dist_dir.mkdir(parents=True, exist_ok=True)
    
    # Build skills (with team filtering)
    print("📦 Building skills...")
    skills_count, skills_skipped = build_skills(src_dir, dist_dir, target_teams, hierarchy)
    
    # Always build other components (rules, workflows, etc.)
    # For rules, pass the target team (use first if multiple, or None for "all")
    target_team = list(target_teams)[0] if target_teams and len(target_teams) == 1 else None
    print("\n📜 Building rules...")
    rules_count = build_rules(src_dir, dist_dir, target_team)
    
    print("\n⚡ Building workflows...")
    workflows_count = build_workflows(src_dir, dist_dir)
    
    print("\n📄 Building templates...")
    templates_count = build_templates(src_dir, dist_dir)
    
    print("\n⚙️  Building configs...")
    configs_count = build_configs(src_dir, dist_dir)
    
    # Summary
    print("\n" + "=" * 40)
    print(f"✅ Build complete!")
    print(f"   Skills:    {skills_count}" + (f" (skipped {skills_skipped})" if skills_skipped else ""))
    print(f"   Rules:     {rules_count}")
    print(f"   Workflows: {workflows_count}")
    print(f"   Templates: {templates_count}")
    print(f"   Configs:   {configs_count}")
    print(f"\n📁 Output: {dist_dir}")


if __name__ == "__main__":
    main()
