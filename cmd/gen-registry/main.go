package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v3"
)

const repoURL = "https://github.com/iamy4n-dev/distros"

type distroYAML struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Status      string `yaml:"status"`
}

type registryEntry struct {
	Name         string `yaml:"name"`
	Description  string `yaml:"description"`
	Status       string `yaml:"status"`
	LatestTag    string `yaml:"latestTag"`
	ChangelogURL string `yaml:"changelogUrl"`
}

type registryPayload struct {
	Distros []registryEntry `yaml:"distros"`
}

// Generate reads distros from distrosDir and returns a registry.yaml stamped with tag.
// Supports both flat (distros/{name}) and org-namespaced (distros/{org}/{name}) layouts.
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
		dir := filepath.Join(distrosDir, e.Name())
		if _, err := os.Stat(filepath.Join(dir, "distro.yaml")); err == nil {
			// Leaf distro
			entry, err := toEntry(dir, e.Name(), tag)
			if err != nil {
				return nil, err
			}
			distros = append(distros, entry)
		} else {
			// Org namespace — recurse one level
			subEntries, err := os.ReadDir(dir)
			if err != nil {
				return nil, fmt.Errorf("read org dir %s: %w", e.Name(), err)
			}
			for _, sub := range subEntries {
				if !sub.IsDir() {
					continue
				}
				subDir := filepath.Join(dir, sub.Name())
				entry, err := toEntry(subDir, e.Name()+"/"+sub.Name(), tag)
				if err != nil {
					return nil, err
				}
				distros = append(distros, entry)
			}
		}
	}

	sort.Slice(distros, func(i, j int) bool { return distros[i].Name < distros[j].Name })

	out, err := yaml.Marshal(registryPayload{Distros: distros})
	if err != nil {
		return nil, fmt.Errorf("marshal registry: %w", err)
	}
	return out, nil
}

func toEntry(dir, label, tag string) (registryEntry, error) {
	dy, err := readDistroYAML(filepath.Join(dir, "distro.yaml"))
	if err != nil {
		return registryEntry{}, fmt.Errorf("distros/%s: %w", label, err)
	}
	return registryEntry{
		Name:         dy.Name,
		Description:  dy.Description,
		Status:       dy.Status,
		LatestTag:    tag,
		ChangelogURL: fmt.Sprintf("%s/releases/tag/%s", repoURL, tag),
	}, nil
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
