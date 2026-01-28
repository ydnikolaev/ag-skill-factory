#!/usr/bin/env python3
"""
Generate JSON Schemas from all YAML schemas.

Usage:
    python3 scripts/generate_all_schemas.py          # Generate all
    python3 scripts/generate_all_schemas.py --type=X # Generate single (skills, documents, rules, workflows)

Output:
    src/_meta/schema/{type}/{type}-schema.json
"""

import argparse
import json
import yaml
from pathlib import Path


def load_yaml(path: Path) -> dict:
    """Load YAML file."""
    if not path.exists():
        return {}
    with open(path) as f:
        return yaml.safe_load(f) or {}


def generate_skills_schema(schema_dir: Path) -> dict:
    """Generate skills JSON Schema (existing logic)."""
    factory_enums = load_yaml(schema_dir / "enums" / "factory.yaml")
    runtime_enums = load_yaml(schema_dir / "enums" / "runtime.yaml")
    
    return {
        "$schema": "http://json-schema.org/draft-07/schema#",
        "$id": "https://antigravity.dev/skill-schema-v3.json",
        "title": "Antigravity Skill Schema V3",
        "description": "Schema for validating Antigravity skill definitions",
        "type": "object",
        "required": ["name", "description", "version", "phase", "category", "presets"],
        "properties": {
            "name": {"type": "string", "pattern": "^[a-z][a-z0-9-]*$"},
            "description": {"type": "string", "minLength": 10},
            "version": {"type": "string", "pattern": "^\\d+\\.\\d+\\.\\d+$"},
            "phase": {"type": "string", "enum": factory_enums.get("phases", [])},
            "category": {"type": "string", "enum": factory_enums.get("categories", [])},
            "presets": {"type": "array", "items": {"type": "string", "enum": factory_enums.get("presets", [])}, "minItems": 1},
            "pre_handoff": {
                "type": "object",
                "properties": {
                    "protocols": {"type": "array", "items": {"type": "string", "enum": runtime_enums.get("protocols", [])}},
                    "checks": {"type": "array", "items": {"type": "string", "enum": runtime_enums.get("checks", [])}}
                }
            }
        }
    }


def generate_documents_schema(schema_dir: Path) -> dict:
    """Generate documents JSON Schema."""
    enums = load_yaml(schema_dir / "enums.yaml")
    
    return {
        "$schema": "http://json-schema.org/draft-07/schema#",
        "$id": "https://antigravity.dev/document-schema-v1.json",
        "title": "Antigravity Document Schema V1",
        "description": "Schema for validating document template frontmatter",
        "type": "object",
        "required": ["status", "owner", "lifecycle", "work_unit"],
        "properties": {
            "status": {"type": "string", "enum": enums.get("statuses", [])},
            "owner": {"type": "string", "pattern": "^@[a-z][a-z0-9-]*$"},
            "lifecycle": {"type": "string", "enum": enums.get("lifecycles", [])},
            "work_unit": {"type": "string"},
            "created": {"type": "string"},
            "updated": {"type": "string"},
            "upstream": {
                "type": "array",
                "items": {
                    "type": "object",
                    "properties": {
                        "doc_type": {"type": "string"},
                        "owner": {"type": "string"}
                    }
                }
            },
            "downstream": {
                "type": "array",
                "items": {
                    "type": "object",
                    "properties": {
                        "skill": {"type": "string"},
                        "doc_type": {"type": "string"}
                    }
                }
            }
        }
    }


def generate_rules_schema(schema_dir: Path) -> dict:
    """Generate rules JSON Schema."""
    enums = load_yaml(schema_dir / "enums.yaml")
    
    return {
        "$schema": "http://json-schema.org/draft-07/schema#",
        "$id": "https://antigravity.dev/rule-schema-v1.json",
        "title": "Antigravity Rule Schema V1",
        "description": "Schema for validating rule frontmatter",
        "type": "object",
        "required": ["trigger"],
        "properties": {
            "trigger": {"type": "string", "enum": enums.get("triggers", [])},
            "description": {"type": "string"},
            "globs": {"type": "array", "items": {"type": "string"}}
        }
    }


def generate_workflows_schema(schema_dir: Path) -> dict:
    """Generate workflows JSON Schema."""
    return {
        "$schema": "http://json-schema.org/draft-07/schema#",
        "$id": "https://antigravity.dev/workflow-schema-v1.json",
        "title": "Antigravity Workflow Schema V1",
        "description": "Schema for validating workflow frontmatter",
        "type": "object",
        "required": ["description"],
        "properties": {
            "description": {"type": "string", "minLength": 5}
        }
    }


GENERATORS = {
    "skills": (generate_skills_schema, "skill-schema.json"),
    "documents": (generate_documents_schema, "document-schema.json"),
    "rules": (generate_rules_schema, "rule-schema.json"),
    "workflows": (generate_workflows_schema, "workflow-schema.json"),
}


def main():
    parser = argparse.ArgumentParser(description="Generate JSON Schemas")
    parser.add_argument("--type", help="Schema type to generate", choices=GENERATORS.keys())
    args = parser.parse_args()
    
    base_dir = Path("src/_meta/schema")
    
    types_to_generate = [args.type] if args.type else GENERATORS.keys()
    
    for schema_type in types_to_generate:
        generator, output_name = GENERATORS[schema_type]
        schema_dir = base_dir / schema_type
        output_path = schema_dir / output_name
        
        json_schema = generator(schema_dir)
        
        with open(output_path, "w") as f:
            json.dump(json_schema, f, indent=2)
        
        print(f"✅ Generated {output_path}")


if __name__ == "__main__":
    main()
