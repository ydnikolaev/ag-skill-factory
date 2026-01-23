// Package skills implements the CLI commands for ag-skill-factory.
//
// Commands:
//   - install: Bootstrap .agent/ structure in current workspace
//   - update: Pull latest skill changes from factory
//   - backport: Push local skill changes back to factory
//   - list: Show skill inventory with sync status
//
// Configuration is loaded from ~/.config/ag-skills/config.yaml
// with the following options:
//   - source: Path to skill factory (squads/ directory)
//   - global_path: Path for global skill copies
package skills
