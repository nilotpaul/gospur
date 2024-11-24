package util

import (
	"fmt"
	"strings"

	"github.com/yosssi/gohtml"
)

const (
	basicHomeBodyExampleHTML = `
<body>
  <main class="container">
      <h1>{{ .Ctx.Title }}</h1>
      <img
        src="public/golang.jpg"
        class="rounded-md"
        height="500"
        width="500"
      />
      <p>{{ .Ctx.Desc }}</p>
  </main>
</body>`
	tailwindHomeBodyExampleHTML = `
  <body
    class="max-w-full bg-gray-300 dark:text-slate-200 text-black dark:bg-gray-950"
  >
    <!-- Container (MAX-WIDTH -> 48rem) -->
    <main class="max-w-3xl mx-auto">
      <div class="flex items-center gap-y-6 mt-4 flex-col justify-center">
        <h1 class="text-4xl my-4 text-blue-600 font-bold">
          {{ .Ctx.Title }}
        </h1>
        <img
          src="public/golang.jpg"
          class="rounded-md"
          height="500"
          width="500"
        />
        <p class="text-lg font-medium">{{ .Ctx.Desc }}</p>
      </div>
    </main>
  </body>  `

	basicErrorBodyExampleHTML = `
<body class="container">
  <h1>{{ .Ctx.FullError }}</h1>
</body>`
	tailwindErrorBodyExampleHTML = `
<body class="flex items-center justify-center">
    <h1 class="font-bold">{{ .Ctx.FullError }}</h1>
</body>`
)

func generatePageContent(page string, cfg StackConfig) []byte {
	var result string

	switch page {
	case "Home.html":
		result = processRawHomePageData(cfg)
	case "Error.html":
		result = processRawErrorPageData(cfg)
	default:
	}

	return []byte(gohtml.Format(result))
}

func processRawHomePageData(cfg StackConfig) string {
	homeHTML := fmt.Sprintf(`{{ define "Home" }}
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    %s
    <!-- For live reloading -->
    {{ if .IsDev }}
    <script src="http://localhost:35729/livereload.js"></script>
    {{ end }}
    %s

    <title>{{ .Ctx.Title }}</title>
  </head>
  %s
</html>
{{ end }}
`,
		generateHeadStyles(cfg),
		generateHeadScripts(cfg),
		generateHomeHTMLBody(cfg),
	)

	return homeHTML
}

func processRawErrorPageData(cfg StackConfig) string {
	errorHTML := fmt.Sprintf(`{{ define "Error" }}
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    %s
    <!-- For live reloading -->
    {{ if .IsDev }}
    <script src="http://localhost:35729/livereload.js"></script>
    {{ end }}

    <title>{{ .Ctx.Msg }}</title>
  </head>
  %s
</html>
{{ end }}
`,
		generateHeadStyles(cfg),
		generateErrorHTMLBody(cfg),
	)

	return errorHTML
}

func generateHomeHTMLBody(cfg StackConfig) string {
	if cfg.CssStrategy == "Tailwind" {
		return tailwindHomeBodyExampleHTML
	}
	return basicHomeBodyExampleHTML
}

func generateErrorHTMLBody(cfg StackConfig) string {
	if cfg.CssStrategy == "Tailwind" {
		return tailwindErrorBodyExampleHTML
	}
	return basicErrorBodyExampleHTML
}

func generateHeadScripts(cfg StackConfig) string {
	scripts := []string{"<!-- Bundled Javascript -->"}

	if contains(cfg.Extras, "HTMX") {
		scripts = append(scripts, `<script defer src="public/bundle/htmx.js"></script>`)
	}
	if cfg.UILibrary == "Preline" {
		scripts = append(scripts, `<script defer src="public/bundle/preline.js"></script>`)
	}
	if len(scripts) == 1 {
		return ""
	}

	return strings.Join(scripts, "\n")
}

func generateHeadStyles(_ StackConfig) string {
	styles := []string{
		"<!-- Styles -->",
		`<link rel="stylesheet" href="public/bundle/globals.css" />`,
	}

	return strings.Join(styles, "\n")
}
