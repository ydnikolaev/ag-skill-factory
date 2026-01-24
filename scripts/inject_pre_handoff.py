#!/usr/bin/env python3
"""Inject Pre-Handoff Validation section into all skills."""

import os
import re

SQUADS_DIR = "/Users/yuranikolaev/Developer/antigravity/ag-skill-factory/squads"

PRE_HANDOFF_SECTION = '''## Pre-Handoff Validation (Hard Stop)

> [!CAUTION]
> **MANDATORY self-check before `notify_user` or delegation.**

| # | Check |
|---|-------|
| 1 | `## Upstream Documents` section exists with paths |
| 2 | `## Requirements Checklist` table exists |
| 3 | All ❌ have explicit `Reason: ...` |
| 4 | Document in `review/` folder |
| 5 | `ARTIFACT_REGISTRY.md` updated |

**If ANY unchecked → DO NOT PROCEED.**

'''

def inject_pre_handoff(skill_path):
    """Inject Pre-Handoff Validation before Handoff Protocol."""
    with open(skill_path, 'r') as f:
        content = f.read()
    
    # Skip if already has Pre-Handoff Validation
    if 'Pre-Handoff Validation' in content:
        return False
    
    # Find ## Handoff Protocol and insert before it
    pattern = r'\n## Handoff Protocol'
    if re.search(pattern, content):
        new_content = re.sub(pattern, '\n' + PRE_HANDOFF_SECTION + '## Handoff Protocol', content)
        with open(skill_path, 'w') as f:
            f.write(new_content)
        return True
    return False

def main():
    updated = []
    skipped = []
    
    for skill_name in os.listdir(SQUADS_DIR):
        skill_dir = os.path.join(SQUADS_DIR, skill_name)
        if not os.path.isdir(skill_dir):
            continue
        if skill_name.startswith('_') or skill_name == 'references':
            continue
            
        skill_file = os.path.join(skill_dir, 'SKILL.md')
        if os.path.exists(skill_file):
            if inject_pre_handoff(skill_file):
                updated.append(skill_name)
            else:
                skipped.append(skill_name)
    
    print(f"✅ Updated: {len(updated)} skills")
    for s in updated:
        print(f"   - {s}")
    if skipped:
        print(f"⏭️ Skipped: {len(skipped)} (already have Pre-Handoff)")
        for s in skipped:
            print(f"   - {s}")

if __name__ == '__main__':
    main()
