# distros

Source-of-truth for verified, named dev environment compositions consumed by [dpod-seed](https://github.com/duyanh-y4n/dpod-seed).

## Directory layout

```
distros/
  <name>/
    distro.yaml   # composition spec
    README.md     # human-readable description
registry.yaml     # auto-generated on tag push — do not edit by hand
cmd/gen-registry/ # tool that generates registry.yaml
```

## distro.yaml schema

```yaml
name: my-distro
description: Short human-readable description
devcontainer: <profile>@<sha>
packages:
  - <package>@<sha>
```

## Contribution workflow

1. **Scaffold** a new distro:
   ```sh
   dpod-seed scaffold distro <name>
   ```
2. Edit `distros/<name>/distro.yaml` — set `devcontainer` and `packages` pins.
3. Add `distros/<name>/README.md` describing the distro.
4. Open a **pull request** — CI runs `gen-registry` in dry-run mode and validates all `distro.yaml` files.
5. After review and merge, a maintainer pushes a `vX.Y.Z` tag.
6. The release workflow regenerates `registry.yaml` and commits it to `main`.

## Release

```sh
git tag vX.Y.Z && git push origin vX.Y.Z
```

CI will regenerate and commit `registry.yaml` automatically.
