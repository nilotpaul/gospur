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

	// With tailwind3, tailwind.config.js is needed.
	mockStackCfg.CssStrategy = "Tailwind3"
	skip := skipProjectfiles("tailwind.config.js", mockStackCfg)
	a.False(skip)

	// With tailwind4, tailwind.config.js is not needed.
	mockStackCfg.CssStrategy = "Tailwind4"
	skip = skipProjectfiles("tailwind.config.js", mockStackCfg)
	a.True(skip)

	// With Vanilla CSS, tailwind.config.js is not needed.
	mockStackCfg.CssStrategy = "Vanilla"
	skip = skipProjectfiles("tailwind.config.js", mockStackCfg)
	a.True(skip)
}
