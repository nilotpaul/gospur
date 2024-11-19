package template

import (
	_ "embed"
)

//go:embed base/main.go.tmpl
var base []byte

func GetMain() []byte {
	return base
}
