// Package diff provides directory and file comparison utilities.
//
// Features:
//   - Compare two directories recursively
//   - Detect new, modified, and deleted files
//   - Generate unified diff output for file changes
//
// Usage:
//
//	changes, err := diff.CompareDirectories(sourceDir, targetDir)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, change := range changes {
//	    fmt.Println(change)
//	}
package diff
