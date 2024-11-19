# Go Spur

A modern CLI Tool to bootstrap scalable web applications without hassling with JavaScript. It provides pre-configured developer tooling with no bloat, flexible for easy one step deployments.

# Installation

```sh
go install github.com/nilotpaul/gospur@latest
```

# Usage

```
gospur init [project-name]
```

**Go Spur is WIP ⚒️, you'll get a default stack(Go + Echo + Tailwind + HTMX) for now.**

# Docs

Read detailed usage and examples of every stack configured.

- [Go + Echo + Tailwind + HTMX](https://github.com/nilotpaul/gospur/tree/go-templates-htmx)

# Known Issues

1. Black color used for text not visible in black background terminals. [#1](https://github.com/nilotpaul/gospur/issues/1)
2. When initialising a project in the current directory, it always says directory is not empty. [#2](https://github.com/nilotpaul/gospur/issues/3)
