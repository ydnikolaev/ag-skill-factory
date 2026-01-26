#!/usr/bin/env python3
"""
Generate skill matrix and pipelines from SKILL.md frontmatter.

Usage:
    python3 scripts/generate_pipelines.py
    
Output:
    blueprint/_meta/_skills/skill-matrix.yaml - Full skill matrix
    blueprint/_meta/_pipelines/PIPELINE_<preset>.md - Pipeline per preset
"""

import os
import yaml
from pathlib import Path


def extract_frontmatter(skill_path: Path) -> dict:
    """Extract YAML frontmatter from SKILL.md."""
    skill_md = skill_path / "SKILL.md"
    if not skill_md.exists():
        return {}
    
    content = skill_md.read_text()
    if not content.startswith("---"):
        return {}
    
    end = content.find("\n---\n", 4)
    if end == -1:
        return {}
    
    frontmatter_text = content[4:end]
    try:
        return yaml.safe_load(frontmatter_text) or {}
    except yaml.YAMLError:
        return {}


def build_skill_matrix(blueprint_skills: Path) -> dict:
    """Build complete skill matrix from all SKILL.md files."""
    matrix = {
        "skills": {},
        "phases": {},
        "handoffs": [],
    }
    
    for skill_path in sorted(blueprint_skills.iterdir()):
        if not skill_path.is_dir() or skill_path.name.startswith("."):
            continue
        
        fm = extract_frontmatter(skill_path)
        if not fm:
            continue
        
        skill_name = skill_path.name
        matrix["skills"][skill_name] = {
            "name": fm.get("name", skill_name),
            "description": fm.get("description", ""),
            "version": fm.get("version", "1.0.0"),
            "phase": fm.get("phase", "utility"),
            "category": fm.get("category", "utility"),
            "receives_from": fm.get("receives_from", []),
            "delegates_to": fm.get("delegates_to", []),
            "outputs": fm.get("outputs", []),
        }
        
        # Build phase groupings
        phase = fm.get("phase", "utility")
        if phase not in matrix["phases"]:
            matrix["phases"][phase] = []
        matrix["phases"][phase].append(skill_name)
        
        # Build handoffs
        for target in fm.get("delegates_to", []):
            for output in fm.get("outputs", []):
                matrix["handoffs"].append({
                    "from": skill_name,
                    "to": target,
                    "doc_type": output.get("doc_type") or output.get("artifact", "").rsplit(".", 1)[0],
                    "path": output.get("path", ""),
                })
    
    return matrix


def generate_doc_types(matrix: dict, output_path: Path):
    """Generate doc-types.yaml from skill matrix outputs."""
    doc_types = {"types": {}}
    
    for skill_name, skill_data in matrix["skills"].items():
        phase = skill_data.get("phase", "utility")
        
        for output in skill_data.get("outputs", []):
            # Support both doc_type (new) and artifact (legacy)
            doc_type = output.get("doc_type") or output.get("artifact", "")
            if not doc_type:
                continue
            
            # Remove .md/.yaml extension if present (legacy artifact format)
            if "." in doc_type:
                doc_type = doc_type.rsplit(".", 1)[0]
            
            # Get consumers from delegates_to
            consumers = skill_data.get("delegates_to", [])
            
            # Get category from doc_category or derive from path
            category = output.get("doc_category", "")
            if not category:
                path = output.get("path", "")
                # Extract category from path like project/docs/active/specs/
                parts = path.rstrip("/").split("/")
                category = parts[-1] if parts else "other"
            
            doc_types["types"][doc_type] = {
                "phase": phase,
                "creator": skill_name,
                "consumers": consumers,
                "category": category,
                "doc_type": doc_type,
                "template": f"_{doc_type}.md",
                "lifecycle": output.get("lifecycle", "per-feature"),
            }
    
    output_path.parent.mkdir(parents=True, exist_ok=True)
    output_path.parent.mkdir(parents=True, exist_ok=True)
    with open(output_path, "w") as f:
        f.write("# Auto-generated from skill frontmatter\n")
        f.write(f"# Run: python3 scripts/generate_pipelines.py\n\n")
        yaml.dump(doc_types, f, default_flow_style=False, allow_unicode=True, sort_keys=False)
    
    return len(doc_types["types"])


def resolve_preset_skills(preset_name: str, presets: dict, all_skills: set) -> list:
    """Resolve full skill list for a preset."""
    preset = presets.get(preset_name, {})
    skills = set()
    
    if preset.get("skills") == "*":
        return sorted(all_skills)
    
    extends = preset.get("extends", [])
    if isinstance(extends, str):
        extends = [extends]
    
    for parent in extends:
        parent_skills = resolve_preset_skills(parent, presets, all_skills)
        skills.update(parent_skills)
    
    own_skills = preset.get("skills", [])
    if isinstance(own_skills, list):
        skills.update(own_skills)
    
    return sorted(skills)


