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

SRC_DIR = Path(__file__).parent.parent / "src"
SKILLS_DIR = SRC_DIR / "skills"
HIERARCHY_FILE = SRC_DIR / "_meta" / "preset-hierarchy.yaml"
OUTPUT_FILE = SRC_DIR / "_meta" / "presets.yaml"


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
        
        # Add extends if present in config (it's called 'inherits' in hierarchy, 'extends' in presets.yaml)
        inherits = config.get("inherits")
        if inherits:
            # If single item list, unwrap it to string if that was the original style, 
            # OR keep as list. The original file showed:
            # - extends: core (single)
            # - extends: [backend, frontend] (multiple)
            # - extends: core (single)
            # Let's match that behavior.
            if len(inherits) == 1:
                preset_data["extends"] = inherits[0]
            else:
                preset_data["extends"] = inherits

        if config.get("includes_all"):
            preset_data["skills"] = "*"
        else:
            # Get all skills for this preset (including inherited)
            all_resolved_skills = preset_skills.get(preset_name, set())
            
            # Find inherited skills to exclude
            inherited_skills = set()
            if inherits:
                # inherits can be single string or list. Normalize to list.
                if isinstance(inherits, str):
                    parents_list = [inherits]
                else:
                    parents_list = inherits
                
                # For each direct parent, get ITS full set of skills (which includes its parents)
                # We can just use preset_skills[parent] because we populated it fully in the first pass
                for p in parents_list:
                    inherited_skills.update(preset_skills.get(p, set()))

            # Delta skills = All resolved - Inherited
            delta_skills = sorted(list(all_resolved_skills - inherited_skills))

            if delta_skills:
                preset_data["skills"] = delta_skills
            elif not inherits:
                 # If no inheritance and no skills, maybe show empty list?
                 # core has skills. minimal has skills.
                 # If delta is empty and we extend something, we omit skills key (like backend in manual)
                 pass
            else:
                 # If delta is empty AND we inherit, we omit the key.
                 pass


        
        output[preset_name] = preset_data
    
    # Custom Dumper to handle specific formatting (flow style for extends list?)
    # or just use default Block style. "extends: [backend, frontend]" is flow style for that key.
    # PyYAML default is block style for lists.
    # To exactly match, we might need a custom representer, but let's try default first.
    # The requirement is "identic to before". 
    # Before:
    # extends: core
    # extends:
    #   - backend
    #   - frontend
    # OR
    # extends: [backend, frontend]
    # Let's start with standard YAML.

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
