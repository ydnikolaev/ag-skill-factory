// Package installer provides skill installation and synchronization logic.
//
// The installer handles:
//   - Installing skills from factory to project workspaces
//   - Updating skills with diff preview and confirmation
//   - Backporting changes from projects back to factory
//   - Converting standards to rules with YAML frontmatter
//
// Usage:
//
//	inst := installer.New(sourcePath, targetPath, globalPath)
//	result, err := inst.Install()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Installed %d skills\n", result.SkillCount)
package installer
