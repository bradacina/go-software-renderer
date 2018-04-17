package renderer

import (
	"math"
)

// VectorFromVertex creates a Vector from a Vertex
func VectorFromVertex(head, tail *Vertex) *Vector {
	return &Vector{Vertex{X: head.X - tail.X, Y: head.Y - tail.Y, Z: head.Z - tail.Z}}
}

// CrossProduct performs cross product of 2 Vectors
func CrossProduct(v1, v2 *Vector) *Vector {
	return &Vector{Vertex{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X}}
}

// DotProduct performs dot product of 2 Vectors
func DotProduct(v1, v2 *Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// MagnitudeSquared calculates the squared magnitude of a Vector
func MagnitudeSquared(v *Vector) float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func MulScalar(v *Vector, scalar float64) *Vector {
	return &Vector{Vertex{
		X: v.X * scalar,
		Y: v.Y * scalar,
		Z: v.Z * scalar}}
}

func Minus(v1, v2 *Vector) *Vector {
	return &Vector{Vertex{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z}}
}

// Normalize will normalize the length of a Vector
func Normalize(v *Vector) {
	mag := math.Sqrt(MagnitudeSquared(v))

	v.X /= mag
	v.Y /= mag
	v.Z /= mag
}

func Flip(v *Vector) *Vector {
	return MulScalar(v, -1)
}
