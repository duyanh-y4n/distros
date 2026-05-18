package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v3"
)

const repoURL = "https://github.com/duyanh-y4n/distros"

type distroYAML struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type registryEntry struct {
	Name         string `yaml:"name"`
	Description  string `yaml:"description"`
	LatestTag    string `yaml:"latestTag"`
	ChangelogURL string `yaml:"changelogUrl"`
}

type registryPayload struct {
	Distros []registryEntry `yaml:"distros"`
}

// Generate reads distros from distrosDir and returns a registry.yaml stamped with tag.
func Generate(distrosDir, tag string) ([]byte, error) {
	entries, err := os.ReadDir(distrosDir)
	if err != nil {
		return nil, fmt.Errorf("read distros dir: %w", err)
	}

	var distros []registryEntry
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		dy, err := readDistroYAML(filepath.Join(distrosDir, e.Name(), "distro.yaml"))
		if err != nil {
			return nil, fmt.Errorf("distros/%s: %w", e.Name(), err)
		}
		distros = append(distros, registryEntry{
			Name:         dy.Name,
			Description:  dy.Description,
			LatestTag:    tag,
			ChangelogURL: fmt.Sprintf("%s/releases/tag/%s", repoURL, tag),
		})
	}

	sort.Slice(distros, func(i, j int) bool { return distros[i].Name < distros[j].Name })

	out, err := yaml.Marshal(registryPayload{Distros: distros})
	if err != nil {
		return nil, fmt.Errorf("marshal registry: %w", err)
	}
	return out, nil
}

func readDistroYAML(path string) (*distroYAML, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read distro.yaml: %w", err)
	}
	var dy distroYAML
	if err := yaml.Unmarshal(data, &dy); err != nil {
		return nil, fmt.Errorf("parse distro.yaml: %w", err)
	}
	if dy.Name == "" {
		return nil, fmt.Errorf("distro.yaml: missing required field 'name'")
	}
	return &dy, nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: gen-registry <distros-dir> <tag>")
		os.Exit(1)
	}
	out, err := Generate(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Print(string(out))
}
