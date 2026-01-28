#!/usr/bin/env python3
"""
Validate document templates against Document Schema.

Usage:
    python3 scripts/validate_documents.py              # Validate all
    python3 scripts/validate_documents.py --doc=name   # Validate single

Validations:
    - Frontmatter presence
    - Required fields (status, owner, lifecycle, work_unit)
    - Valid enum values
    - Naming convention (_doc-type.md)
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


def extract_frontmatter(doc_path: Path) -> dict:
    """Extract YAML frontmatter from document."""
    content = doc_path.read_text()
    if not content.startswith("---"):
        return None
    
    lines = content.split('\n')
    for i, line in enumerate(lines[1:], 1):
        if line.strip() == "---":
            try:
                # Replace template placeholders with valid YAML strings
                yaml_content = '\n'.join(lines[1:i])
                yaml_content = yaml_content.replace("{WORK_UNIT}", "PLACEHOLDER_WORK_UNIT")
                yaml_content = yaml_content.replace("{DATE}", "PLACEHOLDER_DATE")
                yaml_content = yaml_content.replace("{", "").replace("}", "")
                # Quote @-prefixed values (skill names)
                yaml_content = re.sub(r': (@[\w-]+)', r': "\1"', yaml_content)
                return yaml.safe_load(yaml_content) or {}
            except Exception as e:
                return None
    return None


def validate_document(doc_path: Path, enums: dict) -> list:
    """Validate a single document template."""
    errors = []
    doc_name = doc_path.name
    
    # Check naming convention
    if not doc_name.startswith("_"):
        errors.append(f"{doc_name}: should start with underscore")
    
    fm = extract_frontmatter(doc_path)
    if not fm:
        errors.append(f"{doc_name}: no valid frontmatter")
        return errors
    
    # Required fields (work_unit optional for 'living' lifecycle)
    lifecycle = fm.get("lifecycle")
    required = ["status", "owner", "lifecycle"]
    if lifecycle != "living":
        required.append("work_unit")
    
    for field in required:
        if field not in fm:
            errors.append(f"{doc_name}: missing required field '{field}'")
    
    # Validate status enum
    status = fm.get("status")
    if status and status not in enums.get("statuses", []):
        errors.append(f"{doc_name}: invalid status '{status}'")
    
    # Validate lifecycle enum
    lifecycle = fm.get("lifecycle")
    if lifecycle and lifecycle not in enums.get("lifecycles", []):
        errors.append(f"{doc_name}: invalid lifecycle '{lifecycle}'")
    
    # Validate owner format
    owner = fm.get("owner", "")
    if owner and not owner.startswith("@"):
        errors.append(f"{doc_name}: owner should start with @ (got '{owner}')")
    
    return errors


def main():
    parser = argparse.ArgumentParser(description="Validate document templates")
    parser.add_argument("--doc", help="Validate single document")
    args = parser.parse_args()
    
    # Load enums from shared
    shared_dir = Path("src/_meta/schema/shared")
    enums = load_yaml(shared_dir / "runtime.yaml")
    
    # Find documents
    docs_dir = Path("src/templates/documents")
    
    if args.doc:
        doc_paths = [docs_dir / args.doc]
    else:
        doc_paths = [p for p in docs_dir.glob("*.md")]
    
    # Validate
    all_errors = []
    validated = 0
    
    for doc_path in doc_paths:
        if not doc_path.exists():
            print(f"⚠️  Document not found: {doc_path}")
            continue
        
        errors = validate_document(doc_path, enums)
        all_errors.extend(errors)
        validated += 1
    
    # Report
    if all_errors:
        print("=== VALIDATION ERRORS ===\n")
        for error in all_errors:
            print(f"❌ {error}")
        print(f"\n❌ {len(all_errors)} errors in {validated} documents")
        sys.exit(1)
    else:
        print(f"✅ {validated} documents validated successfully")
        sys.exit(0)


if __name__ == "__main__":
    main()
