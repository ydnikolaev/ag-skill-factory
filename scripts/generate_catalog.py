#!/usr/bin/env python3
"""
Generate skill documentation pages from blueprint.

Usage:
    python3 scripts/generate_catalog.py
"""

import os
import re
from pathlib import Path


def extract_frontmatter(content: str) -> dict:
    """Extract YAML frontmatter from SKILL.md."""
    if not content.startswith("---"):
        return {}
    
    end = content.find("\n---\n", 4)
    if end == -1:
        return {}
    
    frontmatter = content[4:end]
    result = {}
    for line in frontmatter.split("\n"):
        if ":" in line:
            key, value = line.split(":", 1)
            result[key.strip()] = value.strip()
    return result


def generate_skill_page(skill_path: Path, docs_path: Path) -> dict:
    """Generate a skill documentation page."""
    skill_md = skill_path / "SKILL.md"
    if not skill_md.exists():
        return None
    
    content = skill_md.read_text()
    frontmatter = extract_frontmatter(content)
    
    name = frontmatter.get("name", skill_path.name)
    description = frontmatter.get("description", "No description")
    version = frontmatter.get("version", "1.0.0")
    
    # Extract body after frontmatter
    body_start = content.find("\n---\n", 4)
    if body_start != -1:
        body = content[body_start + 5:]
    else:
        body = content
    
    # Create skill page
    skill_doc = f"""# {name}

> {description}

**Version:** {version}

---

{body}
"""
    
    # Write skill page
    skill_doc_path = docs_path / "skills" / f"{skill_path.name}.md"
    skill_doc_path.parent.mkdir(parents=True, exist_ok=True)
    skill_doc_path.write_text(skill_doc)
    
    return {"name": name, "slug": skill_path.name, "description": description, "version": version}


def main():
    root = Path(__file__).parent.parent
    blueprint_skills = root / "blueprint" / "skills"
    website_path = root / "website"
    
    if not blueprint_skills.exists():
        print("❌ blueprint/skills not found")
        return
    
    skills = []
    for skill_path in sorted(blueprint_skills.iterdir()):
        if skill_path.is_dir() and not skill_path.name.startswith("."):
            info = generate_skill_page(skill_path, website_path)
            if info:
                skills.append(info)
                print(f"✅ {info['name']}")
    
    print(f"\n📚 Generated {len(skills)} skill pages in website/skills/")


if __name__ == "__main__":
    main()
