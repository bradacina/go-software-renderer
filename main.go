package main

import (
	"log"
	"math/rand"

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

func min(a, b float64) float64 {
	if a > b {
		return b
	}

	return a
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}

	return b
}

func drawObj(o *obj.Obj, b *renderer.Buffer) {

	color := renderer.ARGB{Alpha: 255}
	vertices := [3]renderer.Vertex{}
	halfHeight := float64(b.Height-1) / 2
	halfWidth := float64(b.Width-1) / 2

	for _, f := range o.Faces {

		color.Red = byte(rand.Intn(256))
		color.Green = byte(rand.Intn(256))
		color.Blue = byte(rand.Intn(256))

		objVertexToRendererVertex(o.Vertices[f[0]-1], &vertices[0], halfWidth, halfHeight)
		objVertexToRendererVertex(o.Vertices[f[1]-1], &vertices[1], halfWidth, halfHeight)
		objVertexToRendererVertex(o.Vertices[f[2]-1], &vertices[2], halfWidth, halfHeight)

		b.Triangle(&vertices, &color)
	}
}

func objVertexToRendererVertex(ov *obj.Vertex, rv *renderer.Vertex, halfWidth, halfHeight float64) {
	rv.X = int((ov.X + 1) * halfHeight)
	rv.Y = int((ov.Y + 1) * halfWidth)

	log.Println(rv)
}
