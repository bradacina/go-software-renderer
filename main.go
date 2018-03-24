package main

import (
	"github.com/bradacina/go-software-renderer/obj"
	"github.com/bradacina/go-software-renderer/renderer"
	"github.com/bradacina/go-software-renderer/tga"
)

func main() {
	image := renderer.NewBuffer(1000, 1000, true)
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

func drawObj(o *obj.Obj, gb *renderer.Buffer) {

	color := renderer.ARGB{Alpha: 255}
	a, b, c := renderer.Vertex{}, renderer.Vertex{}, renderer.Vertex{}

	light := &renderer.Vector{renderer.Vertex{0, 0, 1}}

	renderer.Normalize(light)

	for _, f := range o.Faces {

		objVertexToRenderVertex(o.Vertices[f[0]-1], &a)
		objVertexToRenderVertex(o.Vertices[f[1]-1], &b)
		objVertexToRenderVertex(o.Vertices[f[2]-1], &c)
		color.Alpha = 255
		renderer.GrayShade(&a, &b, &c, light, &color)

		if color.Alpha == 0 {
			continue
		}

		gb.Triangle(&a, &b, &c, &color)
	}
}

func objVertexToPoint(ov *obj.Vertex, rv *renderer.Point, halfWidth, halfHeight float64) {
	rv.X = int((ov.X + 1) * halfHeight)
	rv.Y = int((ov.Y + 1) * halfWidth)
}

func objVertexToRenderVertex(ov *obj.Vertex, rv *renderer.Vertex) {
	rv.X = ov.X
	rv.Y = ov.Y
	rv.Z = ov.Z
}
