package template

import (
	"embed"
)

//go:embed base/*
var base embed.FS

//go:embed page/*
var page embed.FS

//go:embed api/*
var api embed.FS

//go:embed public/golang.jpg
var img []byte

func GetBaseFiles() embed.FS {
	return base
}

func GetPageFiles() embed.FS {
	return page
}

func GetAPIFiles() embed.FS {
	return api
}

func GetGolangImage() []byte {
	return img
}
