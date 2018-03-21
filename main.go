package main

import (
	"github.com/bradacina/go-software-renderer/renderer"
	"github.com/bradacina/go-software-renderer/tga"
)

func main() {
	image := renderer.NewBuffer(1000, 1000)

	image.DrawLine(0, 0, 200, 999, &renderer.ARGB{255, 255, 255, 255})
	image.DrawLine(0, 0, 999, 200, &renderer.ARGB{255, 255, 255, 255})
	image.DrawLine(0, 0, 999, 999, &renderer.ARGB{255, 255, 255, 255})
	tga.Save(image.Width, image.Height, image.Data, "test.tga")
}
