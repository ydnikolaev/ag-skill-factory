// Package doctor provides diagnostic checks for the blueprint.
package doctor

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Result holds the diagnostic results.
type Result struct {
	Errors   []string
	Warnings []string
}

// Check runs all diagnostics on the given blueprint path.
func Check(blueprintPath string) (*Result, error) {
	result := &Result{}

	skillsPath := filepath.Join(blueprintPath, "skills")
	if _, err := os.Stat(skillsPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("skills directory not found: %s", skillsPath)
	}

	// Get all skill names for reference checking
	skillNames := make(map[string]bool)
	entries, err := os.ReadDir(skillsPath)
	if err != nil {
		return nil, err
	}
	for _, e := range entries {
		if e.IsDir() {
			skillNames[e.Name()] = true
		}
	}

	// Add factory skills
	factorySkills := []string{"skill-creator", "skill-factory-expert", "skill-interviewer", "skill-updater", "workflow-creator"}
	for _, s := range factorySkills {
		skillNames[s] = true
	}

	// Check each skill
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		skillPath := filepath.Join(skillsPath, entry.Name())
		checkSkill(skillPath, skillNames, result)
	}

	return result, nil
}

func checkSkill(skillPath string, knownSkills map[string]bool, result *Result) {
	skillMD := filepath.Join(skillPath, "SKILL.md")
	content, err := os.ReadFile(skillMD)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("%s: cannot read SKILL.md", filepath.Base(skillPath)))
		return
	}

	text := string(content)
	skillName := filepath.Base(skillPath)

	// Check for broken "See" references
	seePattern := regexp.MustCompile(`See\s+([a-zA-Z_]+/[^\s\)]+)`)
	matches := seePattern.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		refPath := match[1]
		fullPath := filepath.Join(skillPath, refPath)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			result.Errors = append(result.Errors, fmt.Sprintf("%s: broken link '%s'", skillName, refPath))
		}
	}

	// Check for unknown @skill references
	skillRefPattern := regexp.MustCompile(`(?:^|[^a-zA-Z0-9])@([a-z][a-z0-9-]+)(?:[^a-zA-Z0-9@.]|$)`)
	skillMatches := skillRefPattern.FindAllStringSubmatch(text, -1)
	for _, match := range skillMatches {
		refSkill := match[1]
		if !knownSkills[refSkill] {
			result.Warnings = append(result.Warnings, fmt.Sprintf("%s: unknown skill @%s", skillName, refSkill))
		}
	}

	// Check for hardcoded paths
	hardcodedPattern := regexp.MustCompile(`(/Users/|/home/|C:\\)`)
	scanner := bufio.NewScanner(strings.NewReader(text))
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if hardcodedPattern.MatchString(line) {
			result.Warnings = append(result.Warnings, fmt.Sprintf("%s:%d: hardcoded path detected", skillName, lineNum))
		}
	}
}
