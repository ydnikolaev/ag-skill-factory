// Package presets loads and resolves skill preset bundles.
package presets

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// rawPreset is the intermediate YAML structure.
type rawPreset struct {
	Description string      `yaml:"description"`
	Extends     yaml.Node   `yaml:"extends"`
	Skills      yaml.Node   `yaml:"skills"`
}

// Preset defines a bundle of skills.
type Preset struct {
	Description string
	Extends     []string
	Skills      []string
	AllSkills   bool // true if skills: "*"
}

// Config holds all presets from presets.yaml.
type Config map[string]Preset

// Load reads presets from the given blueprint path.
func Load(blueprintPath string) (Config, error) {
	presetsFile := filepath.Join(blueprintPath, "_meta", "presets.yaml")
	data, err := os.ReadFile(presetsFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read presets: %w", err)
	}

	var raw map[string]rawPreset
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse presets: %w", err)
	}

	config := make(Config)
	for name, rp := range raw {
		p := Preset{Description: rp.Description}
		
		// Parse extends
		p.Extends = parseStringOrSlice(rp.Extends)
		
		// Parse skills
		if rp.Skills.Kind == yaml.ScalarNode && rp.Skills.Value == "*" {
			p.AllSkills = true
		} else {
			p.Skills = parseStringOrSlice(rp.Skills)
		}
		
		config[name] = p
	}

	return config, nil
}

func parseStringOrSlice(node yaml.Node) []string {
	if node.Kind == yaml.ScalarNode {
		if node.Value != "" {
			return []string{node.Value}
		}
		return nil
	}
	if node.Kind == yaml.SequenceNode {
		var result []string
		for _, n := range node.Content {
			if n.Kind == yaml.ScalarNode {
				result = append(result, n.Value)
			}
		}
		return result
	}
	return nil
}

// ResolveSkills returns the full list of skills for a preset, resolving extends.
func (c Config) ResolveSkills(presetName string, allSkills []string) ([]string, error) {
	preset, ok := c[presetName]
	if !ok {
		return nil, fmt.Errorf("unknown preset: %s", presetName)
	}

	// If all skills
	if preset.AllSkills {
		return allSkills, nil
	}

	skillSet := make(map[string]bool)

	// Resolve extends first
	for _, ext := range preset.Extends {
		extSkills, err := c.ResolveSkills(ext, allSkills)
		if err != nil {
			return nil, err
		}
		for _, s := range extSkills {
			skillSet[s] = true
		}
	}

	// Add own skills
	for _, s := range preset.Skills {
		skillSet[s] = true
	}

	// Convert to slice
	result := make([]string, 0, len(skillSet))
	for s := range skillSet {
		result = append(result, s)
	}

	return result, nil
}

// List returns all preset names with descriptions.
func (c Config) List() []struct{ Name, Description string } {
	var list []struct{ Name, Description string }
	for name, preset := range c {
		// Skip internal stuff
		if strings.HasPrefix(name, "_") {
			continue
		}
		list = append(list, struct{ Name, Description string }{name, preset.Description})
	}
	return list
}
