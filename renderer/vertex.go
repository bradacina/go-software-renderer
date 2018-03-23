package renderer

import (
	"math"
)

func VectorFromVertex(head, tail *Vertex) *Vector {
	return &Vector{Vertex{X: head.X - tail.X, Y: head.Y - tail.Y, Z: head.Z - tail.Z}}
}

func CrossProduct(v1, v2 *Vector) *Vector {
	return &Vector{Vertex{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X}}
}

func MagnitudeSquared(v *Vector) float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func Normalize(v *Vector) {
	mag := math.Sqrt(MagnitudeSquared(v))

	v.X /= mag
	v.Y /= mag
	v.Z /= mag
}
