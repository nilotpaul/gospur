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

## Create a new project
```sh
gospur init [project-name]
```
## Update the CLI
```sh
gospur update
``` 

Check more options by:
```sh
# help for cli
gospur --help
# help for each command
gospur init --help
```

# Docs

Read detailed usage and examples of every stack configured.

**(Stacks)**
- [Go + Echo + Templates](/docs/go-echo-templates.md)
- [Go + Fiber + Templates](/docs/go-fiber-templates.md)
- [Go + Chi + Templates](/docs/go-chi-templates.md)
- [Go + Seperate Client](/docs/go-seperate-client.md)

**(Others)**
- [Development Usage](/docs/development-usage.md)
- [Recommendations](/docs/recommendations/index.md)

# Configuration Options

The configuration options include settings from various web frameworks to different rendering modes. For a detailed list, please check the [Configuration Options Docs](/docs/configuration.md).

# Coming Soon

- More Framework Options.
- Different Rendering Strategies (~~seperate client~~, [templ](https://templ.guide)).
- More examples and documentation.
- Please suggest More 🙏🏼
