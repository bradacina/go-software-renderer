package main

import (
	"log"
	"math"

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
	a, b, c := renderer.AfineVertex{}, renderer.AfineVertex{}, renderer.AfineVertex{}

	// texture coords
	at, bt, ct := renderer.Point{}, renderer.Point{}, renderer.Point{}

	light := &renderer.Vector{renderer.Vertex{0, 50, 50}}

	renderer.Normalize(light)

	texHeight := float64(texture.Height)
	texWidth := float64(texture.Width)

	cameraLocation := renderer.Vertex{60.0, -0.4, 0.0}
	cameraDirection := renderer.Vertex{0, 0, 0}
	cameraUp := renderer.Vector{renderer.Vertex{0, 1.0, 0}}

	camera := renderer.NewCamera()
	camera.LookAt(&cameraLocation, &cameraDirection, &cameraUp)

	projectionMatrix := renderer.Ortographic(-1, 1, -1, 1, -1, 1)
	viewPortMatrix := renderer.ViewPort(0, 0, 1024, 1024)

	var temp renderer.Mat4x4
	var viewPipeline renderer.Mat4x4
	renderer.Mul4x4(&camera.ModelView, projectionMatrix, &temp)

	//renderer.Mul4x4(&temp, viewPortMatrix, &viewPipeline)

	log.Println("ViewPort", viewPortMatrix)
	log.Println("Projection", projectionMatrix)

	log.Println("ViewModel x Projection", temp)
	log.Println("All", viewPipeline)

	// viewPipeline = temp
	viewPipeline = temp

	var transA, transB, transC renderer.AfineVertex
	var postA, postB, postC renderer.Vertex

	maxX, maxY, maxZ := -math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64
	minX, minY, minZ := math.MaxFloat64, math.MaxFloat64, math.MaxFloat64

	for _, f := range o.Faces {
		v1Idx := f.VertexIndex[0] - 1
		v2Idx := f.VertexIndex[1] - 1
		v3Idx := f.VertexIndex[2] - 1
		objVertexToRenderAfineVertex(o.Vertices[v1Idx], &a)
		objVertexToRenderAfineVertex(o.Vertices[v2Idx], &b)
		objVertexToRenderAfineVertex(o.Vertices[v3Idx], &c)

		vt1Idx := f.VertexTextureIndex[0] - 1
		vt2Idx := f.VertexTextureIndex[1] - 1
		vt3Idx := f.VertexTextureIndex[2] - 1

		objTexVertexToPoint(o.VerticesTexture[vt1Idx], &at, texWidth, texHeight)
		objTexVertexToPoint(o.VerticesTexture[vt2Idx], &bt, texWidth, texHeight)
		objTexVertexToPoint(o.VerticesTexture[vt3Idx], &ct, texWidth, texHeight)

		renderer.Mul4x4WithAfineVertex(&viewPipeline, &a, &transA)
		renderer.Mul4x4WithAfineVertex(&viewPipeline, &b, &transB)
		renderer.Mul4x4WithAfineVertex(&viewPipeline, &c, &transC)

		renderer.AfineVertexToVertex(&transA, &postA)
		renderer.AfineVertexToVertex(&transB, &postB)
		renderer.AfineVertexToVertex(&transC, &postC)

		bbox(&maxX, &minX, &maxY, &minY, &maxZ, &minZ, &postA, &postB, &postC)

		gb.TexturedTriangle(&postA, &postB, &postC, &at, &bt, &ct, texture, light)
	}

	log.Println("minX", minX, "maxX", maxX)
	log.Println("minY", minY, "maxY", maxY)
	log.Println("minZ", minZ, "maxZ", maxZ)
}

func bbox(maxX, minX, maxY, minY, maxZ, minZ *float64, a, b, c *renderer.Vertex) {
	*maxX = math.Max(*maxX, math.Max(a.X, math.Max(b.X, c.X)))
	*minX = math.Min(*minX, math.Min(a.X, math.Min(b.X, c.X)))
	*maxY = math.Max(*maxY, math.Max(a.Y, math.Max(b.Y, c.Y)))
	*minY = math.Min(*minY, math.Min(a.Y, math.Min(b.Y, c.Y)))
	*maxZ = math.Max(*maxZ, math.Max(a.Z, math.Max(b.Z, c.Z)))
	*minZ = math.Min(*minZ, math.Min(a.Z, math.Min(b.Z, c.Z)))
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

func objVertexToRenderAfineVertex(ov *obj.Vertex, rv *renderer.AfineVertex) {
	rv.X = ov.X
	rv.Y = ov.Y
	rv.Z = ov.Z
	rv.W = 1
}
