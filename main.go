package main

import (
	"github.com/bradacina/go-software-renderer/obj"
	"github.com/bradacina/go-software-renderer/renderer"
	"github.com/bradacina/go-software-renderer/tga"
)

func main() {
	image := renderer.NewBuffer(1000, 1000)
	o := obj.Load("african_head.obj")

	drawObj(o, image)
	tga.Save(image.Width, image.Height, image.Data, "test.tga")
}

func min(a, b float32) float32 {
	if a > b {
		return b
	}

	return a
}

func max(a, b float32) float32 {
	if a > b {
		return a
	}

	return b
}

func drawObj(o *obj.Obj, b *renderer.Buffer) {

	color := renderer.ARGB{255, 255, 255, 255}
	height := float32(b.Height-1) / 2
	width := float32(b.Width-1) / 2

	for _, f := range o.Faces {
		for i := 0; i < 3; i++ {
			v1 := o.Vertices[f[i]-1]
			v2 := o.Vertices[f[(i+1)%3]-1]

			x1 := int((v1.X + 1) * height)
			y1 := int((v1.Y + 1) * width)

			x2 := int((v2.X + 1) * height)
			y2 := int((v2.Y + 1) * width)

			b.DrawLine(x1, y1, x2, y2, &color)
		}
	}
}
