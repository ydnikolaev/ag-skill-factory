import os
import re

# Configuration
SKILLS_DIR = "blueprint/skills"
META_SECTIONS_DIR = "blueprint/_meta/_skills/sections"

# Mapping: Section Header -> Shared File Name
SECTION_MAPPING = {
    "Language Requirements": "language-requirements.md",
    "Team Collaboration": "team-collaboration.md",
    "Pre-Handoff Validation (Hard Stop)": "pre-handoff-validation.md",
    "Pre-Handoff Validation": "pre-handoff-validation.md",
    "Handoff Protocol": "handoff-protocol.md",
    "Git Protocol": "git-protocol.md",
    "Git Protocol (Hard Stop)": "git-protocol.md",
    "Tech Debt Protocol": "tech-debt-protocol.md",
    "Tech Debt Protocol (Hard Stop)": "tech-debt-protocol.md",
    "Document Lifecycle": "document-structure-protocol.md",
    "Resources": "resources.md"
}

# Sections that are purely static and should NOT have specific content appended
# (If a skill has content, it might be replaced by the static shared one, or we should verify)
STATIC_SECTIONS = {
    "Language Requirements", 
    "Tech Debt Protocol",
    "Tech Debt Protocol (Hard Stop)"
}

def read_file(path):
    if not os.path.exists(path):
        return None
    with open(path, "r", encoding="utf-8") as f:
        return f.read()

def write_file(path, content):
    with open(path, "w", encoding="utf-8") as f:
        f.write(content)

def extract_section(content, header):
    """
    Extracts section content (excluding header) until the next generic header.
    Returns (content, start_index, end_index) or None.
    """
    # Regex to find the header (at start of line)
    # We look for ## HeaderName
    pattern = r'(^|\n)## ' + re.escape(header) + r'\s*\n'
    match = re.search(pattern, content)
    
    if not match:
        return None
    
    start_idx = match.end() # Start of body
    header_start_idx = match.start()
    if content[header_start_idx] == '\n':
        header_start_idx += 1 # Skip the leading newline if matched
    
    # Find next header ## ...
    next_header = re.search(r'(^|\n)## ', content[start_idx:])
    
    if next_header:
        end_idx = start_idx + next_header.start()
    else:
        end_idx = len(content)
        
    body = content[start_idx:end_idx].strip()
    return body, header_start_idx, end_idx

def append_to_shared_file(shared_filename, skill_name, body):
    path = os.path.join(META_SECTIONS_DIR, shared_filename)
    content = read_file(path)
    if not content:
        print(f"Error: Shared file not found: {path}")
        return
    
    # Check if skill block already exists
    if f"<!-- ===== {skill_name} =====" in content:
        print(f"    - Block for {skill_name} already exists in {shared_filename}, skipping append.")
        return

    # Prepare new block
    # Ensure there is a newline before
    if not content.endswith("\n"):
        content += "\n"
        
    new_block = f"\n<!-- ===== {skill_name} =====\n{body}\n===== /{skill_name} ===== -->\n"
    write_file(path, content + new_block)
    print(f"    - Appended content to {shared_filename}")

def process_skill(skill_path, skill_name):
    content = read_file(skill_path)
    if not content:
        return

    original_content = content
    modified = False
    
    print(f"Processing {skill_name}...")

    # We iterate over mapping keys. 
    # Important: Check content for each key.
    
    for header, filename in SECTION_MAPPING.items():
        # Check if already included
        search_include = f"<!-- INCLUDE: _meta/_skills/sections/{filename} -->"
        if search_include in content:
            continue
            
        extracted = extract_section(content, header)
        if not extracted:
            # Maybe the skill doesn't have this section, skip
            continue
            
        body, start, end = extracted
        
        # If section found, we need to decide:
        # 1. Is it static? If so, just replace.
        # 2. Is it specific? If so, append to shared file THEN replace.
        
        is_static = header in STATIC_SECTIONS
        
        # Heuristic: If body is suspiciously short or identical to generic, maybe treat as static? 
        # But for now, trust the configuration.
        # Actually Language Requirements is definitely static.
        
        if not is_static:
            # Check if empty (sometimes sections are empty placeholders)
            if body and len(body) > 10:
                append_to_shared_file(filename, skill_name, body)
            else:
                 print(f"    - Section {header} is empty/short, not appending to shared file.")

        # Replace in content
        # We replace the WHOLE section (header + body) with the include
        # content slice replacement
        replacement = f"<!-- INCLUDE: _meta/_skills/sections/{filename} -->\n\n"
        
        # We must be careful with indices because replacement changes length.
        # But we are doing string manipulation. 
        # Better to do replacements carefully or re-read?
        # A simple way is to use regex sub, but we already calculated indices.
        # Since we might perform multiple replacements, indices shift.
        # SOLUTION: We will perform replacements sequentially on the string 
        # and re-search for the next header in the *new* string? 
        # OR just use re.sub with a callback?
        
        # Let's perform one replacement and restart the loop? 
        # No, simpler: Read, Replace, Write, Read again? No inefficient.
        
        # Let's just use string replacement by regex since we know the header.
        pattern = r'(^|\n)## ' + re.escape(header) + r'\s*\n.*?(?=\n## |\Z)'
        # FLAGS: DOTALL for .*? to match newlines
        
        # We need to construct the replacement carefully.
        # If the file had "## Header\nBody\n\n", we replace with "<!-- INCLUDE ... -->\n\n"
        
        def replacer(match):
            # Keep the leading newline if capture group 1 matched it
            prefix = match.group(1)
            return prefix + replacement.strip() + "\n\n"

        # Apply replacement
        new_content = re.sub(pattern, replacer, content, flags=re.DOTALL)
        
        if new_content != content:
            content = new_content
            modified = True
            print(f"    - Replaced {header} with INCLUDE")
            
    if modified:
        write_file(skill_path, content)
        print(f"    -> Updated {skill_name}")

def main():
    if not os.path.exists(SKILLS_DIR):
        print(f"Directory {SKILLS_DIR} not found.")
        return

    # Loop through blueprint/skills
    for skill_name in sorted(os.listdir(SKILLS_DIR)):
        skill_path = os.path.join(SKILLS_DIR, skill_name, "SKILL.md")
        if os.path.exists(skill_path):
            process_skill(skill_path, skill_name)

if __name__ == "__main__":
    main()
