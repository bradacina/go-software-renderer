package renderer

import (
	"log"
	"math"
)

func (buf *Buffer) TexturedTriangle(
	a, b, c *Vertex,
	at, bt, ct *Point,
	tex *Buffer,
	light *Vector) {

	// bring the triangle into 2D (Buffer) space
	ap := buf.Vertex2Point(a)
	bp := buf.Vertex2Point(b)
	cp := buf.Vertex2Point(c)

	topLeft, bottomRight := bBox(ap, bp, cp)

	tempVertex := Point{}

	vertices := [3]*Point{ap, bp, cp}

	// fill in the triangle using the pixels in the bounding box
	for i := topLeft.X; i <= bottomRight.X; i++ {
		for j := topLeft.Y; j <= bottomRight.Y; j++ {

			// test if pixel is within the Buffer
			if i < 0 || i >= buf.Height ||
				j < 0 || j >= buf.Width {
				continue
			}

			tempVertex.X = i
			tempVertex.Y = j

			u, v, w := barycentric(&tempVertex, &vertices)

			// test if pixel is within the triangle
			if u < 0 || v < 0 || w < 0 {
				continue
			}

			// depth buffer test
			if buf.DepthBuf != nil {
				z := u*a.Z + v*b.Z + w*c.Z
				if buf.DepthBuf[i*buf.Width+j] > z {
					continue
				}

				buf.DepthBuf[i*buf.Width+j] = z
			}

			color := resolveColor(at, bt, ct, tex, u, v, w)

			if color.Alpha <= 0 {
				log.Println("tried to read from outside texture buffer")
				log.Println(at, bt, ct, u, v, w)
				continue
			}

			//ColorIntensity(a, b, c, light, color)

			buf.Draw(i, j, color)
		}
	}
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
	for i := topLeft.X; i <= bottomRight.X; i++ {
		for j := topLeft.Y; j <= bottomRight.Y; j++ {

			// test if pixel is within the Buffer
			if i < 0 || i >= buf.Width ||
				j < 0 || j >= buf.Height {
				continue
			}

			tempVertex.X = i
			tempVertex.Y = j

			u, v, w := barycentric(&tempVertex, &vertices)

			// test if pixel is within the triangle
			if u < 0 || v < 0 || w < 0 {
				continue
			}

			// depth buffer test
			if buf.DepthBuf != nil {
				z := u*a.Z + v*b.Z + w*c.Z
				if buf.DepthBuf[i*buf.Width+j] > z {
					continue
				}

				buf.DepthBuf[i*buf.Width+j] = z
			}

			buf.Draw(j, i, color)
		}
	}
}

// Vertex2Point creates a Point from a Vertex
func (buf *Buffer) Vertex2Point(v *Vertex) *Point {
	return &Point{int((v.X + 1) * buf.halfWidth), int((v.Y + 1) * buf.halfHeight)}
}

func resolveColor(at, bt, ct *Point, tex *Buffer, u, v, w float64) *RGBA {

	p := Vector{Vertex{
		X: float64(at.X)*u + float64(bt.X)*v + float64(ct.X)*w,
		Y: float64(at.Y)*u + float64(bt.Y)*v + float64(ct.Y)*w}}

	color := tex.Read(int(p.X), int(p.Y))

	return color
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

	u = float64(dot11*dot02-dot01*dot12) * invDenom
	v = float64(dot00*dot12-dot01*dot02) * invDenom
	w = 1 - u - v
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
	maxX = max(max(max(minX, a.X), b.X), c.X)
	maxY = max(max(max(minY, a.Y), b.Y), c.Y)

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
