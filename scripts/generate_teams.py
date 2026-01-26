#!/usr/bin/env python3
"""
Generate TEAM files for each preset from presets.yaml.

Usage:
    python3 scripts/generate_teams.py
    
Output:
    blueprint/_meta/_teams/TEAM_<preset>.md for each preset
"""

import os
import yaml
from pathlib import Path


def extract_description(skill_path: Path) -> str:
    """Extract description from SKILL.md frontmatter."""
    skill_md = skill_path / "SKILL.md"
    if not skill_md.exists():
        return "No description"
    
    content = skill_md.read_text()
    if not content.startswith("---"):
        return "No description"
    
    end = content.find("\n---\n", 4)
    if end == -1:
        return "No description"
    
    frontmatter = content[4:end]
    for line in frontmatter.split("\n"):
        if line.startswith("description:"):
            return line.split(":", 1)[1].strip()
    
    return "No description"


def resolve_skills(preset_name: str, presets: dict, blueprint_skills: Path) -> list:
    """Resolve full skill list for a preset, including extends."""
    preset = presets.get(preset_name, {})
    skills = set()
    
    # Handle 'all' preset
    if preset.get("skills") == "*":
        for skill_path in blueprint_skills.iterdir():
            if skill_path.is_dir() and not skill_path.name.startswith("."):
                skills.add(skill_path.name)
        return sorted(skills)
    
    # Handle extends
    extends = preset.get("extends", [])
    if isinstance(extends, str):
        extends = [extends]
    
    for parent in extends:
        parent_skills = resolve_skills(parent, presets, blueprint_skills)
        skills.update(parent_skills)
    
    # Add own skills
    own_skills = preset.get("skills", [])
    if isinstance(own_skills, list):
        skills.update(own_skills)
    
    return sorted(skills)


def generate_team_file(preset_name: str, skill_names: list, blueprint_skills: Path, output_dir: Path, preset_desc: str):
    """Generate TEAM_<preset>.md file."""
    output_file = output_dir / f"TEAM_{preset_name}.md"
    
    lines = [
        "---",
        "trigger: model_decision",
        f"description: Team roster for {preset_name} preset. Apply when collaborating or delegating.",
        "---",
        "",
        f"# Team Roster ({preset_name})",
        "",
        f"> {preset_desc}",
        "",
        "| Skill | Description |",
        "|-------|-------------|",
    ]
    
    for skill_name in skill_names:
        skill_path = blueprint_skills / skill_name
        desc = extract_description(skill_path)
        lines.append(f"| `{skill_name}` | {desc} |")
    
    lines.extend([
        "",
        "## Usage",
        "",
        "Reference skills with `@skill-name` in skill collaboration sections.",
        "",
    ])
    
    output_file.write_text("\n".join(lines))
    return len(skill_names)


def main():
    root = Path(__file__).parent.parent
    blueprint_skills = root / "blueprint" / "skills"
    presets_file = root / "blueprint" / "_meta" / "presets.yaml"
    output_dir = root / "blueprint" / "_meta" / "_teams"
    
    if not blueprint_skills.exists():
        print("❌ blueprint/skills not found")
        return
    
    if not presets_file.exists():
        print("❌ blueprint/_meta/presets.yaml not found")
        return
    
    # Create output directory
    output_dir.mkdir(parents=True, exist_ok=True)
    
    # Load presets
    with open(presets_file) as f:
        presets = yaml.safe_load(f)
    
    print("📝 Generating team files...")
    
    for preset_name, preset_config in presets.items():
        skill_names = resolve_skills(preset_name, presets, blueprint_skills)
        preset_desc = preset_config.get("description", f"{preset_name} preset")
        count = generate_team_file(preset_name, skill_names, blueprint_skills, output_dir, preset_desc)
        print(f"  ✅ TEAM_{preset_name}.md ({count} skills)")
    
    print(f"\n📁 Generated {len(presets)} team files in blueprint/_meta/_teams/")


if __name__ == "__main__":
    main()
