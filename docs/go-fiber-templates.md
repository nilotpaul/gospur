# Go + Fiber + Templates

This is a minimal project template designed to be highly configurable for your requirements.

# Prerequisites

- Go
- Node.js with your preferred package manager (e.g., npm, yarn, or pnpm)
- [wgo](https://github.com/bokwoon95/wgo) for live server reload.

# Installation

**Run: `gospur init [project-name]`**

## Post Installation

```sh
# Needed for live reload
go install github.com/bokwoon95/wgo@latest
# Install node Deps
npm install
# Install Go Deps
go mod tidy
```

**To start dev server run:**

The Fiber server takes a few milliseconds more than others to start up, during reload it might feel a little slower, in that case pls [reduce or remove the delay](/docs/development-usage.md#if-auto-browser-reload-feels-slow).

```sh
make dev
```

**To start prod server run:**

```
make
```

# Deployment

You only need:

- The built binary in `bin` folder.

> **Note: All the assets in `public` and `web` folder will be embedded in the binary.**

- Commands to build for production:
```sh
# build cmd:
node ./esbuild.config.js
go build -tags '!dev' -o bin/build

# run cmd: 
ENVIRONMENT=PRODUCTION ./bin/build
```

# How easy it is to use?

> **Note: By default it'll use the root layout**

## Simple Example
```go
func handleGetHome(c *fiber.Ctx) error {
	return c.Render("Home", map[string]any{
		"Title": "GoSpur",
		"Desc":  "Best for building Full-Stack Applications with minimal JavaScript",
	})
}
```
```html
<h1 class="text-4xl">{{ .Ctx.Title }}</h1>
<p class="mt-4">{{ .Ctx.Desc }}</p>
```
Only this much code is needed to render a page.

## With Custom Layout
```go
func handleGetHome(c *fiber.Ctx) error {
	return c.Render("Other", map[string]any{
		"Title": "Other Page",
	}, "layouts/Layout.html")
}
```

# Templates

By default you'll get the stack with Go HTML Templates, but Fiber supports many templating engines like django.

It's very easy to swap but in our case there're few extra steps.
[See all supported engines](https://docs.gofiber.io/guide/templates#supported-engines).

## Using Django

Install and fix the import `github.com/gofiber/template/django/v3`
> **Note: Keep track of the version it might change in future.**

```go
// (Only change the part shown)
//
// build_dev.go
func LoadTemplates() *django.Engine {
	return django.New("web", ".html")
}
// build_prod.go
func LoadTemplates() *django.Engine {
	subFS, err := fs.Sub(templateFS, "web")
	if err != nil {
		panic(err)
	}

	return html.NewFileSystem(http.FS(subFS), ".html")
}
// api/api.go
type ServerConfig struct {
	LoadTemplates func() *django.Engine
}
type TemplatesEngine struct {
	engine *django.Engine
}
```

# Styling

- If you've selected tailwind, then no extra configuration is needed, start adding classes in any html.
- You can always use plain css (even with tailwind).

# Quick Tips

- **HTML Routes:** Render templates using handlers like the example above.
- **JSON Routes:** Prefix API endpoints with `/api/json`. The configuration ensures JSON responses even on errors.

For example, `/api/json/example` will always return a JSON response, whereas `/example` would render a template or custom HTML error pages.

# Advanced Usage

**You can also install any npm library and use it.**

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
    <!-- Optionally defer if needed eg. </script defer>...</script> -->
    <script src="/public/bundle/some-library.js"></script>
    ```

# Links to Documentation

- [Fiber](https://docs.gofiber.io)
- [Esbuild](https://esbuild.github.io)
- [TailwindCSS](https://tailwindcss.com)