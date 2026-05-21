package main

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"
)

type registryFile struct {
	Distros []registryEntry `yaml:"distros"`
}

func writeDistro(t *testing.T, dir, name, content string) {
	t.Helper()
	p := filepath.Join(dir, name)
	if err := os.MkdirAll(p, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(p, "distro.yaml"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func parseRegistry(t *testing.T, data []byte) registryFile {
	t.Helper()
	var rf registryFile
	if err := yaml.Unmarshal(data, &rf); err != nil {
		t.Fatalf("invalid registry.yaml: %v", err)
	}
	return rf
}

func TestGenerate_EmptyDir_EmitsEmptyList(t *testing.T) {
	dir := t.TempDir()

	out, err := Generate(dir, "v1.0.0")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	rf := parseRegistry(t, out)
	if rf.Distros != nil && len(rf.Distros) != 0 {
		t.Errorf("want empty distros list, got %v", rf.Distros)
	}
}

func TestGenerate_MissingName_ReturnsError(t *testing.T) {
	dir := t.TempDir()
	writeDistro(t, dir, "bad-distro", `description: no name field
devcontainer: arch-base@abc123
`)

	_, err := Generate(dir, "v1.0.0")
	if err == nil {
		t.Fatal("want error for missing name, got nil")
	}
	if !containsAll(err.Error(), "bad-distro", "name") {
		t.Errorf("error %q should mention the distro dir and the missing field", err.Error())
	}
}

func containsAll(s string, subs ...string) bool {
	for _, sub := range subs {
		found := false
		for i := 0; i <= len(s)-len(sub); i++ {
			if s[i:i+len(sub)] == sub {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func TestGenerate_TwoDistros_SortedByName(t *testing.T) {
	dir := t.TempDir()
	writeDistro(t, dir, "zebra-distro", `name: zebra-distro
description: Comes last alphabetically
devcontainer: arch-base@abc123
packages: []
`)
	writeDistro(t, dir, "alpha-distro", `name: alpha-distro
description: Comes first alphabetically
devcontainer: arch-base@abc123
packages: []
`)

	out, err := Generate(dir, "v2.0.0")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	rf := parseRegistry(t, out)
	if len(rf.Distros) != 2 {
		t.Fatalf("want 2 distros, got %d", len(rf.Distros))
	}
	if rf.Distros[0].Name != "alpha-distro" {
		t.Errorf("first entry: want %q, got %q", "alpha-distro", rf.Distros[0].Name)
	}
	if rf.Distros[1].Name != "zebra-distro" {
		t.Errorf("second entry: want %q, got %q", "zebra-distro", rf.Distros[1].Name)
	}
}

func TestGenerate_SingleDistro(t *testing.T) {
	dir := t.TempDir()
	writeDistro(t, dir, "example", `
name: example
description: An example distro
status: stable
devcontainer: arch-base@abc123
packages:
  - shell-zsh@def456
`)

	out, err := Generate(dir, "v1.2.3")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	rf := parseRegistry(t, out)
	if len(rf.Distros) != 1 {
		t.Fatalf("want 1 distro, got %d", len(rf.Distros))
	}

	e := rf.Distros[0]
	if e.Name != "example" {
		t.Errorf("name: want %q, got %q", "example", e.Name)
	}
	if e.Description != "An example distro" {
		t.Errorf("description: want %q, got %q", "An example distro", e.Description)
	}
	if e.Status != "stable" {
		t.Errorf("status: want %q, got %q", "stable", e.Status)
	}
	if e.LatestTag != "v1.2.3" {
		t.Errorf("latestTag: want %q, got %q", "v1.2.3", e.LatestTag)
	}
	wantURL := "https://github.com/iamy4n-dev/distros/releases/tag/v1.2.3"
	if e.ChangelogURL != wantURL {
		t.Errorf("changelogUrl: want %q, got %q", wantURL, e.ChangelogURL)
	}
}

func writeOrgDistro(t *testing.T, dir, org, name, content string) {
	t.Helper()
	p := filepath.Join(dir, org, name)
	if err := os.MkdirAll(p, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(p, "distro.yaml"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestGenerate_OrgNamespacedDistro(t *testing.T) {
	dir := t.TempDir()
	writeOrgDistro(t, dir, "acme", "my-distro", `name: acme/my-distro
description: Org-namespaced distro
status: stable
devcontainer: arch-base@v1.0.0
packages: []
`)

	out, err := Generate(dir, "v1.0.0")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	rf := parseRegistry(t, out)
	if len(rf.Distros) != 1 {
		t.Fatalf("want 1 distro, got %d: %v", len(rf.Distros), rf.Distros)
	}
	if rf.Distros[0].Name != "acme/my-distro" {
		t.Errorf("name: want %q, got %q", "acme/my-distro", rf.Distros[0].Name)
	}
}

func TestGenerate_MixedFlatAndOrg(t *testing.T) {
	dir := t.TempDir()
	writeDistro(t, dir, "flat-distro", `name: flat-distro
description: Flat layout distro
status: experimental
devcontainer: arch-base@v1.0.0
packages: []
`)
	writeOrgDistro(t, dir, "acme", "org-distro", `name: acme/org-distro
description: Org-namespaced distro
status: stable
devcontainer: arch-base@v1.0.0
packages: []
`)

	out, err := Generate(dir, "v1.0.0")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	rf := parseRegistry(t, out)
	if len(rf.Distros) != 2 {
		t.Fatalf("want 2 distros, got %d: %v", len(rf.Distros), rf.Distros)
	}
}
