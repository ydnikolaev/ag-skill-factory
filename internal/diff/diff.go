package diff

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	godiff "github.com/sergi/go-diff/diffmatchpatch"
)

// CompareDirectories compares two directories and returns a list of changes.
func CompareDirectories(dir1, dir2 string) ([]string, error) {
	var changes []string

	err := filepath.Walk(dir2, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		change, err := compareFile(dir1, dir2, path)
		if err != nil {
			return err
		}
		if change != "" {
			changes = append(changes, change)
		}
		return nil
	})

	return changes, err
}

// compareFile compares a single file between two directories.
func compareFile(dir1, dir2, path string) (string, error) {
	relPath, err := filepath.Rel(dir2, path)
	if err != nil {
		return "", err
	}

	localPath := filepath.Join(dir1, relPath)

	localInfo, err := os.Stat(localPath)
	if os.IsNotExist(err) {
		return color.GreenString("+ %s (new file)", relPath), nil
	}

	sourceInfo, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	if localInfo.Size() != sourceInfo.Size() {
		return generateDiffForFile(localPath, path, relPath)
	}

	return compareFileContents(localPath, path, relPath)
}

// compareFileContents compares the actual contents of two files.
func compareFileContents(localPath, sourcePath, relPath string) (string, error) {
	localContent, err := os.ReadFile(localPath)
	if err != nil {
		return "", err
	}

	sourceContent, err := os.ReadFile(sourcePath)
	if err != nil {
		return "", err
	}

	if string(localContent) != string(sourceContent) {
		return generateDiffForFile(localPath, sourcePath, relPath)
	}

	return "", nil
}

// generateDiffForFile generates a diff string for display.
func generateDiffForFile(file1, file2, relPath string) (string, error) {
	diffStr, err := generateFileDiff(file1, file2)
	if err != nil {
		return "", err
	}
	if diffStr != "" {
		return "M " + relPath + ":\n" + diffStr, nil
	}
	return "", nil
}

// generateFileDiff generates a unified diff between two files.
func generateFileDiff(file1, file2 string) (string, error) {
	content1, err := os.ReadFile(file1)
	if err != nil {
		return "", err
	}

	content2, err := os.ReadFile(file2)
	if err != nil {
		return "", err
	}

	dmp := godiff.New()
	diffs := dmp.DiffMain(string(content1), string(content2), true)

	if !hasActualChanges(diffs) {
		return "", nil
	}

	return formatDiff(diffs), nil
}

// hasActualChanges checks if there are non-equal diffs.
func hasActualChanges(diffs []godiff.Diff) bool {
	for _, d := range diffs {
		if d.Type != godiff.DiffEqual {
			return true
		}
	}
	return false
}

// formatDiff formats diff output with colors.
func formatDiff(diffs []godiff.Diff) string {
	var result strings.Builder
	for _, d := range diffs {
		switch d.Type {
		case godiff.DiffInsert:
			result.WriteString(color.GreenString("+ %s", d.Text))
		case godiff.DiffDelete:
			result.WriteString(color.RedString("- %s", d.Text))
		case godiff.DiffEqual:
			// Skip equal parts for brevity.
		}
	}
	return result.String()
}
