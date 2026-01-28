#!/usr/bin/env python3
"""
Validate all skills against Schema V3.

Dynamically parses skill-schema.yaml to validate frontmatter.

Usage:
    python3 scripts/validate_skills.py              # Validate all
    python3 scripts/validate_skills.py --skill=name # Validate single
    python3 scripts/validate_skills.py --strict     # Fail on warnings

Validations:
    - Required fields (from schema)
    - Enum values (from enums/*.yaml)
    - Field types (string, list, object)
    - Nested structures validation
"""

import argparse
import re
import sys
import yaml
from pathlib import Path
from typing import Any


def load_yaml(path: Path) -> dict:
    """Load YAML file."""
    if not path.exists():
        return {}
    with open(path) as f:
        return yaml.safe_load(f) or {}


def extract_frontmatter(skill_path: Path) -> dict | None:
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


class SchemaValidator:
    """Validates skill frontmatter against skill-schema.yaml."""
    
    # V3 required fields (top-level)
    REQUIRED_FIELDS = {
        "name", "description", "version", "phase", "category", "presets"
    }
    
    # V3 optional fields that should exist in migrated skills
    V3_FIELDS = {
        "scope", "tags",  # IDENTITY
        "mcp_servers", "allowed_tools", "dependencies", "context", "reads", "produces",  # CAPABILITIES
        "receives_from", "delegates_to", "return_paths",  # WORKFLOW
        "requires", "creates", "updates", "archives",  # DOCUMENTS
        "pre_handoff", "quality_gates",  # VALIDATION
        "required_sections"  # SECTIONS
    }
    
    def __init__(self, schema_dir: Path):
        self.schema = load_yaml(schema_dir / "skill-schema.yaml")
        self.factory = load_yaml(schema_dir / "enums" / "factory.yaml")
        self.runtime = load_yaml(schema_dir / "enums" / "runtime.yaml")
        self.errors = []
        self.warnings = []
    
    def _get_enum_keys(self, source: dict, key: str) -> list:
        """Get enum values from source (handles both list and dict formats)."""
        data = source.get(key, [])
        if isinstance(data, dict):
            return list(data.keys())
        return data if isinstance(data, list) else []
    
    def validate(self, fm: dict, skill_name: str) -> tuple[list, list]:
        """Validate frontmatter against schema. Returns (errors, warnings)."""
        self.errors = []
        self.warnings = []
        self.skill_name = skill_name
        
        # Check required fields
        for field in self.REQUIRED_FIELDS:
            if field not in fm or fm[field] is None:
                self.errors.append(f"{skill_name}: missing required field '{field}'")
        
        # Check V3 optional fields (warn if missing for migration tracking)
        for field in self.V3_FIELDS:
            if field not in fm:
                self.warnings.append(f"{skill_name}: V3 field '{field}' not present")
        
        # Validate specific fields
        self._validate_name(fm.get("name"))
        self._validate_version(fm.get("version"))
        self._validate_description(fm.get("description"))
        self._validate_enums(fm)
        self._validate_nested_structures(fm)
        
        return self.errors, self.warnings
    
    def _validate_name(self, name: str | None):
        """Validate skill name format."""
        if name and not re.match(r'^[a-z][a-z0-9-]*$', name):
            self.errors.append(f"{self.skill_name}: 'name' must be lowercase-with-hyphens")
    
    def _validate_version(self, version: str | None):
        """Validate version format."""
        if version and not re.match(r'^\d+\.\d+\.\d+$', version):
            self.errors.append(f"{self.skill_name}: 'version' must be semver (e.g. 3.0.0)")
    
    def _validate_description(self, desc: str | None):
        """Validate description."""
        if desc and len(desc) < 10:
            self.errors.append(f"{self.skill_name}: 'description' must be at least 10 characters")
    
    def _validate_enums(self, fm: dict):
        """Validate enum values against enums/*.yaml."""
        # Phase
        phase = fm.get("phase")
        valid_phases = self._get_enum_keys(self.factory, "phases")
        if phase and phase not in valid_phases:
            self.errors.append(f"{self.skill_name}: invalid phase '{phase}'")
        
        # Category
        category = fm.get("category")
        valid_categories = self._get_enum_keys(self.factory, "categories")
        if category and category not in valid_categories:
            self.errors.append(f"{self.skill_name}: invalid category '{category}'")
        
        # Scope
        scope = fm.get("scope")
        valid_scopes = self._get_enum_keys(self.runtime, "scopes")
        if scope and scope not in valid_scopes:
            self.errors.append(f"{self.skill_name}: invalid scope '{scope}'")
        
        # Presets
        presets = fm.get("presets", [])
        valid_presets = self._get_enum_keys(self.factory, "presets")
        if isinstance(presets, list):
            for preset in presets:
                if preset not in valid_presets:
                    self.errors.append(f"{self.skill_name}: invalid preset '{preset}'")
        
        # MCP servers
        mcp_servers = fm.get("mcp_servers", [])
        valid_mcp = self._get_enum_keys(self.runtime, "mcp_servers")
        if isinstance(mcp_servers, list):
            for server in mcp_servers:
                if valid_mcp and server not in valid_mcp:
                    self.warnings.append(f"{self.skill_name}: unknown mcp_server '{server}'")
        
        # Protocols (in pre_handoff)
        pre_handoff = fm.get("pre_handoff", {})
        if isinstance(pre_handoff, dict):
            protocols = pre_handoff.get("protocols", [])
            valid_protocols = self._get_enum_keys(self.runtime, "protocols")
            if isinstance(protocols, list):
                for protocol in protocols:
                    if protocol not in valid_protocols:
                        self.errors.append(f"{self.skill_name}: invalid protocol '{protocol}'")
            
            # Checks
            checks = pre_handoff.get("checks", [])
            valid_checks = self._get_enum_keys(self.runtime, "checks")
            if isinstance(checks, list):
                for check in checks:
                    if check not in valid_checks:
                        self.errors.append(f"{self.skill_name}: invalid check '{check}'")
        
        # Triggers in creates/updates/archives
        valid_triggers = self._get_enum_keys(self.runtime, "triggers")
        for section in ["creates", "updates", "archives"]:
            items = fm.get(section, [])
            if isinstance(items, list):
                for item in items:
                    if isinstance(item, dict):
                        trigger = item.get("trigger")
                        if trigger and trigger not in valid_triggers:
                            self.warnings.append(f"{self.skill_name}: unknown trigger '{trigger}' in {section}")
    
    def _validate_nested_structures(self, fm: dict):
        """Validate nested object structures."""
        # context must be object with required/optional
        context = fm.get("context")
        if context is not None:
            if not isinstance(context, dict):
                self.errors.append(f"{self.skill_name}: 'context' must be object with required/optional")
            else:
                for key in ["required", "optional"]:
                    items = context.get(key, [])
                    if items and not isinstance(items, list):
                        self.errors.append(f"{self.skill_name}: 'context.{key}' must be list")
                    elif isinstance(items, list):
                        for i, item in enumerate(items):
                            if not isinstance(item, dict):
                                self.errors.append(f"{self.skill_name}: 'context.{key}[{i}]' must be object")
                            elif "path" not in item:
                                self.warnings.append(f"{self.skill_name}: 'context.{key}[{i}]' missing 'path'")
        
        # pre_handoff must be object with protocols/checks
        pre_handoff = fm.get("pre_handoff")
        if pre_handoff is not None and not isinstance(pre_handoff, dict):
            self.errors.append(f"{self.skill_name}: 'pre_handoff' must be object with protocols/checks")
        
        # List fields type check
        list_fields = [
            "tags", "mcp_servers", "allowed_tools", "dependencies",
            "reads", "produces", "presets", "receives_from", "delegates_to",
            "return_paths", "requires", "creates", "updates", "archives",
            "quality_gates", "required_sections"
        ]
        for field in list_fields:
            value = fm.get(field)
            if value is not None and not isinstance(value, list):
                self.errors.append(f"{self.skill_name}: '{field}' must be list")


