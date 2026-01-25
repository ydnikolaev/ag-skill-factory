#!/usr/bin/env python3
"""
Validate entire blueprint consistency.

Checks:
1. presets.yaml completeness - every skill in blueprint/skills/ mentioned in at least one preset
2. TEAM.md synced - would not change if regenerated

Usage:
    python3 validate_blueprint.py <blueprint-path>
    python3 validate_blueprint.py blueprint/
"""
import os
import sys
import subprocess
import tempfile
import yaml


def check_presets_completeness(blueprint_path: str) -> list:
    """Check that every skill is mentioned in at least one preset."""
    errors = []
    
    skills_dir = os.path.join(blueprint_path, "skills")
    presets_file = os.path.join(blueprint_path, "_meta", "presets.yaml")
    
    if not os.path.isdir(skills_dir):
        errors.append(f"Skills directory not found: {skills_dir}")
        return errors
    
    if not os.path.exists(presets_file):
        errors.append(f"Presets file not found: {presets_file}")
        return errors
    
    # Get all skills
    all_skills = set()
    for item in os.listdir(skills_dir):
        skill_path = os.path.join(skills_dir, item)
        if os.path.isdir(skill_path) and os.path.exists(os.path.join(skill_path, "SKILL.md")):
            all_skills.add(item)
    
    # Get skills from presets
    with open(presets_file, 'r') as f:
        presets = yaml.safe_load(f)
    
    mentioned_skills = set()
    for preset_name, preset_data in presets.items():
        if isinstance(preset_data, dict):
            skills = preset_data.get("skills", [])
            if skills == "*":
                # "all" preset covers everything
                mentioned_skills = all_skills.copy()
                break
            elif isinstance(skills, list):
                mentioned_skills.update(skills)
    
    # Find missing skills
    missing = all_skills - mentioned_skills
    if missing:
        errors.append(f"Skills not in any preset: {sorted(missing)}")
    
    return errors


def check_team_synced(blueprint_path: str) -> list:
    """Check that TEAM.md is in sync with skills directory."""
    errors = []
    
    team_file = os.path.join(blueprint_path, "rules", "TEAM.md")
    skills_dir = os.path.join(blueprint_path, "skills")
    
    if not os.path.exists(team_file):
        errors.append(f"TEAM.md not found: {team_file}")
        return errors
    
    if not os.path.isdir(skills_dir):
        errors.append(f"Skills directory not found: {skills_dir}")
        return errors
    
    # Get skills from directory
    dir_skills = set()
    for item in os.listdir(skills_dir):
        skill_path = os.path.join(skills_dir, item)
        if os.path.isdir(skill_path) and os.path.exists(os.path.join(skill_path, "SKILL.md")):
            dir_skills.add(item)
    
    # Get skills from TEAM.md
    team_skills = set()
    with open(team_file, 'r') as f:
        for line in f:
            if line.startswith("| `") and "` |" in line:
                skill_name = line.split("`")[1]
                team_skills.add(skill_name)
    
    # Compare
    missing_in_team = dir_skills - team_skills
    extra_in_team = team_skills - dir_skills
    
    if missing_in_team:
        errors.append(f"Skills missing in TEAM.md: {sorted(missing_in_team)}")
    if extra_in_team:
        errors.append(f"Extra skills in TEAM.md (not in directory): {sorted(extra_in_team)}")
    
    return errors


def validate_blueprint(blueprint_path: str) -> bool:
    """Validate entire blueprint consistency."""
    print(f"🔍 Validating blueprint at {blueprint_path}...")
    errors = []
    
    # Check 1: Presets completeness
    print("📦 Checking presets.yaml completeness...")
    preset_errors = check_presets_completeness(blueprint_path)
    if preset_errors:
        errors.extend(preset_errors)
    else:
        print("✅ All skills covered by presets")
    
    # Check 2: TEAM.md sync
    print("📋 Checking TEAM.md sync...")
    team_errors = check_team_synced(blueprint_path)
    if team_errors:
        errors.extend(team_errors)
    else:
        print("✅ TEAM.md is in sync")
    
    # Print results
    if errors:
        print("\n❌ ERRORS:")
        for e in errors:
            print(f"   • {e}")
        return False
    else:
        print("\n🎉 Blueprint is consistent!")
        return True


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: validate_blueprint.py <blueprint-path>")
        print("Example: validate_blueprint.py blueprint/")
        sys.exit(1)
    
    success = validate_blueprint(sys.argv[1])
    sys.exit(0 if success else 1)
