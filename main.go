package main

import (
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
	image.FlipVertically()
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

	// vertex normals
	an, bn, cn := renderer.Vector{}, renderer.Vector{}, renderer.Vector{}

	// texture coords
	at, bt, ct := renderer.Point{}, renderer.Point{}, renderer.Point{}

	light := &renderer.Vector{renderer.Vertex{0.0, -1.0, 1.0}}

	renderer.Normalize(light)

	texHeight := float64(texture.Height)
	texWidth := float64(texture.Width)

	cameraLocation := renderer.Vertex{-3.0, -3.0, -10.0}
	cameraDirection := renderer.Vertex{0, 0, 0}
	cameraUp := renderer.Vector{renderer.Vertex{0, 1.0, 0}}

	camera := renderer.NewCamera()
	camera.LookAt(&cameraLocation, &cameraDirection, &cameraUp)

	projectionMatrix := renderer.Ortographic(-1, 1, -1, 1, 0, 100)
	viewPortMatrix := renderer.ViewPort(0, 0, 1024, 1024)

	var temp renderer.Mat4x4
	var viewPipeline renderer.Mat4x4

	renderer.Mul4x4(projectionMatrix, &camera.ModelView, &temp)
	renderer.Mul4x4(viewPortMatrix, &temp, &viewPipeline)

	var transA, transB, transC renderer.AfineVertex
	var postA, postB, postC renderer.Vertex
	var postAN, postBN, postCN renderer.Vector

	for _, f := range o.Faces {
		// vertex shader
		v1Idx := f.VertexIndex[0] - 1
		v2Idx := f.VertexIndex[1] - 1
		v3Idx := f.VertexIndex[2] - 1
		objVertexToRenderAfineVertex(o.Vertices[v1Idx], &a)
		objVertexToRenderAfineVertex(o.Vertices[v2Idx], &b)
		objVertexToRenderAfineVertex(o.Vertices[v3Idx], &c)

		renderer.Mul4x4WithAfineVertex(&viewPipeline, &a, &transA)
		renderer.Mul4x4WithAfineVertex(&viewPipeline, &b, &transB)
		renderer.Mul4x4WithAfineVertex(&viewPipeline, &c, &transC)

		renderer.AfineVertexToVertex(&transA, &postA)
		renderer.AfineVertexToVertex(&transB, &postB)
		renderer.AfineVertexToVertex(&transC, &postC)

		vt1Idx := f.VertexTextureIndex[0] - 1
		vt2Idx := f.VertexTextureIndex[1] - 1
		vt3Idx := f.VertexTextureIndex[2] - 1

		objTexVertexToPoint(o.VerticesTexture[vt1Idx], &at, texWidth, texHeight)
		objTexVertexToPoint(o.VerticesTexture[vt2Idx], &bt, texWidth, texHeight)
		objTexVertexToPoint(o.VerticesTexture[vt3Idx], &ct, texWidth, texHeight)

		// vertex normals
		vn1Idx := f.VertexNormalIndex[0] - 1
		vn2Idx := f.VertexNormalIndex[1] - 1
		vn3Idx := f.VertexNormalIndex[2] - 1

		objVertexToRenderVector(o.VerticesNormal[vn1Idx], &an)
		objVertexToRenderVector(o.VerticesNormal[vn2Idx], &bn)
		objVertexToRenderVector(o.VerticesNormal[vn3Idx], &cn)

		renderer.Mul3x3WithVector(&camera.NormalMatrix, &an, &postAN)
		renderer.Mul3x3WithVector(&camera.NormalMatrix, &bn, &postBN)
		renderer.Mul3x3WithVector(&camera.NormalMatrix, &cn, &postCN)

		// fragment shader
		//textureShader := renderer.NewTextureShader(&postA, &postB, &postC, &at, &bt, &ct, texture, light)
		//gb.TexturedTriangle(&postA, &postB, &postC, textureShader)

		gts := renderer.NewGouraudTextureShader(
			&postA, &postB, &postC, &postAN, &postBN, &postCN, light, &at, &bt, &ct, texture)
		gb.TexturedTriangle(&postA, &postB, &postC, gts)
	}
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

func objVertexToRenderVector(ov *obj.Vertex, rv *renderer.Vector) {
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
