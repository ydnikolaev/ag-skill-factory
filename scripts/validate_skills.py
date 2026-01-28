#!/usr/bin/env python3
"""
Validate all skills against Schema V3.

Usage:
    python3 scripts/validate_skills.py              # Validate all
    python3 scripts/validate_skills.py --skill=name # Validate single
    python3 scripts/validate_skills.py --check-docs # Include doc checks

Validations:
    - JSON Schema structure
    - Enum values
    - Name format (lowercase-with-hyphens)
    - Version format (semver)
    - Required fields
    - Preset validity
    - Doc type existence
"""

import argparse
import re
import sys
import yaml
from pathlib import Path


def load_yaml(path: Path) -> dict:
    """Load YAML file."""
    with open(path) as f:
        return yaml.safe_load(f) or {}


def extract_frontmatter(skill_path: Path) -> dict:
    """Extract YAML frontmatter from SKILL.md."""
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


def validate_name(name: str) -> list:
    """Validate skill name format."""
    errors = []
    if not name:
        errors.append("name is required")
    elif not re.match(r'^[a-z][a-z0-9-]*$', name):
        errors.append(f"name '{name}' must be lowercase-with-hyphens")
    return errors


def validate_version(version: str) -> list:
    """Validate version format."""
    errors = []
    if not version:
        errors.append("version is required")
    elif not re.match(r'^\d+\.\d+\.\d+$', version):
        errors.append(f"version '{version}' must be semver (e.g. 1.0.0)")
    return errors


def validate_description(desc: str) -> list:
    """Validate description."""
    errors = []
    if not desc:
        errors.append("description is required")
    elif len(desc) < 10:
        errors.append("description must be at least 10 characters")
    return errors


def validate_enums(fm: dict, factory: dict, runtime: dict) -> list:
    """Validate enum values against schema."""
    errors = []
    
    # Phase
    phase = fm.get("phase")
    if phase and phase not in factory.get("phases", []):
        errors.append(f"invalid phase '{phase}' (valid: {factory.get('phases', [])})")
    
    # Category
    category = fm.get("category")
    if category and category not in factory.get("categories", []):
        errors.append(f"invalid category '{category}' (valid: {factory.get('categories', [])})")
    
    # Presets
    presets = fm.get("presets", [])
    valid_presets = factory.get("presets", [])
    for preset in presets:
        if preset not in valid_presets:
            errors.append(f"invalid preset '{preset}' (valid: {valid_presets})")
    
    # Protocols
    pre_handoff = fm.get("pre_handoff", {})
    protocols = pre_handoff.get("protocols", [])
    valid_protocols = runtime.get("protocols", [])
    for protocol in protocols:
        if protocol not in valid_protocols:
            errors.append(f"invalid protocol '{protocol}'")
    
    # Checks
    checks = pre_handoff.get("checks", [])
    valid_checks = runtime.get("checks", [])
    for check in checks:
        if check not in valid_checks:
            errors.append(f"invalid check '{check}'")
    
    return errors


def validate_skill(skill_path: Path, factory: dict, runtime: dict, doc_types: dict) -> list:
    """Validate a single skill."""
    errors = []
    
    fm = extract_frontmatter(skill_path)
    if not fm:
        errors.append(f"{skill_path.name}: no valid frontmatter")
        return errors
    
    skill_name = fm.get("name", skill_path.name)
    
    # Identity validations
    errors.extend([f"{skill_name}: {e}" for e in validate_name(fm.get("name"))])
    errors.extend([f"{skill_name}: {e}" for e in validate_version(fm.get("version"))])
    errors.extend([f"{skill_name}: {e}" for e in validate_description(fm.get("description"))])
    
    # Enum validations
    errors.extend([f"{skill_name}: {e}" for e in validate_enums(fm, factory, runtime)])
    
    # Required fields
    if not fm.get("phase"):
        errors.append(f"{skill_name}: phase is required")
    if not fm.get("category"):
        errors.append(f"{skill_name}: category is required")
    if not fm.get("presets"):
        errors.append(f"{skill_name}: presets is required (at least one)")
    
    # Doc type validation
    creates = fm.get("creates", [])
    for create in creates:
        doc_type = create.get("doc_type")
        if doc_type and doc_types and doc_type not in doc_types.get("types", {}):
            errors.append(f"{skill_name}: unknown doc_type '{doc_type}'")
    
    return errors


def main():
    parser = argparse.ArgumentParser(description="Validate skills against Schema V3")
    parser.add_argument("--skill", help="Validate single skill")
    parser.add_argument("--check-docs", action="store_true", help="Validate doc types")
    args = parser.parse_args()
    
    # Load schema files
    schema_dir = Path("src/_meta/skills/schema")
    factory = load_yaml(schema_dir / "enums" / "factory.yaml")
    runtime = load_yaml(schema_dir / "enums" / "runtime.yaml")
    
    # Load doc types if checking docs
    doc_types = {}
    if args.check_docs:
        doc_types_path = Path("src/_meta/doc-types.yaml")
        if doc_types_path.exists():
            doc_types = load_yaml(doc_types_path)
    
    # Find skills
    skills_dir = Path("src/skills")
    private_dir = skills_dir / "private"
    
    if args.skill:
        skill_paths = [skills_dir / args.skill]
        if not skill_paths[0].exists():
            skill_paths = [private_dir / args.skill]
    else:
        skill_paths = [p for p in skills_dir.iterdir() if p.is_dir() and not p.name.startswith(".") and p.name != "private"]
        if private_dir.exists():
            skill_paths.extend([p for p in private_dir.iterdir() if p.is_dir()])
    
    # Validate
    all_errors = []
    validated = 0
    
    for skill_path in skill_paths:
        if not skill_path.exists():
            print(f"⚠️  Skill not found: {skill_path}")
            continue
        
        errors = validate_skill(skill_path, factory, runtime, doc_types)
        all_errors.extend(errors)
        validated += 1
    
    # Report
    if all_errors:
        print("=== VALIDATION ERRORS ===\n")
        for error in all_errors:
            print(f"❌ {error}")
        print(f"\n❌ {len(all_errors)} errors in {validated} skills")
        sys.exit(1)
    else:
        print(f"✅ {validated} skills validated successfully")
        sys.exit(0)


if __name__ == "__main__":
    main()
