# Ultimate Golang Stack

**Build scalable full-stack solutions using just Go without hassling with JS.**

This is a minimal project template designed to be highly configurable for your requirements. It includes only two essential dependency.

# Features

- Auto server & browser reload (like HMR).
- [TailwindCSS](https://tailwindcss.com) & [Preline](https://preline.co) pre-configured for building modern UIs.
- [HTMX](https://htmx.org) for enhanced HTML interactivity.
- Bundling JavaScript with [esbuild](https://esbuild.github.io) (No third-party CDNs).
- To run in production, all you need is `public` folder.

# Prerequisites

- Go
- Node.js with your preferred package manager (e.g., npm, yarn, or pnpm)
- [wgo](https://github.com/bokwoon95/wgo) for live server reload.

# Installation

**After cloing the repo run these:**

```sh
# Needed for live reload
go install github.com/bokwoon95/wgo@latest
# Install node Deps
npm install
# Install Go Deps
go mod tidy
```

**To start dev server run:**

```sh
make watch
```

**To start prod server run:**

```
make
```

# Deployment

- You only need the built binary by Go and the bundled CSS & JS to run in production.
- Environment Variables can be loaded via `.env` file or runtime (already configured).

# How easy it is to use?

```go
func handleGetHome(c echo.Context) error {
	return c.Render(http.StatusOK, "Home", map[string]any{
		"IsDEV": true,
		"Title": "Go + Echo + HTMX",
		"Desc":  "Best for building Full-Stack Applications with minimal JavaScript",
	})
}
```

```html
{{ define "Home" }}
<h1 class="text-4xl">{{ .Title }}</h1>
<p class="mt-4">{{ .Desc }}</p>
{{ end }}
```

Only this much code is needed to render a page and dynamic content gets even more easier with Go Templates & HTMX combined with PrelineUI where components are by default interactive.

# Quick Tips

- **HTML Routes:** Render templates using handlers like the example above.
- **JSON Routes:** Prefix API endpoints with `/api/json`. The configuration ensures JSON responses even on errors.

For example, `/api/json/example` will always return a JSON response, whereas /`example` would render a template or custom HTML error pages.

# Advanced Usage

**You can also install any node library and use it.**

1.  Install the library you want.
2.  Update the esbuild configuration:

    ```js
    build({
      // Add the main entrypoint
      entryPoints: ["node_modules/some-library/index.js"],
    });
    ```

3.  Include the bundled script in your templates:
    your lib will be bundled and store in `public/bundle`, find the exact path and include in your templates.

    ```html
    <script src="/public/bundle/some-library/index.js"></script>
    ```

# Links to Documentation

- [Echo](https://echo.labstack.com)
- [TailwindCSS](https://tailwindcss.com)
- [Preline](https://preline.co)
- [HTMX](https://htmx.org)
- [Esbuild](https://esbuild.github.io)
