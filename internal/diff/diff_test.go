package diff

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCompareDirectories_IdenticalDirs(t *testing.T) {
	dir1 := t.TempDir()
	dir2 := t.TempDir()

	content := []byte("same content")
	writeTestFile(t, filepath.Join(dir1, "file.md"), content)
	writeTestFile(t, filepath.Join(dir2, "file.md"), content)

	changes, err := CompareDirectories(dir1, dir2)
	if err != nil {
		t.Fatalf("CompareDirectories failed: %v", err)
	}

	if len(changes) != 0 {
		t.Errorf("Expected 0 changes for identical dirs, got %d", len(changes))
	}
}

func TestCompareDirectories_NewFile(t *testing.T) {
	dir1 := t.TempDir()
	dir2 := t.TempDir()

	writeTestFile(t, filepath.Join(dir2, "new_file.md"), []byte("new content"))

	changes, err := CompareDirectories(dir1, dir2)
	if err != nil {
		t.Fatalf("CompareDirectories failed: %v", err)
	}

	if len(changes) != 1 {
		t.Errorf("Expected 1 change (new file), got %d", len(changes))
	}

	if len(changes) > 0 && !containsSubstring(changes[0], "new file") {
		t.Errorf("Expected 'new file' marker, got: %s", changes[0])
	}
}

func TestCompareDirectories_ModifiedFile(t *testing.T) {
	dir1 := t.TempDir()
	dir2 := t.TempDir()

	writeTestFile(t, filepath.Join(dir1, "file.md"), []byte("old content"))
	writeTestFile(t, filepath.Join(dir2, "file.md"), []byte("new content"))

	changes, err := CompareDirectories(dir1, dir2)
	if err != nil {
		t.Fatalf("CompareDirectories failed: %v", err)
	}

	if len(changes) != 1 {
		t.Errorf("Expected 1 change (modified), got %d", len(changes))
	}
}

func TestCompareDirectories_NestedDirs(t *testing.T) {
	dir1 := t.TempDir()
	dir2 := t.TempDir()

	createTestDir(t, filepath.Join(dir1, "subdir"))
	createTestDir(t, filepath.Join(dir2, "subdir"))

	writeTestFile(t, filepath.Join(dir1, "subdir", "file.md"), []byte("old"))
	writeTestFile(t, filepath.Join(dir2, "subdir", "file.md"), []byte("new"))

	changes, err := CompareDirectories(dir1, dir2)
	if err != nil {
		t.Fatalf("CompareDirectories failed: %v", err)
	}

	if len(changes) != 1 {
		t.Errorf("Expected 1 change in nested dir, got %d", len(changes))
	}
}

func TestGenerateFileDiff_ShowsChanges(t *testing.T) {
	dir := t.TempDir()

	file1 := filepath.Join(dir, "old.md")
	file2 := filepath.Join(dir, "new.md")

	writeTestFile(t, file1, []byte("line1\nline2\n"))
	writeTestFile(t, file2, []byte("line1\nline3\n"))

	diffStr, err := generateFileDiff(file1, file2)
	if err != nil {
		t.Fatalf("generateFileDiff failed: %v", err)
	}

	if diffStr == "" {
		t.Error("Expected non-empty diff for different files")
	}
}

func TestGenerateFileDiff_EmptyForIdentical(t *testing.T) {
	dir := t.TempDir()

	file1 := filepath.Join(dir, "file1.md")
	file2 := filepath.Join(dir, "file2.md")

	content := []byte("same content")
	writeTestFile(t, file1, content)
	writeTestFile(t, file2, content)

	diffStr, err := generateFileDiff(file1, file2)
	if err != nil {
		t.Fatalf("generateFileDiff failed: %v", err)
	}

	if diffStr != "" {
		t.Errorf("Expected empty diff for identical files, got: %s", diffStr)
	}
}

// Helper functions.

func writeTestFile(t *testing.T, path string, content []byte) {
	t.Helper()
	if err := os.WriteFile(path, content, 0o644); err != nil {
		t.Fatalf("Failed to write test file %s: %v", path, err)
	}
}

func createTestDir(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatalf("Failed to create test dir %s: %v", path, err)
	}
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
