#!/usr/bin/env python3
"""
Validate rules against Rule Schema.

Usage:
    python3 scripts/validate_rules.py           # Validate all
    python3 scripts/validate_rules.py --rule=X  # Validate single

Validations:
    - Frontmatter presence
    - Required fields (trigger)
    - Valid trigger enum
    - Naming convention (SCREAMING_SNAKE_CASE.md)
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


def extract_frontmatter(rule_path: Path) -> dict:
    """Extract YAML frontmatter from rule."""
    content = rule_path.read_text()
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


def validate_rule(rule_path: Path, enums: dict) -> list:
    """Validate a single rule."""
    errors = []
    rule_name = rule_path.name
    
    # Check naming convention (SCREAMING_SNAKE_CASE.md)
    if not re.match(r'^[A-Z][A-Z0-9_]*\.md$', rule_name):
        errors.append(f"{rule_name}: should be SCREAMING_SNAKE_CASE.md")
    
    fm = extract_frontmatter(rule_path)
    if not fm:
        errors.append(f"{rule_name}: no valid frontmatter")
        return errors
    
    # Required field: trigger
    trigger = fm.get("trigger")
    if not trigger:
        errors.append(f"{rule_name}: missing required field 'trigger'")
    elif trigger not in enums.get("triggers", []):
        errors.append(f"{rule_name}: invalid trigger '{trigger}'")
    
    # model_decision requires description
    if trigger == "model_decision" and not fm.get("description"):
        errors.append(f"{rule_name}: model_decision trigger requires 'description'")
    
    # glob trigger requires globs
    if trigger == "glob" and not fm.get("globs"):
        errors.append(f"{rule_name}: glob trigger requires 'globs' field")
    
    return errors


def main():
    parser = argparse.ArgumentParser(description="Validate rules")
    parser.add_argument("--rule", help="Validate single rule")
    args = parser.parse_args()
    
    # Load enums
    schema_dir = Path("src/_meta/schema/rules")
    enums = load_yaml(schema_dir / "enums" / "enums.yaml")
    
    # Find rules
    rules_dir = Path("src/rules")
    
    if args.rule:
        rule_paths = [rules_dir / args.rule]
    else:
        rule_paths = [p for p in rules_dir.glob("*.md")]
    
    # Validate
    all_errors = []
    validated = 0
    
    for rule_path in rule_paths:
        if not rule_path.exists():
            print(f"⚠️  Rule not found: {rule_path}")
            continue
        
        errors = validate_rule(rule_path, enums)
        all_errors.extend(errors)
        validated += 1
    
    # Report
    if all_errors:
        print("=== VALIDATION ERRORS ===\n")
        for error in all_errors:
            print(f"❌ {error}")
        print(f"\n❌ {len(all_errors)} errors in {validated} rules")
        sys.exit(1)
    else:
        print(f"✅ {validated} rules validated successfully")
        sys.exit(0)


if __name__ == "__main__":
    main()
