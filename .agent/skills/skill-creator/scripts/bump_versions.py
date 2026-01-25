#!/usr/bin/env python3
"""
Auto-bump skill versions based on staged or changed files.

Usage:
    python3 bump_versions.py              # Bump based on staged files
    python3 bump_versions.py --all        # Bump all changed skills (vs last commit)
    python3 bump_versions.py --major      # Force major bump
    python3 bump_versions.py --minor      # Force minor bump
"""
import os
import sys
import re
import subprocess


def get_staged_skills() -> list:
    """Get list of staged SKILL.md files."""
    result = subprocess.run(
        ["git", "diff", "--cached", "--name-only"],
        capture_output=True,
        text=True
    )
    files = result.stdout.strip().split('\n')
    skills = []
    for f in files:
        if f and 'SKILL.md' in f and ('blueprint/skills/' in f or 'blueprint/private/' in f):
            skills.append(f)
    return skills


def get_changed_skills() -> list:
    """Get list of changed SKILL.md files vs last commit."""
    result = subprocess.run(
        ["git", "diff", "--name-only", "HEAD"],
        capture_output=True,
        text=True
    )
    files = result.stdout.strip().split('\n')
    skills = []
    for f in files:
        if f and 'SKILL.md' in f and ('blueprint/skills/' in f or 'blueprint/private/' in f):
            skills.append(f)
    return skills


def parse_version(content: str) -> tuple:
    """Extract version from frontmatter."""
    match = re.search(r'^version:\s*(\d+)\.(\d+)\.(\d+)', content, re.MULTILINE)
    if match:
        return int(match.group(1)), int(match.group(2)), int(match.group(3))
    return 1, 0, 0


def bump_version(major: int, minor: int, patch: int, bump_type: str) -> str:
    """Bump version according to type."""
    if bump_type == 'major':
        return f"{major + 1}.0.0"
    elif bump_type == 'minor':
        return f"{major}.{minor + 1}.0"
    else:  # patch
        return f"{major}.{minor}.{patch + 1}"


def detect_bump_type(skill_path: str) -> str:
    """Detect bump type from diff content."""
    result = subprocess.run(
        ["git", "diff", "--cached", skill_path],
        capture_output=True,
        text=True
    )
    diff = result.stdout
    
    # Check for new sections (minor bump)
    if re.search(r'^\+## ', diff, re.MULTILINE):
        return 'minor'
    
    # Check for breaking changes marker
    if 'BREAKING:' in diff or 'BREAKING CHANGE' in diff:
        return 'major'
    
    # Default to patch
    return 'patch'


def update_skill_version(skill_path: str, bump_type: str) -> tuple:
    """Update version in skill file. Returns (old_version, new_version)."""
    with open(skill_path, 'r') as f:
        content = f.read()
    
    major, minor, patch = parse_version(content)
    old_version = f"{major}.{minor}.{patch}"
    new_version = bump_version(major, minor, patch, bump_type)
    
    # Replace version in frontmatter
    new_content = re.sub(
        r'^version:\s*\d+\.\d+\.\d+',
        f'version: {new_version}',
        content,
        count=1,
        flags=re.MULTILINE
    )
    
    with open(skill_path, 'w') as f:
        f.write(new_content)
    
    return old_version, new_version


def stage_file(skill_path: str):
    """Stage the updated file."""
    subprocess.run(["git", "add", skill_path], check=True)


def main():
    force_type = None
    use_all = False
    
    for arg in sys.argv[1:]:
        if arg == '--major':
            force_type = 'major'
        elif arg == '--minor':
            force_type = 'minor'
        elif arg == '--patch':
            force_type = 'patch'
        elif arg == '--all':
            use_all = True
    
    if use_all:
        skills = get_changed_skills()
    else:
        skills = get_staged_skills()
    
    if not skills:
        print("📦 No skill changes to version")
        return
    
    print(f"🔍 Found {len(skills)} skill(s) to version...")
    
    for skill_path in skills:
        skill_name = os.path.basename(os.path.dirname(skill_path))
        
        if force_type:
            bump_type = force_type
        else:
            bump_type = detect_bump_type(skill_path)
        
        old_ver, new_ver = update_skill_version(skill_path, bump_type)
        stage_file(skill_path)
        
        print(f"   ✅ {skill_name}: {old_ver} → {new_ver} ({bump_type})")
    
    print(f"\n🎉 Bumped {len(skills)} skill version(s)")


if __name__ == "__main__":
    main()
