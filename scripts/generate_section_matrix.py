#!/usr/bin/env python3
"""
Generate section-matrix.yaml from SKILL.md frontmatter required_sections.

Usage:
    python3 scripts/generate_section_matrix.py
    
Output:
    src/_meta/section-matrix.yaml - Auto-generated skill required sections
"""

import os
import yaml
from pathlib import Path
from typing import Dict, List, Any


# Section order (top to bottom in SKILL.md)
SECTION_ORDER = [
    "frontmatter",
    "header",
    "mode_block",
    "when_to_activate",
    "role_boundary",
    "brain_to_docs",
    "tech_stack",
    "critical_rules",
    "workflow",
    "protocols",
    "team_collaboration",
    "when_to_delegate",
    "document_lifecycle",
    "pre_handoff",
    "handoff_protocol",
    "artifact_ownership",
    "resources",
    "language_requirements",
    "best_practices",
]

# Section header mappings (for grep validation)
SECTION_HEADERS = {
    "frontmatter": "^---",
    "language_requirements": "## Language Requirements|INCLUDE.*language-requirements",
    "tech_stack": "## Tech Stack",
    "workflow": "## Workflow",
    "team_collaboration": "## Team Collaboration",
    "when_to_delegate": "## When to Delegate",
    "brain_to_docs": "## (Iteration Protocol|Brain to Docs)|INCLUDE.*brain-to-docs",
    "document_lifecycle": "## Document Lifecycle",
    "handoff_protocol": "## Handoff Protocol",
    "pre_handoff": "## Pre-Handoff Validation",
    "protocols": "## (TDD|Git|Tech Debt) Protocol",
    "critical_rules": "## Critical Rules",
    "when_to_activate": "## When to Activate",
    "role_boundary": "## Role Boundary",
    "artifact_ownership": "## Artifact Ownership",
    "best_practices": "## Antigravity Best Practices",
}


def extract_frontmatter(skill_path: Path) -> dict:
    """Extract YAML frontmatter from SKILL.md."""
    skill_md = skill_path / "SKILL.md"
    if not skill_md.exists():
        return {}
    
    content = skill_md.read_text()
    if not content.startswith("---"):
        return {}
    
    # Find closing ---
    end = content.find("\n---\n", 4)
    if end == -1:
        # Try alternative: just "---" at start of line
        lines = content.split('\n')
        end_line = None
        for i, line in enumerate(lines[1:], 1):
            if line.strip() == "---":
                end_line = i
                break
        if end_line is None:
            return {}
        frontmatter_text = '\n'.join(lines[1:end_line])
    else:
        frontmatter_text = content[4:end]
    
    try:
        return yaml.safe_load(frontmatter_text) or {}
    except yaml.YAMLError as e:
        print(f"  ⚠️ YAML error in {skill_path.name}: {e}")
        return {}


def build_section_matrix(skills_path: Path, private_skills_path: Path = None) -> dict:
    """Build section matrix from all SKILL.md files."""
    matrix = {
        "section_order": SECTION_ORDER,
        "categories": {},
        "section_headers": SECTION_HEADERS,
    }
    
    # Collect skills by category
    category_skills: Dict[str, List[str]] = {}
    category_required: Dict[str, set] = {}
    
    all_skills_paths = [skills_path]
    if private_skills_path and private_skills_path.exists():
        all_skills_paths.append(private_skills_path)
    
    for base_path in all_skills_paths:
        for skill_path in sorted(base_path.iterdir()):
            if not skill_path.is_dir() or skill_path.name.startswith("."):
                continue
            
            fm = extract_frontmatter(skill_path)
            if not fm:
                continue
            
            skill_name = skill_path.name
            category = fm.get("category", "utility")
            required_sections = fm.get("required_sections", [])
            
            # Initialize category if needed
            if category not in category_skills:
                category_skills[category] = []
                category_required[category] = set()
            
            category_skills[category].append(skill_name)
            
            # Add required sections
            for section in required_sections:
                category_required[category].add(section)
    
    # Build final matrix structure
    for category in sorted(category_skills.keys()):
        matrix["categories"][category] = {
            "skills": sorted(category_skills[category]),
            "required": sorted(category_required[category]),
        }
    
    return matrix


def main():
    # Find project root
    script_path = Path(__file__).parent
    project_root = script_path.parent
    
    skills_path = project_root / "src" / "skills"
    private_skills_path = project_root / "src" / "skills" / "private"
    output_path = project_root / "src" / "_meta" / "section-matrix.yaml"
    
    print("📊 Generating section-matrix.yaml...")
    
    matrix = build_section_matrix(skills_path, private_skills_path)
    
    # Write YAML with header comment
    output_content = """# Skill Section Matrix (AUTO-GENERATED)
# 
# DO NOT EDIT MANUALLY!
# Generated from SKILL.md frontmatter required_sections.
# Run: python3 scripts/generate_section_matrix.py
#
# Defines the order and applicability of sections across skill types.
# Used by validation and skill-updater for compliance checks.

"""
    output_content += yaml.dump(matrix, default_flow_style=False, sort_keys=False, allow_unicode=True)
    
    output_path.write_text(output_content)
    
    # Print summary
    total_skills = sum(len(cat["skills"]) for cat in matrix["categories"].values())
    print(f"  ✅ section-matrix.yaml ({len(matrix['categories'])} categories, {total_skills} skills)")
    
    for cat_name, cat_data in matrix["categories"].items():
        print(f"     - {cat_name}: {len(cat_data['skills'])} skills, {len(cat_data['required'])} required sections")


if __name__ == "__main__":
    main()
