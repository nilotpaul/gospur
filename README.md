# Go Spur

A modern CLI Tool to bootstrap scalable web applications without hassling with JavaScript. It provides pre-configured developer tooling with no bloat, flexible for easy one step deployments.

# What's better?

- Only the necessary pre-configuration (full-control).
- Auto JavaScript Bundling (Bring any npm library).
- Very Fast Live Reload (server & browser).
- `make dev` for dev and `make` for prod (one-click).
- Extra options like tailwind, vanilla css, HTMX. 


# Installation

```sh
go install -u github.com/nilotpaul/gospur@latest
```

or without installation

```sh
go run github.com/nilotpaul/gospur@latest init [project-name]
```

# Usage

```
gospur init [project-name]
```

# Docs

Read detailed usage and examples of every stack configured.

- [Go + Echo + Tailwind + HTMX](https://github.com/nilotpaul/gospur/tree/go-echo-tailwind-htmx)

# Coming Soon

- More Framework Options.
- Deployments strategies like docker.
- Better templating with [templ](https://templ.guide) and [fiber](https://docs.gofiber.io) (django syntax).
- Please suggest More 🙏🏼