def validate_skill(skill_path: Path, validator: SchemaValidator) -> tuple[list, list]:
    """Validate a single skill. Returns (errors, warnings)."""
    fm = extract_frontmatter(skill_path)
    if not fm:
        return [f"{skill_path.name}: no valid frontmatter"], []
    
    skill_name = fm.get("name", skill_path.name)
    return validator.validate(fm, skill_name)


def main():
    parser = argparse.ArgumentParser(description="Validate skills against Schema V3")
    parser.add_argument("--skill", help="Validate single skill")
    parser.add_argument("--strict", action="store_true", help="Fail on warnings")
    parser.add_argument("--no-warnings", action="store_true", help="Hide warnings")
    args = parser.parse_args()
    
    # Initialize validator
    schema_dir = Path("src/_meta/schema/skills")
    validator = SchemaValidator(schema_dir)
    
    # Find skills
    skills_dir = Path("src/skills")
    private_dir = skills_dir / "private"
    
    if args.skill:
        skill_paths = [skills_dir / args.skill]
        if not skill_paths[0].exists():
            skill_paths = [private_dir / args.skill]
    else:
        skill_paths = [
            p for p in skills_dir.iterdir() 
            if p.is_dir() and not p.name.startswith(".") and p.name != "private"
        ]
        if private_dir.exists():
            skill_paths.extend([p for p in private_dir.iterdir() if p.is_dir()])
    
    # Validate
    all_errors = []
    all_warnings = []
    validated = 0
    
    for skill_path in sorted(skill_paths):
        if not skill_path.exists():
            print(f"⚠️  Skill not found: {skill_path}")
            continue
        
        errors, warnings = validate_skill(skill_path, validator)
        all_errors.extend(errors)
        all_warnings.extend(warnings)
        validated += 1
    
    # Report warnings
    if all_warnings and not args.no_warnings:
        print(f"=== WARNINGS ({len(all_warnings)}) ===\n")
        # Group by skill
        by_skill = {}
        for w in all_warnings:
            skill = w.split(":")[0]
            by_skill.setdefault(skill, []).append(w)
        
        for skill, warnings in sorted(by_skill.items()):
            v3_missing = [w for w in warnings if "V3 field" in w]
            other = [w for w in warnings if "V3 field" not in w]
            if v3_missing:
                fields = [w.split("'")[1] for w in v3_missing]
                print(f"⚠️  {skill}: missing V3 fields: {', '.join(fields)}")
            for w in other:
                print(f"⚠️  {w}")
        print()
    
    # Report errors
    if all_errors:
        print("=== VALIDATION ERRORS ===\n")
        for error in all_errors:
            print(f"❌ {error}")
        print(f"\n❌ {len(all_errors)} errors in {validated} skills")
        sys.exit(1)
    elif all_warnings and args.strict:
        print(f"\n⚠️  {len(all_warnings)} warnings (strict mode)")
        sys.exit(1)
    else:
        warn_msg = f" ({len(all_warnings)} warnings)" if all_warnings else ""
        print(f"✅ {validated} skills validated successfully{warn_msg}")
        sys.exit(0)


if __name__ == "__main__":
    main()
