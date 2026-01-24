package installer

import (
	"strings"
)

// rewriteStandardsPaths transforms _standards/ references to .agent/rules/.
// Example: `_standards/TDD_PROTOCOL.md` → `.agent/rules/tdd_protocol.md`
func rewriteStandardsPaths(content string) string {
	// Pattern: _standards/SOMETHING.md → .agent/rules/something.md
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if strings.Contains(line, "_standards/") {
			lines[i] = rewriteStandardsLine(line)
		}
	}
	return strings.Join(lines, "\n")
}

// rewriteStandardsLine rewrites a single line's _standards references.
func rewriteStandardsLine(line string) string {
	// Find _standards/XXX.md patterns and replace
	// Handle: `_standards/TDD_PROTOCOL.md` → `.agent/rules/tdd_protocol.md`
	result := line

	// Find all occurrences of _standards/something.md
	idx := strings.Index(result, "_standards/")
	for idx != -1 {
		// Find the end of the path (space, `, ), ], or end of line)
		endIdx := idx + len("_standards/")
		for endIdx < len(result) {
			c := result[endIdx]
			if c == ' ' || c == '`' || c == ')' || c == ']' || c == '"' || c == '\'' {
				break
			}
			endIdx++
		}

		oldPath := result[idx:endIdx]
		// Extract filename from _standards/FILENAME.md
		filename := strings.TrimPrefix(oldPath, "_standards/")
		newPath := ".agent/rules/" + strings.ToLower(filename)

		result = result[:idx] + newPath + result[endIdx:]

		// Look for next occurrence
		idx = strings.Index(result[idx+len(newPath):], "_standards/")
		if idx != -1 {
			idx += idx + len(newPath)
		}
	}

	return result
}
