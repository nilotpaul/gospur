package util

import (
	"fmt"
	"strings"
)

const (
	basicHomeBodyExampleHTML = `
<body>
  <main>
    <div>
      <h1>{{ index . "Ctx" "Title" }}</h1>
      <img
        src="public/golang.jpg"
        class="rounded-md"
        height="500"
        width="500"
      />
      <p>{{ index . "Ctx" "Desc" }}</p>
    </div>

    <div>
      <button>Light</button>
    </div>
  </main>
</body>`
	tailwindHomeBodyExampleHTML = `
  <body
    class="max-w-full min-h-screen bg-gray-300 dark:text-slate-200 text-black dark:bg-gray-950"
  >
    <!-- Container (MAX-WIDTH -> 48rem) -->
    <main class="max-w-3xl mx-auto">
      <div class="flex items-center gap-y-6 mt-4 flex-col justify-center">
        <h1 class="text-4xl mt-6 mb-0 text-blue-600 font-bold">
          {{ index . "Ctx" "Title" }}
        </h1>
        <img
          src="public/golang.jpg"
          class="rounded-md"
          height="500"
          width="500"
        />
        <p class="text-lg font-medium">{{ index . "Ctx" "Desc" }}</p>
      </div>

      <div class="flex items-center mt-4 justify-center gap-4">
        <button
          type="button"
          class="hs-dark-mode hs-dark-mode-active:hidden inline-flex items-center gap-x-2 py-2 px-3 bg-black/80 rounded-full text-sm text-white hover:bg-black/70 focus:outline-none focus:bg-black/70"
          data-hs-theme-click-value="dark"
        >
          <svg
            class="shrink-0 size-4"
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z"></path>
          </svg>
          Dark
        </button>
        <button
          type="button"
          class="hs-dark-mode hs-dark-mode-active:inline-flex hidden items-center gap-x-2 py-2 px-3 bg-white/10 rounded-full text-sm text-white hover:bg-white/20 focus:outline-none focus:bg-white/20"
          data-hs-theme-click-value="light"
        >
          <svg
            class="shrink-0 size-4"
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <circle cx="12" cy="12" r="4"></circle>
            <path d="M12 2v2"></path>
            <path d="M12 20v2"></path>
            <path d="m4.93 4.93 1.41 1.41"></path>
            <path d="m17.66 17.66 1.41 1.41"></path>
            <path d="M2 12h2"></path>
            <path d="M20 12h2"></path>
            <path d="m6.34 17.66-1.41 1.41"></path>
            <path d="m19.07 4.93-1.41 1.41"></path>
          </svg>
          Light
        </button>
      </div>
    </main>
  </body>  `

	basicErrorBodyExampleHTML = `
<body>
  <h1>{{ index . "Ctx" "FullError" }}</h1>
</body>`
	tailwindErrorBodyExampleHTML = `
<body class="flex items-center justify-center">
    <h1 class="font-bold">{{ index . "Ctx" "FullError" }}</h1>
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

	return []byte(result)
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
    {{ if index .IsDev }}
    <script src="http://localhost:35729/livereload.js"></script>
    {{ end }}
    %s

    <title>{{ index . "Ctx" "Title" }}</title>
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

    <title>{{ index . "Ctx" "Msg" }}</title>
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
	if cfg.CssStrategy == "Vanilla CSS" {
		return addMultilineIndentation(basicHomeBodyExampleHTML, 1)
	}
	if cfg.CssStrategy == "Tailwind" {
		return addMultilineIndentation(tailwindHomeBodyExampleHTML, 1)
	}
	return addMultilineIndentation(basicHomeBodyExampleHTML, 1)
}

func generateErrorHTMLBody(cfg StackConfig) string {
	if cfg.CssStrategy == "Vanilla CSS" {
		return addMultilineIndentation(basicErrorBodyExampleHTML, 1)
	}
	if cfg.CssStrategy == "Tailwind" {
		return addMultilineIndentation(tailwindErrorBodyExampleHTML, 1)
	}
	return addMultilineIndentation(basicErrorBodyExampleHTML, 1)
}

func generateHeadScripts(cfg StackConfig) string {
	scripts := []string{"<!-- Bundled Javascript -->"}

	if contains(cfg.Extras, "HTMX") {
		scripts = append(scripts, "\t\t"+`<script defer src="public/bundle/htmx.org/dist/htmx.js"></script>`)
	}
	if cfg.UILibrary == "Preline" {
		scripts = append(scripts, "\t\t"+`<script defer src="public/bundle/preline/preline.js"></script>`)
	}
	if len(scripts) == 1 {
		return ""
	}

	return strings.Join(scripts, "\n")
}

func generateHeadStyles(cfg StackConfig) string {
	styles := []string{"<!-- Styles -->"}

	if cfg.CssStrategy == "Tailwind" {
		styles = append(styles, "\t\t"+`<link rel="stylesheet" href="public/bundle/styles.css" />`)
	}
	if len(styles) == 1 {
		return ""
	}

	return strings.Join(styles, "\n")
}
