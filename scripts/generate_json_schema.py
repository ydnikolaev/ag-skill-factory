#!/usr/bin/env python3
"""
Generate JSON Schema from skill-schema.yaml.

Usage:
    python3 scripts/generate_json_schema.py

Output:
    src/_meta/skills/schema/skill-schema.json
"""

import json
import yaml
from pathlib import Path


def load_yaml(path: Path) -> dict:
    """Load YAML file."""
    with open(path) as f:
        return yaml.safe_load(f) or {}


def yaml_type_to_json_schema(yaml_type: str) -> dict:
    """Convert YAML type to JSON Schema type."""
    if yaml_type == "string":
        return {"type": "string"}
    elif yaml_type == "enum":
        return {"type": "string"}  # Will be populated with enum values
    elif yaml_type.startswith("list["):
        inner = yaml_type[5:-1]
        if inner == "string":
            return {"type": "array", "items": {"type": "string"}}
        elif inner == "object":
            return {"type": "array", "items": {"type": "object"}}
        else:
            return {"type": "array"}
    elif yaml_type == "object":
        return {"type": "object"}
    else:
        return {"type": "string"}


def generate_json_schema(schema_dir: Path) -> dict:
    """Generate JSON Schema from YAML schema."""
    schema = load_yaml(schema_dir / "skill-schema.yaml")
    factory_enums = load_yaml(schema_dir / "enums" / "factory.yaml")
    runtime_enums = load_yaml(schema_dir / "enums" / "runtime.yaml")
    
    json_schema = {
        "$schema": "http://json-schema.org/draft-07/schema#",
        "$id": "https://antigravity.dev/skill-schema-v3.json",
        "title": "Antigravity Skill Schema V3",
        "description": "Schema for validating Antigravity skill definitions",
        "type": "object",
        "required": ["name", "description", "version", "phase", "category", "presets"],
        "properties": {
            # Identity
            "name": {
                "type": "string",
                "pattern": "^[a-z][a-z0-9-]*$",
                "description": "Skill name (lowercase-with-hyphens)"
            },
            "description": {
                "type": "string",
                "minLength": 10,
                "description": "WHAT it does and WHEN to use it"
            },
            "version": {
                "type": "string",
                "pattern": "^\\d+\\.\\d+\\.\\d+$",
                "description": "Semantic version"
            },
            "phase": {
                "type": "string",
                "enum": factory_enums.get("phases", []),
                "description": "Pipeline phase"
            },
            "category": {
                "type": "string",
                "enum": factory_enums.get("categories", []),
                "description": "Output category"
            },
            "scope": {
                "type": "string",
                "enum": ["project", "global"],
                "default": "project"
            },
            "tags": {
                "type": "array",
                "items": {"type": "string"}
            },
            
            # Workflow
            "presets": {
                "type": "array",
                "items": {
                    "type": "string",
                    "enum": factory_enums.get("presets", [])
                },
                "minItems": 1
            },
            "receives_from": {
                "type": "array",
                "items": {
                    "type": "object",
                    "properties": {
                        "skill": {"type": "string"},
                        "docs": {"type": "array", "items": {"type": "string"}}
                    },
                    "required": ["skill"]
                }
            },
            "delegates_to": {
                "type": "array",
                "items": {
                    "type": "object",
                    "properties": {
                        "skill": {"type": "string"},
                        "docs": {"type": "array", "items": {"type": "string"}}
                    },
                    "required": ["skill"]
                }
            },
            "return_paths": {
                "type": "array",
                "items": {
                    "type": "object",
                    "properties": {
                        "skill": {"type": "string"},
                        "docs": {"type": "array", "items": {"type": "string"}}
                    }
                }
            },
            
            # Documents
            "creates": {
                "type": "array",
                "items": {
                    "type": "object",
                    "properties": {
                        "doc_type": {"type": "string"},
                        "path": {"type": "string"},
                        "lifecycle": {
                            "type": "string",
                            "enum": runtime_enums.get("lifecycles", [])
                        },
                        "initial_status": {"type": "string"},
                        "trigger": {"type": "string"}
                    },
                    "required": ["doc_type", "path"]
                }
            },
            "requires": {
                "type": "array",
                "items": {
                    "type": "object",
                    "properties": {
                        "doc_type": {"type": "string"},
                        "status": {"type": "string"}
                    }
                }
            },
            "updates": {
                "type": "array",
                "items": {
                    "type": "object",
                    "properties": {
                        "doc_type": {"type": "string"},
                        "new_status": {"type": "string"}
                    }
                }
            },
            "archives": {
                "type": "array",
                "items": {
                    "type": "object",
                    "properties": {
                        "doc_type": {"type": "string"},
                        "destination": {"type": "string"}
                    }
                }
            },
            
            # Capabilities
            "mcp_servers": {
                "type": "array",
                "items": {"type": "string"}
            },
            "dependencies": {
                "type": "array",
                "items": {"type": "string"}
            },
            "context": {
                "type": "object",
                "properties": {
                    "required": {"type": "array"},
                    "optional": {"type": "array"}
                }
            },
            "reads": {
                "type": "array",
                "items": {
                    "type": "object",
                    "properties": {
                        "type": {"type": "string"},
                        "from": {"type": "string"}
                    }
                }
            },
            "produces": {
                "type": "array",
                "items": {
                    "type": "object",
                    "properties": {
                        "type": {"type": "string"}
                    }
                }
            },
            
            # Validation
            "pre_handoff": {
                "type": "object",
                "properties": {
                    "protocols": {
                        "type": "array",
                        "items": {
                            "type": "string",
                            "enum": runtime_enums.get("protocols", [])
                        }
                    },
                    "checks": {
                        "type": "array",
                        "items": {
                            "type": "string",
                            "enum": runtime_enums.get("checks", [])
                        }
                    }
                }
            },
            "quality_gates": {
                "type": "array",
                "items": {
                    "type": "object",
                    "properties": {
                        "type": {"type": "string"},
                        "value": {}
                    }
                }
            },
            
            # Extensions
            "extensions": {
                "type": "object",
                "properties": {
                    "custom_triggers": {"type": "array", "items": {"type": "string"}},
                    "custom_statuses": {"type": "array", "items": {"type": "string"}},
                    "custom_checks": {"type": "array", "items": {"type": "string"}}
                }
            }
        }
    }
    
    return json_schema


def main():
    schema_dir = Path("src/_meta/skills/schema")
    output_path = schema_dir / "skill-schema.json"
    
    json_schema = generate_json_schema(schema_dir)
    
    with open(output_path, "w") as f:
        json.dump(json_schema, f, indent=2)
    
    print(f"✅ Generated {output_path}")


if __name__ == "__main__":
    main()
