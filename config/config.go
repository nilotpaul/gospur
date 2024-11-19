package config

import "github.com/manifoldco/promptui"

var (
	ErrMsg     = promptui.Styler(promptui.FGRed)
	SuccessMsg = promptui.Styler(promptui.FGGreen)
	NormalMsg  = promptui.Styler(promptui.FGBlack)
)

var (
	WebFrameworkOpts = []string{
		"Echo",
	}
	UILibraryOpts = []string{
		"Preline (requires tailwind)",
	}
	ExtraOpts = []string{
		"Tailwind",
		"HTMX",
	}
)

var (
	ProjectBaseFiles = map[string]string{
		"config/env.go":          "base/env.go.tmpl",
		"web/styles/globals.css": "base/globals.css.tmpl",
		".gitignore":             "base/gitignore.tmpl",
		"Makefile":               "base/makefile.tmpl",
		"README.md":              "base/readme.md.tmpl",
		"esbuild.config.js":      "base/esbuild.config.js.tmpl",
		"package.json":           "base/package.json.tmpl",
		"tailwind.config.js":     "base/tailwind.config.js.tmpl",
		"main.go":                "base/main.go.tmpl",
	}

	ProjectPageFiles = map[string]string{
		"web/Home.html":  "page/error.html.echo.tmpl",
		"web/Error.html": "page/home.html.echo.tmpl",
	}

	ProjectAPIFiles = map[string]string{
		"api/api.go":     "api/api.go.echo.tmpl",
		"api/route.go":   "api/route.go.echo.tmpl",
		"api/handler.go": "api/handler.go.echo.tmpl",
	}
)
