package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipProjectfiles(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	mockStackCfg := StackConfig{
		WebFramework: "Echo",
	}

	// With tailwind (v3), tailwind.config.js is needed.
	mockStackCfg.CssStrategy = "Tailwind (v3)"
	skip := skipProjectfiles("tailwind.config.js", mockStackCfg)
	a.False(skip)

	// With tailwind (v4), tailwind.config.js is not needed.
	mockStackCfg.CssStrategy = "Tailwind (v4)"
	skip = skipProjectfiles("tailwind.config.js", mockStackCfg)
	a.True(skip)

	// With Vanilla CSS, tailwind.config.js is not needed.
	mockStackCfg.CssStrategy = "Vanilla"
	skip = skipProjectfiles("tailwind.config.js", mockStackCfg)
	a.True(skip)
}
