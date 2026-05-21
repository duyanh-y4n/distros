---
name: y4n/frontend-web-base
display_name: Frontend Web Base
description: Debian Slim environment for web frontend development and testing
status: stable
devcontainer: debian-slim
tags: [node, typescript, react, playwright, debian]
packages: [cli-essentials, git-extras, zsh-config, neovim-config, node-runtime, bun, browser-test, js-linting]
---

# Frontend Web Base

Debian Slim devpod for web frontend development. Primary use: running dev servers, builds, and browser tests inside the container.

## Contents

| Layer | Packages |
|-------|---------|
| Shared | cli-essentials, git-extras, zsh-config, neovim-config |
| Role | node-runtime, bun, browser-test, js-linting |

## Usage

```yaml
# dpod.yaml
distro: y4n/frontend-web-base@v0.1.0
```

## Common overrides

Add power-user TUI tools:

```yaml
distro: y4n/frontend-web-base@v0.1.0
packages:
  - tui-power@v0.1.0
```
