#!/usr/bin/env python3
"""
Generate rules-matrix.yaml from src/rules/ frontmatter.

Usage:
    python3 scripts/generate_rules.py

Output:
    src/_meta/rules-matrix.yaml
"""

import yaml
import re
from pathlib import Path


def parse_frontmatter(content: str) -> dict:
    """Extract YAML frontmatter from markdown file."""
    match = re.match(r'^---\n(.*?)\n---', content, re.DOTALL)
    if match:
        try:
            return yaml.safe_load(match.group(1)) or {}
        except yaml.YAMLError:
            return {}
    return {}


def extract_first_line(content: str) -> str:
    """Extract first non-empty line after frontmatter for description fallback."""
    # Remove frontmatter
    content = re.sub(r'^---\n.*?\n---\n*', '', content, flags=re.DOTALL)
    # Find first blockquote or heading
    match = re.search(r'^>\s*(.+)$', content, re.MULTILINE)
    if match:
        return match.group(1).strip()
    return ""


def main():
    root = Path(__file__).parent.parent
    src_rules = root / "src" / "rules"
    output_file = root / "src" / "_meta" / "rules-matrix.yaml"
    
    if not src_rules.exists():
        print("❌ src/rules not found")
        return
    
    rules = []
    
    for rule_file in sorted(src_rules.glob("*.md")):
        content = rule_file.read_text()
        frontmatter = parse_frontmatter(content)
        
        rule_name = rule_file.stem  # BRAIN_TO_DOCS, GIT_PROTOCOL, etc.
        
        rule_data = {
            "name": rule_name,
            "file": f"rules/{rule_file.name}",
            "trigger": frontmatter.get("trigger", "model_decision"),
            "description": frontmatter.get("description", extract_first_line(content)),
        }
        
        # Optional fields
        if "glob" in frontmatter:
            rule_data["glob"] = frontmatter["glob"]
        if "skills" in frontmatter:
            rule_data["skills"] = frontmatter["skills"]
        
        rules.append(rule_data)
        print(f"  ✅ {rule_name}")
    
    # Build output structure
    output = {
        "# AUTO-GENERATED": "Do not edit manually. Run: python3 scripts/generate_rules.py",
        "rules": rules,
        "triggers": {
            "always_on": "Applied to every conversation",
            "model_decision": "Model decides based on description",
            "manual": "Only via @mention",
            "glob": "When working with matching files"
        }
    }
    
    output_file.parent.mkdir(parents=True, exist_ok=True)
    
    with open(output_file, 'w') as f:
        # Write header comment
        f.write("# Rules Matrix - Auto-generated from src/rules/\n")
        f.write("# Run: python3 scripts/generate_rules.py\n\n")
        yaml.dump(output, f, default_flow_style=False, allow_unicode=True, sort_keys=False)
    
    print(f"\n📁 Generated rules-matrix.yaml ({len(rules)} rules)")


if __name__ == "__main__":
    print("📝 Generating rules matrix...")
    main()
