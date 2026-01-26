#!/usr/bin/env python3
"""
Generate presets.yaml from skill frontmatter + preset-hierarchy.yaml

Usage: python3 scripts/generate-presets.py

Reads:
  - blueprint/skills/*/SKILL.md (frontmatter: presets: [backend, ...])
  - blueprint/_meta/preset-hierarchy.yaml (inheritance rules)

Writes:
  - blueprint/_meta/presets.yaml (generated, no extends)
"""

import os
import re
import yaml
from pathlib import Path
from collections import defaultdict

BLUEPRINT_DIR = Path(__file__).parent.parent / "blueprint"
SKILLS_DIR = BLUEPRINT_DIR / "skills"
HIERARCHY_FILE = BLUEPRINT_DIR / "_meta" / "preset-hierarchy.yaml"
OUTPUT_FILE = BLUEPRINT_DIR / "_meta" / "presets.yaml"


def extract_frontmatter(skill_path: Path) -> dict:
    """Extract YAML frontmatter from SKILL.md"""
    content = skill_path.read_text()
    match = re.match(r'^---\n(.*?)\n---', content, re.DOTALL)
    if match:
        return yaml.safe_load(match.group(1)) or {}
    return {}


def load_hierarchy() -> dict:
    """Load preset hierarchy with inheritance rules"""
    if HIERARCHY_FILE.exists():
        return yaml.safe_load(HIERARCHY_FILE.read_text()) or {}
    return {}


def resolve_parent_presets(preset: str, hierarchy: dict, visited: set = None) -> set:
    """Recursively find all presets that inherit from the given preset"""
    if visited is None:
        visited = set()
    
    parents = set()
    for name, config in hierarchy.items():
        if name in visited:
            continue
        inherits = config.get("inherits", [])
        if preset in inherits:
            parents.add(name)
            visited.add(name)
            # Recursively find parents of this parent
            parents.update(resolve_parent_presets(name, hierarchy, visited))
    
    return parents


def main():
    hierarchy = load_hierarchy()
    
    # Collect skills by preset
    preset_skills = defaultdict(set)
    all_skills = []
    
    for skill_dir in sorted(SKILLS_DIR.iterdir()):
        skill_file = skill_dir / "SKILL.md"
        if not skill_file.exists():
            continue
        
        frontmatter = extract_frontmatter(skill_file)
        skill_name = frontmatter.get("name") or skill_dir.name
        all_skills.append(skill_name)
        
        # Get declared presets
        declared_presets = frontmatter.get("presets", [])
        if isinstance(declared_presets, str):
            declared_presets = [declared_presets]
        
        # Add to declared presets
        for preset in declared_presets:
            preset_skills[preset].add(skill_name)
            
            # Resolve parent presets via inheritance
            parents = resolve_parent_presets(preset, hierarchy)
            for parent in parents:
                preset_skills[parent].add(skill_name)
    
    # Build output
    output = {
        "_generated": "DO NOT EDIT - generated from skill frontmatter",
        "_source": "scripts/generate-presets.py",
    }
    
    # Add presets in order
    preset_order = ["all", "core", "backend", "frontend", "fullstack", "tma", "cli", "minimal"]
    
    for preset_name in preset_order:
        if preset_name not in hierarchy:
            continue
            
        config = hierarchy[preset_name]
        preset_data = {
            "description": config.get("description", ""),
        }
        
        if config.get("includes_all"):
            preset_data["skills"] = "*"
        else:
            skills = sorted(preset_skills.get(preset_name, []))
            if skills:
                preset_data["skills"] = skills
            else:
                preset_data["skills"] = []
        
        output[preset_name] = preset_data
    
    # Write output
    OUTPUT_FILE.write_text(
        "# Auto-generated from skill frontmatter\n"
        "# Run: python3 scripts/generate-presets.py\n\n"
        + yaml.dump(output, default_flow_style=False, sort_keys=False, allow_unicode=True)
    )
    
    print(f"✅ Generated {OUTPUT_FILE}")
    print(f"   Skills: {len(all_skills)}")
    print(f"   Presets: {len([p for p in preset_order if p in hierarchy])}")


if __name__ == "__main__":
    main()
