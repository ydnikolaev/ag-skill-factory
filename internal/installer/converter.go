package installer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

// convertStandardsToRules converts _standards files to rule format.
func (i *Installer) convertStandardsToRules(standardsPath string) (int, error) {
	count := 0

	entries, err := os.ReadDir(standardsPath)
	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		src := filepath.Join(standardsPath, entry.Name())
		dst := filepath.Join(i.Target, "rules", strings.ToLower(entry.Name()))

		if err := i.convertToRule(src, dst); err != nil {
			return count, err
		}
		count++
	}

	return count, nil
}

// convertToRule converts a markdown file to rule format with YAML frontmatter.
func (i *Installer) convertToRule(src, dst string) error {
	content, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	description := extractFirstHeading(string(content))

	var builder strings.Builder
	builder.WriteString("---\n")
	builder.WriteString(fmt.Sprintf("description: %s\n", description))
	builder.WriteString("---\n\n")
	builder.WriteString(string(content))

	return os.WriteFile(dst, []byte(builder.String()), 0o644)
}

// extractFirstHeading extracts the first markdown heading.
func extractFirstHeading(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
	}
	return "Rule"
}

// copyMetaFilesToRules copies TEAM.md and PIPELINE.md to rules.
func (i *Installer) copyMetaFilesToRules(result *InstallResult) error {
	files := []string{"TEAM.md", "PIPELINE.md"}
	for _, file := range files {
		src := filepath.Join(i.Source, file)
		if _, err := os.Stat(src); err != nil {
			continue
		}
		dst := filepath.Join(i.Target, "rules", strings.ToLower(file))
		if err := i.convertToRule(src, dst); err != nil {
			color.Yellow("Warning: failed to convert %s: %v", file, err)
		} else {
			result.RuleCount++
		}
	}
	return nil
}
