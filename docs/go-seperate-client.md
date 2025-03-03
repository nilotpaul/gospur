# Setting Up Your Frontend

You have two options for integrating your frontend with the Go backend:

1. Keep frontend inside this repo

- Create your frontend project inside the `web` folder.
- Do not modify `web/dist` directly - Go requires it to exist.
- Configure your frontend to output static files to `web/dist`, replacing the default one.
- If your build output directory is not `dist`, update it in `build_prod.go`.

2. Keep frontend in a separate repo

- Build the frontend in a CI pipeline (examples given below).
- Merge the generated `dist` folder into `web/dist`.

## Examples
- [Basic Github Actions](#)
- [Docker Github Actions](#)
- [Dockerfile](#)

> **TODO: will update later.**

# Serving a Separate Frontend with Go

Frontend frameworks like React, Vue, Svelte, etc. generate static assets (HTML, CSS, JS), which can be served directly from a Go backend. This simplifies deployment and avoids CORS issues.

# How It Works

- The frontend build (`build`/`dist` folder) is embedded in Go using `embed.FS`.
- If a requested file exists, it's served normally.
- If not, Go serves `index.html` (for SPAs) or a dedicated error page if available (for SSGs).

# SPA vs. SSG Behavior

- **SPA (Single Page App)**: Always fallbacks to `index.html`, and routing/errors are handled in JavaScript.
- **SSG (Static Site Generation)**: If a `404.html` or similar page exists, serve that instead of `index.html`.

# Configuring Fallback Pages (if needed)

Some SSG frameworks may generate a `404.html` for handling missing pages. To serve it properly:

**Echo Example**
```go
// Todo: will fix this later. Echo doesn't have an option to specify a fallback.
e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
	Root:         "web/dist",
	Index:        "index.html",
	Filesystem:   http.FS(web),
}))
```

**Fiber Example**
```go
app.Use(filesystem.New(filesystem.Config{
	Root:         http.FS(subFS),
	Browse:       false,
	Index:        "index.html",
	NotFoundFile: "404.html", // Change this
}))
```

**Chi Example**
```go
mux.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")

	// Check if the requested file exists
	_, err := subFS.Open(path)
	if err != nil {
		// If not found, serve fallback page.
		http.ServeFileFS(w, r, subFS, "404.html") // Change this
		return
	}

    fs.ServeHTTP(w, r)
}))
```