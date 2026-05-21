---
name: y4n/platform-eng-base
display_name: Platform Engineering Base
description: Arch-based environment for cloud-native platform engineers
status: stable
devcontainer: arch-base
tags: [kubernetes, aws, terraform, gitops, arch]
packages: [cli-essentials, git-extras, zsh-config, neovim-config, k8s-tools, aws-cli, observability, iac-tools, flux]
---

# Platform Engineering Base

Arch-based devpod for cloud-native platform engineers working with standard cloud Kubernetes (EKS, GKE, AKS).

## Contents

| Layer | Packages |
|-------|---------|
| Shared | cli-essentials, git-extras, zsh-config, neovim-config |
| Role | k8s-tools, aws-cli, observability, iac-tools, flux |

## Usage

```yaml
# dpod.yaml
distro: y4n/platform-eng-base@v0.1.0
```

## Common overrides

Add GCP support:

```yaml
distro: y4n/platform-eng-base@v0.1.0
packages:
  - gcloud@v0.1.0
```

Add Azure support:

```yaml
distro: y4n/platform-eng-base@v0.1.0
packages:
  - azure-cli@v0.1.0
```

Add ArgoCD CLI:

```yaml
distro: y4n/platform-eng-base@v0.1.0
packages:
  - argocd@v0.1.0
```

Add power-user TUI tools:

```yaml
distro: y4n/platform-eng-base@v0.1.0
packages:
  - tui-power@v0.1.0
```
