package installer

import (
	"testing"
)

func TestRewriteStandardsPaths_BasicTransform(t *testing.T) {
	input := "See `_standards/TDD_PROTOCOL.md` for details"
	expected := "See `.agent/rules/tdd_protocol.md` for details"

	result := rewriteStandardsPaths(input)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestRewriteStandardsPaths_MultipleOccurrences(t *testing.T) {
	input := "Check `_standards/TDD_PROTOCOL.md` and `_standards/GIT_PROTOCOL.md`"
	result := rewriteStandardsPaths(input)

	if contains(result, "_standards/") {
		t.Errorf("Still contains _standards/: %s", result)
	}
	if !contains(result, ".agent/rules/tdd_protocol.md") {
		t.Error("Missing tdd_protocol.md transformation")
	}
	if !contains(result, ".agent/rules/git_protocol.md") {
		t.Error("Missing git_protocol.md transformation")
	}
}

func TestRewriteStandardsPaths_NoChange(t *testing.T) {
	input := "Regular text without standards paths"
	result := rewriteStandardsPaths(input)

	if result != input {
		t.Errorf("Expected no change, got %q", result)
	}
}

func TestRewriteStandardsPaths_MultilineContent(t *testing.T) {
	input := `# Protocol Reference

Please follow these protocols:
- [TDD Protocol](_standards/TDD_PROTOCOL.md)
- [Git Protocol](_standards/GIT_PROTOCOL.md)

More text here.`

	result := rewriteStandardsPaths(input)

	if contains(result, "_standards/") {
		t.Errorf("Still contains _standards/: %s", result)
	}
	if !contains(result, ".agent/rules/tdd_protocol.md") {
		t.Error("Missing tdd_protocol.md in multiline")
	}
}

func TestRewriteStandardsLine_WithBackticks(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "backticks",
			input:    "See `_standards/TDD_PROTOCOL.md` for details",
			expected: "See `.agent/rules/tdd_protocol.md` for details",
		},
		{
			name:     "parentheses",
			input:    "Reference: (_standards/TDD_PROTOCOL.md)",
			expected: "Reference: (.agent/rules/tdd_protocol.md)",
		},
		{
			name:     "brackets",
			input:    "[TDD Protocol](_standards/TDD_PROTOCOL.md)",
			expected: "[TDD Protocol](.agent/rules/tdd_protocol.md)",
		},
		{
			name:     "at end of line",
			input:    "Check _standards/TECH_DEBT_PROTOCOL.md",
			expected: "Check .agent/rules/tech_debt_protocol.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rewriteStandardsLine(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestRewriteStandardsLine_PreservesCase(t *testing.T) {
	// Original filename case is converted to lowercase
	input := "See _standards/MY_CUSTOM_RULE.md"
	result := rewriteStandardsLine(input)

	if !contains(result, "my_custom_rule.md") {
		t.Errorf("Expected lowercase, got %s", result)
	}
}
