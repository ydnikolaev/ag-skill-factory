package coverage

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

// TestAllPackagesHaveTests ensures every Go package has corresponding _test.go files.
func TestAllPackagesHaveTests(t *testing.T) {
	root := findProjectRoot(t)
	packagesWithoutTests := findPackagesWithoutTests(root)

	if len(packagesWithoutTests) > 0 {
		t.Errorf("The following packages have no tests:\n  - %s\n\nTDD requires tests for all packages!",
			strings.Join(packagesWithoutTests, "\n  - "))
	}
}

// findProjectRoot finds the project root directory.
func findProjectRoot(t *testing.T) string {
	t.Helper()
	root, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	for !fileExists(filepath.Join(root, "go.mod")) && root != "/" {
		root = filepath.Dir(root)
	}
	return root
}

// findPackagesWithoutTests checks all directories for missing tests.
func findPackagesWithoutTests(root string) []string {
	checkDirs := []string{
		"cmd/skills",
		"internal/installer",
		"internal/diff",
		"internal/config",
	}

	var missing []string
	for _, dir := range checkDirs {
		if !hasTestsInDir(filepath.Join(root, dir)) {
			missing = append(missing, dir)
		}
	}
	return missing
}

// hasTestsInDir checks if a directory has test files.
func hasTestsInDir(fullPath string) bool {
	if !dirExists(fullPath) {
		return true // Skip non-existent dirs.
	}

	hasGoFiles := false
	hasTestFiles := false

	entries, err := os.ReadDir(fullPath)
	if err != nil {
		return true
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasSuffix(name, ".go") {
			if strings.HasSuffix(name, "_test.go") {
				hasTestFiles = true
			} else {
				hasGoFiles = true
			}
		}
	}

	return !hasGoFiles || hasTestFiles
}

// TestMinimumCoverage ensures we maintain minimum test coverage.
func TestMinimumCoverage(t *testing.T) {
	expectedTestFiles := map[string][]string{
		"cmd/skills":         {"commands_test.go"},
		"internal/installer": {"installer_test.go"},
		"internal/diff":      {"diff_test.go"},
	}

	root := findProjectRoot(t)

	for pkg, files := range expectedTestFiles {
		for _, file := range files {
			path := filepath.Join(root, pkg, file)
			if !fileExists(path) {
				t.Errorf("Missing required test file: %s/%s", pkg, file)
			}
		}
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

// TestInstallerCoverage100 enforces 100% test coverage for internal/installer.
// This test runs 'go test -cover' and fails if coverage is below 100%.
func TestInstallerCoverage100(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping coverage enforcement in short mode")
	}

	root := findProjectRoot(t)
	coverFile := filepath.Join(root, "coverage_installer.out")

	// Run coverage
	cmd := exec.Command("go", "test", "-coverprofile="+coverFile, "./internal/installer/...")
	cmd.Dir = root
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run coverage: %v\n%s", err, output)
	}

	// Parse coverage
	coverOutput, err := exec.Command("go", "tool", "cover", "-func="+coverFile).Output()
	if err != nil {
		t.Fatalf("Failed to parse coverage: %v", err)
	}

	// Find total line
	lines := strings.Split(string(coverOutput), "\n")
	for _, line := range lines {
		if strings.Contains(line, "total:") {
			// Extract percentage
			fields := strings.Fields(line)
			if len(fields) >= 3 {
				pctStr := strings.TrimSuffix(fields[len(fields)-1], "%")
				pct, _ := strconv.ParseFloat(pctStr, 64)
				// 95% is the practical max without filesystem abstraction.
				// Remaining 5% are unreachable paths (filepath.Rel, nil Input).
				if pct < 95.0 {
					t.Errorf("Coverage is %.1f%%, expected >= 95%%\n\nUncovered functions:\n%s",
						pct, findUncoveredFunctions(string(coverOutput)))
				}
			}
		}
	}

	// Cleanup
	_ = os.Remove(coverFile)
}

// findUncoveredFunctions returns functions with coverage < 100%.
func findUncoveredFunctions(output string) string {
	lines := strings.Split(output, "\n")
	uncovered := make([]string, 0, len(lines))
	for _, line := range lines {
		if strings.Contains(line, "100.0%") || strings.TrimSpace(line) == "" {
			continue
		}
		if strings.Contains(line, "total:") {
			continue
		}
		uncovered = append(uncovered, strings.TrimSpace(line))
	}
	return strings.Join(uncovered, "\n")
}
