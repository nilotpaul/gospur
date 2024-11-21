package config

import "github.com/manifoldco/promptui"

// For adding styles to console output.
var (
	ErrMsg     = promptui.Styler(promptui.FGRed)
	SuccessMsg = promptui.Styler(promptui.FGGreen, promptui.FGBold)
	NormalMsg  = promptui.Styler(promptui.FGWhite)
	FaintMsg   = promptui.Styler(promptui.FGFaint)
)

// UILibrary represents an UI Library and `DependsOn`
// which means it can depend on any chosen CSS Strategy (framework).
type UILibrary struct {
	// Name of the UI Library
	Name string

	// An UI Library can depend on any chosen CSS Strategy.
	// If it's independent, `DependsOn` should be an empty string.
	DependsOn string
}

// Prompt options.
var (
	WebFrameworkOpts = []string{
		"Echo",
	}
	ExtraOpts = []string{
		"HTMX",
	}

	CssStrategyOpts = []string{
		"Tailwind",
	}
	UILibraryOpts = map[string][]string{
		"Preline": {"Tailwind"},
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
		"main.go":                "base/main.go.tmpl",
	}

	ProjectPageFiles = map[string]string{
		"web/Home.html":  "page/home.html.echo.tmpl",
		"web/Error.html": "page/error.html.echo.tmpl",
	}

	ProjectAPIFiles = map[string]string{
		"api/api.go":     "api/api.go.echo.tmpl",
		"api/route.go":   "api/route.go.echo.tmpl",
		"api/handler.go": "api/handler.go.echo.tmpl",
	}
)
