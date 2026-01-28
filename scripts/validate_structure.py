#!/usr/bin/env python3
"""
Validate structure consistency across schemas.

Cross-references:
1. Skill paths (creates.path) → must exist in structure schema categories
2. Document template locations → must exist in structure schema
3. Structure schema categories → should be used by at least one skill/doc

Usage:
    python3 scripts/validate_structure.py
"""

import re
import sys
import yaml
from pathlib import Path


def load_yaml(path: Path) -> dict:
    """Load YAML file."""
    if not path.exists():
        return {}
    with open(path) as f:
        return yaml.safe_load(f) or {}


def extract_structure_categories() -> set:
    """Extract valid categories from folder-structure/ template (SSOT).
    
    SSOT: src/templates/folder-structure/active/ is the single source of truth
    for document categories. Validators check against real folder structure.
    """
    categories = set()
    
    # SSOT: Real folder structure at src/templates/folder-structure/active/
    active_dir = Path("src/templates/folder-structure/active")
    
    if not active_dir.exists():
        print(f"⚠️  SSOT folder-structure not found: {active_dir}")
        return categories
    
    # Scan real directories
    for item in active_dir.iterdir():
        if item.is_dir() and not item.name.startswith('.'):
            categories.add(item.name)
    
    return categories


def extract_skill_paths(skills_dir: Path) -> dict:
    """Extract paths from skill YAML frontmatter."""
    skill_paths = {}  # skill_name -> list of paths
    
    for skill_dir in skills_dir.iterdir():
        if not skill_dir.is_dir():
            continue
        
        skill_md = skill_dir / "SKILL.md"
        if not skill_md.exists():
            continue
        
        content = skill_md.read_text()
        if not content.startswith("---"):
            continue
        
        # Extract frontmatter
        lines = content.split('\n')
        for i, line in enumerate(lines[1:], 1):
            if line.strip() == "---":
                try:
                    fm = yaml.safe_load('\n'.join(lines[1:i])) or {}
                    paths = []
                    
                    # Extract creates[].path
                    creates = fm.get("creates", [])
                    if isinstance(creates, list):
                        for c in creates:
                            if isinstance(c, dict) and "path" in c:
                                paths.append(c["path"])
                    
                    if paths:
                        skill_paths[skill_dir.name] = paths
                except:
                    pass
                break
    
    return skill_paths


def extract_document_categories(enums_path: Path) -> set:
    """Extract categories from document enums."""
    enums = load_yaml(enums_path)
    return set(enums.get("categories", []))


def extract_category_from_path(path: str) -> str | None:
    """Extract category from a document path like project/docs/active/backend/."""
    match = re.search(r'project/docs/active/([^/]+)', path)
    if match:
        return match.group(1)
    return None


def main():
    errors = []
    warnings = []
    
    # SSOT: Read from real folder structure
    structure_categories = extract_structure_categories()
    
    # Load document enums
    doc_categories = extract_document_categories(Path("src/_meta/schema/documents/enums/enums.yaml"))
    
    # Extract skill paths
    skill_paths = extract_skill_paths(Path("src/skills"))
    
    # Track usage
    used_categories = set()
    
    # == VALIDATION 1: Skill paths must match structure categories ==
    for skill_name, paths in skill_paths.items():
        for path in paths:
            category = extract_category_from_path(path)
            if category:
                used_categories.add(category)
                if category not in structure_categories:
                    errors.append(f"{skill_name}: path '{path}' uses category '{category}' not in structure schema")
    
    # == VALIDATION 2: Document enum categories must match structure ==
    for cat in doc_categories:
        if cat not in structure_categories:
            errors.append(f"document enum category '{cat}' not in structure schema")
        used_categories.add(cat)
    
    # == VALIDATION 3: Structure categories should be used ==
    for cat in structure_categories:
        if cat not in used_categories:
            warnings.append(f"structure category '{cat}' is not used by any skill or document")
    
    # == REPORT ==
    if errors:
        print("=== VALIDATION ERRORS ===\n")
        for error in errors:
            print(f"❌ {error}")
    
    if warnings:
        print("\n=== WARNINGS ===\n")
        for warning in warnings:
            print(f"⚠️  {warning}")
    
    # Summary
    print(f"\n📊 Structure categories defined: {len(structure_categories)}")
    print(f"📊 Document enum categories: {len(doc_categories)}")
    print(f"📊 Skills with paths: {len(skill_paths)}")
    print(f"📊 Categories used: {len(used_categories)}")
    
    if errors:
        print(f"\n❌ {len(errors)} errors, {len(warnings)} warnings")
        sys.exit(1)
    else:
        print(f"\n✅ Structure validated ({len(warnings)} warnings)")
        sys.exit(0)


if __name__ == "__main__":
    main()
