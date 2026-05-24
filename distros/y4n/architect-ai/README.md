---
name: y4n/architect-ai
display_name: Architect AI
description: AI-powered software architect environment with Claude Code, Codex, OpenCode, and Pi for code exploration and technical decision-making
status: experimental
devcontainer: arch-base
tags: [fish, node, claude-code, codex, opencode, ai-agents, arch]
packages: [fish-config, node-runtime, cli-essentials, tui-power, ai-agent-claude, ai-agent-codex, ai-agent-opencode, ai-agent-pi]
---

# y4n/architect-ai

A development environment for software architects who rely on AI agents for code exploration, codebase understanding, and technical decision-making.

## What's included

| Package | Purpose |
|---------|---------|
| `fish-config` | Fish shell — optimised for navigation and exploration, not script execution |
| `node-runtime` | Node.js LTS via nvm — required by all AI agent CLIs |
| `cli-essentials` | Core tools: fzf, ripgrep, bat, jq, eza, zoxide, atuin |
| `tui-power` | TUI tools: lazygit and friends for fast terminal navigation |
| `ai-agent-claude` | Anthropic's Claude Code — agentic, whole-codebase reasoning |
| `ai-agent-codex` | OpenAI Codex CLI — lightweight terminal coding agent |
| `ai-agent-opencode` | OpenCode — open-source, multi-provider AI coding agent |
| `ai-agent-pi` | Pi Coding Agent — autonomous code understanding and generation |

## Philosophy

A software architect rarely executes builds or runs tests directly — they navigate, explore, and reason. This distro is optimised for that mode of work:

- **Fish shell** for interactive navigation (abbreviations, autosuggestions, no POSIX baggage)
- **CLI essentials** for fast file search and exploration (fzf, ripgrep, bat)
- **Multiple AI agents** so you can pick the right model for each task — Claude for deep reasoning, Codex for quick edits, OpenCode for multi-provider flexibility, Pi for autonomous exploration

## Usage

```yaml
# dpod.yaml
distro: y4n/architect-ai@v0.2.0
```

Run `dpod-seed sync` to materialise the environment.
