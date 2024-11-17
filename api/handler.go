package api

import (
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/nilotpaul/go-echo-htmx/config"
)

func HTTPErrorHandler(err error, c echo.Context) {
	var (
		status  int    = 500
		msg     string = "Internal Server Error"
		fullErr string = "Something Went Wrong"
	)
	if he, ok := err.(*echo.HTTPError); ok {
		status = he.Code
		msg = regexp.MustCompile(`message=([^,]+)`).FindStringSubmatch(he.Error())[1]
		fullErr = he.Error()
	}

	// If the path is prefixed with `/api/json`, send a JSON Response Back.
	// Otherwise, render a Error HTML Page.
	if strings.HasPrefix(c.Request().URL.Path, "/api/json") {
		c.JSON(status, map[string]any{"status": status, "error": msg})
	} else {
		c.Render(http.StatusOK, "Error", map[string]any{
			"IsDEV":     os.Getenv("ENVIRONMENT") == string(config.DEV),
			"FullError": fullErr,
			"Msg":       msg,
		})
	}
}

func handleGetHome(c echo.Context) error {
	return c.Render(http.StatusOK, "Home", map[string]any{
		"IsDEV": os.Getenv("ENVIRONMENT") == string(config.DEV),
		"Title": "Go + Echo + HTMX",
		"Desc":  "Best for building Full-Stack Applications with minimal JavaScript",
	})
}
