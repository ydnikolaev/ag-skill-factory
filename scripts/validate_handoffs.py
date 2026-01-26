#!/usr/bin/env python3
"""
Validate that all handoffs are between skills that share at least one preset.
Uses preset-hierarchy.yaml to resolve inheritance.
"""

import yaml
from pathlib import Path


def load_preset_hierarchy():
    """Load preset hierarchy and compute full inheritance."""
    hierarchy_path = Path("src/_meta/preset-hierarchy.yaml")
    if not hierarchy_path.exists():
        return {}
    
    with open(hierarchy_path) as f:
        hierarchy = yaml.safe_load(f) or {}
    
    # Compute full inheritance for each preset
    def get_all_presets(preset_name, visited=None):
        if visited is None:
            visited = set()
        if preset_name in visited:
            return set()
        visited.add(preset_name)
        
        result = {preset_name}
        preset_def = hierarchy.get(preset_name, {})
        
        # Add inherited presets
        for parent in preset_def.get("inherits", []):
            result.update(get_all_presets(parent, visited.copy()))
        
        return result
    
    # Build full preset map
    full_presets = {}
    for preset_name in hierarchy:
        full_presets[preset_name] = get_all_presets(preset_name)
    
    return full_presets


def resolve_skill_presets(skill_presets, hierarchy):
    """Resolve all presets a skill belongs to (including inherited)."""
    resolved = set()
    for preset in skill_presets:
        resolved.update(hierarchy.get(preset, {preset}))
    return resolved


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


def main():
    # Load preset hierarchy
    hierarchy = load_preset_hierarchy()
    
    # Collect all skills and their presets
    skills = {}
    for path in list(Path("src/skills").iterdir()) + list(Path("src/skills/private").glob("*")):
        if not path.is_dir() or path.name.startswith("."):
            continue
        fm = extract_frontmatter(path)
        if fm and "name" in fm:
            leaf_presets = set(fm.get("presets", []))
            resolved_presets = resolve_skill_presets(leaf_presets, hierarchy)
            skills[fm["name"]] = {
                "leaf_presets": leaf_presets,
                "resolved_presets": resolved_presets,
                "receives_from": [h.get("skill") for h in fm.get("receives_from", []) or [] if isinstance(h, dict)],
                "delegates_to": [h.get("skill") for h in fm.get("delegates_to", []) or [] if isinstance(h, dict)],
                "return_paths": [h.get("skill") for h in fm.get("return_paths", []) or [] if isinstance(h, dict)],
            }

    # Validate
    errors = []
    for skill_name, skill_data in skills.items():
        skill_presets = skill_data["resolved_presets"]
        
        # Check receives_from
        for source in skill_data["receives_from"]:
            if source and source in skills:
                source_presets = skills[source]["resolved_presets"]
                common = skill_presets & source_presets
                if not common:
                    errors.append(f"❌ {skill_name} receives from {source}, but no common preset!")
                    errors.append(f"   {skill_name}: {skill_data['leaf_presets']} → {skill_presets}")
                    errors.append(f"   {source}: {skills[source]['leaf_presets']} → {source_presets}")
        
        # Check delegates_to
        for target in skill_data["delegates_to"]:
            if target and target in skills:
                target_presets = skills[target]["resolved_presets"]
                common = skill_presets & target_presets
                if not common:
                    errors.append(f"❌ {skill_name} delegates to {target}, but no common preset!")
                    errors.append(f"   {skill_name}: {skill_data['leaf_presets']} → {skill_presets}")
                    errors.append(f"   {target}: {skills[target]['leaf_presets']} → {target_presets}")
        
        # Check return_paths
        for target in skill_data["return_paths"]:
            if target and target in skills:
                target_presets = skills[target]["resolved_presets"]
                common = skill_presets & target_presets
                if not common:
                    errors.append(f"❌ {skill_name} has return_path to {target}, but no common preset!")
                    errors.append(f"   {skill_name}: {skill_data['leaf_presets']} → {skill_presets}")
                    errors.append(f"   {target}: {skills[target]['leaf_presets']} → {target_presets}")

    if errors:
        print("=== PRESET-HANDOFF MISMATCHES ===\n")
        for e in errors:
            print(e)
        print(f"\n❌ Found {len([e for e in errors if e.startswith('❌')])} issues")
        exit(1)
    else:
        print("✅ All handoffs have at least one common preset (with inheritance)")
        exit(0)


if __name__ == "__main__":
    main()
