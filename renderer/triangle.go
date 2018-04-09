package renderer

import (
	"math"
)

var DebugFaceIdx int

var isTrianglevertex bool

var debugFaces = []int{}

func isDebugArea(p *Point) bool {
	if p.Y >= 631 && p.Y <= 637 &&
		p.X >= 510 && p.X <= 514 {
		return true
	}

	return false
}

func isDebugFace() bool {
	for _, v := range debugFaces {
		if v == DebugFaceIdx {
			return true
		}
	}

	return false
}

// TexturedTriangle draws a triangle using a shader
func (buf *Buffer) TexturedTriangle(
	a, b, c *Vertex,
	shader Shader) {

	// backface cull
	if shader.ShouldIgnore() {
		return
	}

	// bring the triangle into 2D (Buffer) space
	ap := buf.Vertex2Point(a)
	bp := buf.Vertex2Point(b)
	cp := buf.Vertex2Point(c)

	topLeft, bottomRight := bBox(ap, bp, cp)

	tempVertex := Point{}

	vertices := [3]*Point{ap, bp, cp}

	color := RGBA{}

	// fill in the triangle using the pixels in the bounding box
	for col := topLeft.X; col <= bottomRight.X; col++ {
		for row := topLeft.Y; row <= bottomRight.Y; row++ {

			// test if pixel is within the Buffer
			if col < 0 || col >= buf.Width ||
				row < 0 || row >= buf.Height {
				continue
			}

			tempVertex.X = col
			tempVertex.Y = row

			u, v, w := barycentric(&tempVertex, &vertices)

			// test if pixel is within the triangle
			if u < 0 || v < 0 || w < 0 {
				continue
			}

			// depth buffer test
			if buf.DepthBuf != nil {
				z := u*a.Z + v*b.Z + w*c.Z
				if buf.DepthBuf[row*buf.Width+col] > z {
					continue
				}

				buf.DepthBuf[row*buf.Width+col] = z
			}

			shader.ShadeFragment(u, v, w, &color)

			if color.Alpha <= 0 {
				continue
			}

			if isDebugFace() {
				buf.Draw(col, row, &ColorRed)
				continue
			}

			buf.Draw(col, row, &color)
		}
	}
}

// TriangleMesh draws the outer shape of a triangle
func (buf *Buffer) TriangleMesh(a, b, c *Point, color *RGBA) {
	buf.DrawLine(a.X, a.Y, b.X, b.Y, color)
	buf.DrawLine(a.X, a.Y, c.X, c.Y, color)
	buf.DrawLine(c.X, c.Y, b.X, b.Y, color)
}

//Triangle draws a triangle from 3D space onto the Buffer with specified color
func (buf *Buffer) Triangle(a, b, c *Vertex, color *RGBA) {

	// bring the triangle into 2D (Buffer) space
	ap := buf.Vertex2Point(a)
	bp := buf.Vertex2Point(b)
	cp := buf.Vertex2Point(c)

	topLeft, bottomRight := bBox(ap, bp, cp)

	tempVertex := Point{}

	vertices := [3]*Point{ap, bp, cp}

	// fix in the triangle using the pixels in the bounding box
	for col := topLeft.X; col <= bottomRight.X; col++ {
		for row := topLeft.Y; row <= bottomRight.Y; row++ {

			// test if pixel is within the Buffer
			if col < 0 || col >= buf.Width ||
				row < 0 || row >= buf.Height {
				continue
			}

			tempVertex.X = col
			tempVertex.Y = row

			u, v, w := barycentric(&tempVertex, &vertices)

			// test if pixel is within the triangle
			if u < 0 || v < 0 || w < 0 {
				continue
			}

			// depth buffer test
			if buf.DepthBuf != nil {
				z := u*a.Z + v*b.Z + w*c.Z
				if buf.DepthBuf[row*buf.Width+col] > z {
					continue
				}

				buf.DepthBuf[row*buf.Width+col] = z
			}

			buf.Draw(col, row, color)
		}
	}
}

// Vertex2Point creates a Point from a Vertex
func (buf *Buffer) Vertex2Point(v *Vertex) *Point {
	return &Point{int(v.X), int(v.Y)}
	//return &Point{int(v.X * float64(buf.Width)), int(v.Y * float64(buf.Height))}
	// return &Point{int(math.RoundToEven((v.X + 1) * buf.halfWidth)),
	// 	int(math.RoundToEven((v.Y + 1) * buf.halfHeight))}
}

func barycentric(p *Point, vertices *[3]*Point) (u, v, w float64) {
	// Compute vectors
	v0 := VectorFromPoints(vertices[0], vertices[2])
	v1 := VectorFromPoints(vertices[0], vertices[1])
	v2 := VectorFromPoints(vertices[0], p)

	// Compute dot products
	dot00 := dot(v0, v0)
	dot01 := dot(v0, v1)
	dot02 := dot(v0, v2)
	dot11 := dot(v1, v1)
	dot12 := dot(v1, v2)

	// Compute barycentric coordinates
	invDenom := 1 / float64(dot00*dot11-dot01*dot01)
	if math.IsNaN(invDenom) || math.IsInf(invDenom, 0) {
		return -1, -1, -1
	}

	w = float64(dot11*dot02-dot01*dot12) * invDenom
	v = float64(dot00*dot12-dot01*dot02) * invDenom
	u = 1 - w - v
	return u, v, w
}

func dot(v1, v2 *Point) int {
	return v1.X*v2.X + v1.Y*v2.Y
}

// VectorFromPoints creates a Vector from 2 Points
func VectorFromPoints(tail, head *Point) *Point {
	return &Point{X: head.X - tail.X, Y: head.Y - tail.Y}
}

// Calculates a bounding box for a screen triangle
func bBox(a, b, c *Point) (*Point, *Point) {
	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := math.MinInt32, math.MinInt32

	minX = min(min(min(minX, a.X), b.X), c.X)
	minY = min(min(min(minY, a.Y), b.Y), c.Y)
	maxX = max(max(max(maxX, a.X), b.X), c.X)
	maxY = max(max(max(maxY, a.Y), b.Y), c.Y)

	return &Point{X: minX, Y: minY}, &Point{X: maxX, Y: maxY}
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
