package template

import (
	"embed"
)

//go:embed base/*
var base embed.FS

func GetBaseFiles() embed.FS {
	return base
}
