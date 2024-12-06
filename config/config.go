package config

import "github.com/manifoldco/promptui"

// For adding styles to console output.
var (
	ErrMsg     = promptui.Styler(promptui.FGRed)
	SuccessMsg = promptui.Styler(promptui.FGGreen, promptui.FGBold)
	NormalMsg  = promptui.Styler(promptui.FGWhite)
	FaintMsg   = promptui.Styler(promptui.FGFaint)
)

type ProjectCtx map[string]any

// Prompt options.
var (
	WebFrameworkOpts = []string{
		"Echo",
	}
	ExtraOpts = []string{
		"HTMX",
		"Dockerfile",
	}

	CssStrategyOpts = []string{
		"Tailwind",
		"Vanilla CSS",
	}
	UILibraryOpts = map[string][]string{
		"Preline": {"Tailwind"},
		"DaisyUI": {"Tailwind"},
	}
)

// Project file structure
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
		"Dockerfile":             "base/.dockerfile.tmpl",
		".dockerignore":          "base/.dockerignore.tmpl",
		"main.go":                "base/main.go.tmpl",
	}

	// Template path is not required anymore for pages.
	// We're processing these as raw files.
	ProjectPageFiles = map[string]string{
		"web/Home.html":  "",
		"web/Error.html": "",
	}

	ProjectAPIFiles = map[string]string{
		"api/api.go":     "api/api.go.echo.tmpl",
		"api/route.go":   "api/route.go.echo.tmpl",
		"api/handler.go": "api/handler.go.echo.tmpl",
	}
)
