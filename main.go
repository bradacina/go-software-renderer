package main

import (
	"github.com/bradacina/go-software-renderer/obj"
	"github.com/bradacina/go-software-renderer/renderer"
	"github.com/bradacina/go-software-renderer/tga"
)

func main() {

	image := renderer.NewBuffer(1024, 1024, true)

	o := obj.Load("african_head.obj")

	texture := tgaToBuffer("xxx.tga")
	drawObj(o, texture, image)
	tga.Save(image.Width, image.Height, image.Data, "test.tga")
}

func tgaToBuffer(filename string) *renderer.Buffer {
	width, height, _, data := tga.Load(filename)
	img := renderer.NewBuffer(width, height, false)

	var color renderer.RGBA

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			idx := (j*width + i) * 4
			color.Blue = (*data)[idx]
			color.Green = (*data)[idx+1]
			color.Red = (*data)[idx+2]
			color.Alpha = (*data)[idx+3]

			img.Draw(i, j, &color)
		}
	}

	return img
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

func drawObj(o *obj.Obj, texture *renderer.Buffer, gb *renderer.Buffer) {

	// 3d model triangle coords
	a, b, c := renderer.Vertex{}, renderer.Vertex{}, renderer.Vertex{}

	// texture coords
	at, bt, ct := renderer.Point{}, renderer.Point{}, renderer.Point{}

	light := &renderer.Vector{renderer.Vertex{0, 0, 1}}

	renderer.Normalize(light)

	texHeight := float64(texture.Height)
	texWidth := float64(texture.Width)

	for idx, f := range o.Faces {
		renderer.DebugFaceIdx = idx
		v1Idx := f.VertexIndex[0] - 1
		v2Idx := f.VertexIndex[1] - 1
		v3Idx := f.VertexIndex[2] - 1
		objVertexToRenderVertex(o.Vertices[v1Idx], &a)
		objVertexToRenderVertex(o.Vertices[v2Idx], &b)
		objVertexToRenderVertex(o.Vertices[v3Idx], &c)

		vt1Idx := f.VertexTextureIndex[0] - 1
		vt2Idx := f.VertexTextureIndex[1] - 1
		vt3Idx := f.VertexTextureIndex[2] - 1

		objTexVertexToPoint(o.VerticesTexture[vt1Idx], &at, texWidth, texHeight)
		objTexVertexToPoint(o.VerticesTexture[vt2Idx], &bt, texWidth, texHeight)
		objTexVertexToPoint(o.VerticesTexture[vt3Idx], &ct, texWidth, texHeight)

		gb.TexturedTriangle(&a, &b, &c, &at, &bt, &ct, texture, light)
	}
}

func objTexVertexToPoint(ov *obj.Vertex, rv *renderer.Point, width, height float64) {
	rv.X = int(ov.X * width)
	rv.Y = int(ov.Y * height)
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
