---
name: y4n/docker-ops-base
display_name: Docker Ops Base
description: Arch-based environment for container lifecycle management and security
status: stable
devcontainer: arch-base
tags: [docker, containers, security, registry, arch]
packages: [cli-essentials, git-extras, zsh-config, neovim-config, docker-tools, container-inspect, registry-tools, container-security]
---

# Docker Ops Base

Arch-based devpod for container lifecycle management, inspection, registry operations, and security scanning.

## Contents

| Layer | Packages |
|-------|---------|
| Shared | cli-essentials, git-extras, zsh-config, neovim-config |
| Role | docker-tools, container-inspect, registry-tools, container-security |

## Usage

```yaml
# dpod.yaml
distro: y4n/docker-ops-base@v0.1.0
```

## Common overrides

Add power-user TUI tools:

```yaml
distro: y4n/docker-ops-base@v0.1.0
packages:
  - tui-power@v0.1.0
```
