# GoSpur

<img src="/assets/gospur.png" width="300" alt="GoSpur Logo" />

A modern CLI Tool to bootstrap scalable web applications without hassling with JavaScript. It provides pre-configured developer tooling with no bloat, flexible for easy one step deployments.


# What's better?

- Only the necessary pre-configuration (full-control).
- Auto JavaScript Bundling (Bring any npm library).
- Very Fast Live Reload (server & browser).
- `make dev` for dev and `make` for prod (one-click).
- Extra options like tailwind, vanilla css, HTMX. 


# Installation

```sh
go install github.com/nilotpaul/gospur@latest
```

or without installation

```sh
go run github.com/nilotpaul/gospur@latest init [project-name]
```

or download prebuilt binary

- [Download from here](https://github.com/nilotpaul/gospur/releases/latest)
- Extract and run `./gospur init`

# Usage

```
gospur init [project-name]
```

# Docs

Read detailed usage and examples of every stack configured.

**(Stacks)**
- [Go + Echo + Templates](/docs/go-echo-templates.md)
- [Go + Fiber + Templates](/docs/go-fiber-templates.md)

**(Others)**
- [Development Usage](/docs/development-usage.md)

# Configuration Options

**Web Framework**
- Echo  
- Fiber

**Styling**
- Vanilla CSS  
- Tailwind

**UI Library** 
- Preline  
- DaisyUI

**Extra Options**
- HTMX  
- Dockerfile

# Coming Soon

- More Framework Options.
- More examples and documentation.
- Please suggest More 🙏🏼
