package renderer

import (
	"math"
)

func (b *Buffer) Triangle(vertices *[3]Point, color *ARGB) {
	topLeft, bottomRight := bBox(vertices)

	tempVertex := Point{}

	for i := topLeft.X; i <= bottomRight.X; i++ {
		for j := topLeft.Y; j <= bottomRight.Y; j++ {
			tempVertex.X = i
			tempVertex.Y = j
			if isInTriangle(&tempVertex, vertices) &&
				// clip the triangle if outside the image boundaries
				i >= 0 && i < b.Height &&
				j >= 0 && j < b.Width {
				b.Draw(i, j, color)
			}
		}
	}
}

func isInTriangle(p *Point, vertices *[3]Point) bool {
	// Compute vectors
	v0 := VectorFromPoints(&vertices[0], &vertices[2])
	v1 := VectorFromPoints(&vertices[0], &vertices[1])
	v2 := VectorFromPoints(&vertices[0], p)

	// Compute dot products
	dot00 := dot(v0, v0)
	dot01 := dot(v0, v1)
	dot02 := dot(v0, v2)
	dot11 := dot(v1, v1)
	dot12 := dot(v1, v2)

	// Compute barycentric coordinates
	invDenom := 1 / float64(dot00*dot11-dot01*dot01)
	u := float64(dot11*dot02-dot01*dot12) * invDenom
	v := float64(dot00*dot12-dot01*dot02) * invDenom

	if u < 0 || v < 0 || u+v > 1 {
		return false
	}

	return true
}

func dot(v1, v2 *Point) int {
	return v1.X*v2.X + v1.Y*v2.Y
}

func VectorFromPoints(tail, head *Point) *Point {
	return &Point{X: head.X - tail.X, Y: head.Y - tail.Y}
}

// Calculates a bounding box for a triangle
func bBox(vertices *[3]Point) (*Point, *Point) {
	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := math.MinInt32, math.MinInt32

	for _, v := range *vertices {
		minX = min(minX, v.X)
		minY = min(minY, v.Y)
		maxX = max(maxX, v.X)
		maxY = max(maxY, v.Y)
	}

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