def generate_pipeline_file(preset_name: str, skill_names: list, matrix: dict, 
                           output_dir: Path, preset_desc: str):
    """Generate PIPELINE_<preset>.md file."""
    output_file = output_dir / f"PIPELINE_{preset_name}.md"
    
    # Filter matrix to only preset skills
    preset_skills = {k: v for k, v in matrix["skills"].items() if k in skill_names}
    
    # Build phase summary
    phase_order = ["discovery", "definition", "design", "architecture", "implementation", "delivery", "utility"]
    phases_in_preset = {}
    for skill_name, skill_data in preset_skills.items():
        phase = skill_data.get("phase", "utility")
        if phase not in phases_in_preset:
            phases_in_preset[phase] = []
        phases_in_preset[phase].append(skill_name)
    
    # Build handoffs in preset
    preset_handoffs = [
        h for h in matrix["handoffs"]
        if h["from"] in skill_names and h["to"] in skill_names
    ]
    
    # Build return paths (reverse of handoffs for bug scenarios)
    return_paths = []
    for h in preset_handoffs:
        # qa-lead → backend/frontend is a common return path
        if h["from"] in ["backend-go-expert", "frontend-nuxt"] and h["to"] == "qa-lead":
            return_paths.append({
                "from": "qa-lead",
                "to": h["from"],
                "trigger": f"Bugs found in {h['from'].replace('-', ' ')}"
            })
    
    lines = [
        "---",
        "trigger: model_decision",
        f"description: Pipeline for {preset_name} preset. Skill handoffs and phases.",
        "---",
        "",
        f"# Pipeline ({preset_name})",
        "",
        f"> {preset_desc}",
        "",
        "## Phases",
        "",
        "| Phase | Skills | Outputs |",
        "|-------|--------|---------|",
    ]
    
    for phase in phase_order:
        if phase not in phases_in_preset:
            continue
        skills = phases_in_preset[phase]
        skill_refs = ", ".join([f"`@{s}`" for s in skills])
        outputs = []
        for s in skills:
            for o in preset_skills.get(s, {}).get("outputs", []):
                dt = o.get("doc_type") or o.get("artifact", "")
                if "." in dt:
                    dt = dt.rsplit(".", 1)[0]
                outputs.append(dt)
        outputs_str = ", ".join(outputs) if outputs else "—"
        lines.append(f"| {phase.title()} | {skill_refs} | {outputs_str} |")
    
    if preset_handoffs:
        lines.extend([
            "",
            "## Handoff Matrix",
            "",
            "| From | To | Artifact |",
            "|------|-----|----------|",
        ])
        
        for h in preset_handoffs:
            lines.append(f"| `@{h['from']}` | `@{h['to']}` | {h['doc_type']} |")
    
    if return_paths:
        lines.extend([
            "",
            "## Return Paths",
            "",
            "| From | To | Trigger |",
            "|------|-----|---------|",
        ])
        for rp in return_paths:
            lines.append(f"| `@{rp['from']}` | `@{rp['to']}` | {rp['trigger']} |")
    
    lines.append("")
    
    output_file.write_text("\n".join(lines))
    return len(preset_handoffs)


def main():
    root = Path(__file__).parent.parent
    src_skills = root / "src" / "skills"
    presets_file = root / "src" / "_meta" / "presets.yaml"
    matrix_output = root / "src" / "_meta" / "skill-matrix.yaml"
    doc_types_output = root / "src" / "_meta" / "doc-types.yaml"
    pipelines_output = root / "src" / "_meta" / "pipelines"
    
    if not src_skills.exists():
        print("❌ src/skills not found")
        return
    
    # Build skill matrix
    print("📊 Building skill matrix...")
    matrix = build_skill_matrix(src_skills)
    
    # Save matrix
    matrix_output.parent.mkdir(parents=True, exist_ok=True)
    # Save matrix
    matrix_output.parent.mkdir(parents=True, exist_ok=True)
    with open(matrix_output, "w") as f:
        f.write("# Auto-generated from skill frontmatter\n")
        f.write(f"# Run: python3 scripts/generate_pipelines.py\n\n")
        yaml.dump(matrix, f, default_flow_style=False, allow_unicode=True, sort_keys=False)
    print(f"  ✅ skill-matrix.yaml ({len(matrix['skills'])} skills, {len(matrix['handoffs'])} handoffs)")
    
    # Generate doc-types.yaml
    doc_count = generate_doc_types(matrix, doc_types_output)
    print(f"  ✅ doc-types.yaml ({doc_count} document types)")
    
    # Load presets
    if not presets_file.exists():
        print("❌ src/_meta/presets.yaml not found")
        return
    
    with open(presets_file) as f:
        presets = yaml.safe_load(f)
    
    # Generate pipelines
    print("📝 Generating pipeline files...")
    pipelines_output.mkdir(parents=True, exist_ok=True)
    all_skills = set(matrix["skills"].keys())
    
    for preset_name, preset_config in presets.items():
        if preset_name.startswith("_"):
            continue

        skill_names = resolve_preset_skills(preset_name, presets, all_skills)
        preset_desc = preset_config.get("description", f"{preset_name} preset")
        count = generate_pipeline_file(preset_name, skill_names, matrix, pipelines_output, preset_desc)
        print(f"  ✅ PIPELINE_{preset_name}.md ({len(skill_names)} skills, {count} handoffs)")
    
    print(f"\n📁 Generated {len(presets)} pipeline files in src/_meta/pipelines/")


if __name__ == "__main__":
    main()
