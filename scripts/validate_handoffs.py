#!/usr/bin/env python3
"""Validate that all handoffs are between skills that share at least one preset."""

import yaml
from pathlib import Path

def extract_frontmatter(skill_path):
    skill_md = skill_path / "SKILL.md"
    if not skill_md.exists():
        return None
    content = skill_md.read_text()
    if not content.startswith("---"):
        return None
    lines = content.split('\n')
    for i, line in enumerate(lines[1:], 1):
        if line.strip() == "---":
            try:
                return yaml.safe_load('\n'.join(lines[1:i])) or {}
            except:
                return None
    return None

# Collect all skills and their presets
skills = {}
for path in list(Path("src/skills").iterdir()) + list(Path("src/skills/private").glob("*")):
    if not path.is_dir() or path.name.startswith("."):
        continue
    fm = extract_frontmatter(path)
    if fm and "name" in fm:
        skills[fm["name"]] = {
            "presets": set(fm.get("presets", [])),
            "receives_from": [h.get("skill") for h in fm.get("receives_from", []) or [] if isinstance(h, dict)],
            "delegates_to": [h.get("skill") for h in fm.get("delegates_to", []) or [] if isinstance(h, dict)],
            "return_paths": [h.get("skill") for h in fm.get("return_paths", []) or [] if isinstance(h, dict)],
        }

# Validate
errors = []
for skill_name, skill_data in skills.items():
    skill_presets = skill_data["presets"]
    
    # Check receives_from
    for source in skill_data["receives_from"]:
        if source and source in skills:
            source_presets = skills[source]["presets"]
            common = skill_presets & source_presets
            if not common:
                errors.append(f"❌ {skill_name} receives from {source}, but no common preset!")
                errors.append(f"   {skill_name}: {skill_data['presets']}")
                errors.append(f"   {source}: {source_presets}")
    
    # Check delegates_to
    for target in skill_data["delegates_to"]:
        if target and target in skills:
            target_presets = skills[target]["presets"]
            common = skill_presets & target_presets
            if not common:
                errors.append(f"❌ {skill_name} delegates to {target}, but no common preset!")
                errors.append(f"   {skill_name}: {skill_data['presets']}")
                errors.append(f"   {target}: {target_presets}")
    
    # Check return_paths
    for target in skill_data["return_paths"]:
        if target and target in skills:
            target_presets = skills[target]["presets"]
            common = skill_presets & target_presets
            if not common:
                errors.append(f"❌ {skill_name} has return_path to {target}, but no common preset!")
                errors.append(f"   {skill_name}: {skill_data['presets']}")
                errors.append(f"   {target}: {target_presets}")

if errors:
    print("=== PRESET-HANDOFF MISMATCHES ===\n")
    for e in errors:
        print(e)
    print(f"\n❌ Found {len([e for e in errors if e.startswith('❌')])} issues")
else:
    print("✅ All handoffs have at least one common preset")
