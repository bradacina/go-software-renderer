package renderer

// Buffer is a 2D buffer onto which rendering can be performed.
type Buffer struct {
	Width  int
	Height int

	// pixel information
	Data      []byte
	DebugData []byte

	// depth buf information
	DepthBuf []float64

	// helper to speed up transform of 3D space to 2D space
	halfWidth  float64
	halfHeight float64
}

// RGBA represents a color
type RGBA struct {
	Alpha byte
	Blue  byte
	Green byte
	Red   byte
}

// Point represents a 2D point
type Point struct {
	X int
	Y int
}

// Vertex represents a coordinate in 3D space
type Vertex struct {
	X float64
	Y float64
	Z float64
}

// Vector is just a vector in 3D space
type Vector struct {
	Vertex
}
