#!/usr/bin/env python3
"""
Validate workflows against Workflow Schema.

Usage:
    python3 scripts/validate_workflows.py              # Validate all
    python3 scripts/validate_workflows.py --workflow=X # Validate single

Validations:
    - Frontmatter presence
    - Required fields (description)
    - Naming convention (lowercase-with-hyphens.md)
    - Steps section presence
"""

import argparse
import re
import sys
import yaml
from pathlib import Path


def extract_frontmatter(workflow_path: Path) -> dict:
    """Extract YAML frontmatter from workflow."""
    content = workflow_path.read_text()
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


def validate_workflow(workflow_path: Path) -> list:
    """Validate a single workflow."""
    errors = []
    workflow_name = workflow_path.name
    
    # Check naming convention (lowercase-with-hyphens.md)
    if not re.match(r'^[a-z][a-z0-9-]*\.md$', workflow_name):
        errors.append(f"{workflow_name}: should be lowercase-with-hyphens.md")
    
    fm = extract_frontmatter(workflow_path)
    if not fm:
        errors.append(f"{workflow_name}: no valid frontmatter")
        return errors
    
    # Required field: description
    description = fm.get("description")
    if not description:
        errors.append(f"{workflow_name}: missing required field 'description'")
    elif len(description) < 5:
        errors.append(f"{workflow_name}: description too short (min 5 chars)")
    
    # Check for Steps section
    content = workflow_path.read_text()
    if "## Steps" not in content:
        errors.append(f"{workflow_name}: missing '## Steps' section")
    
    return errors


def main():
    parser = argparse.ArgumentParser(description="Validate workflows")
    parser.add_argument("--workflow", help="Validate single workflow")
    args = parser.parse_args()
    
    # Find workflows
    workflows_dir = Path("src/workflows")
    
    if args.workflow:
        workflow_paths = [workflows_dir / args.workflow]
    else:
        workflow_paths = [p for p in workflows_dir.glob("*.md")]
    
    # Validate
    all_errors = []
    validated = 0
    
    for workflow_path in workflow_paths:
        if not workflow_path.exists():
            print(f"⚠️  Workflow not found: {workflow_path}")
            continue
        
        errors = validate_workflow(workflow_path)
        all_errors.extend(errors)
        validated += 1
    
    # Report
    if all_errors:
        print("=== VALIDATION ERRORS ===\n")
        for error in all_errors:
            print(f"❌ {error}")
        print(f"\n❌ {len(all_errors)} errors in {validated} workflows")
        sys.exit(1)
    else:
        print(f"✅ {validated} workflows validated successfully")
        sys.exit(0)


if __name__ == "__main__":
    main()
