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

var ProjectStructure = map[string]string{
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
