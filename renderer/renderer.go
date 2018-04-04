package renderer

// Buffer is a 2D buffer onto which rendering can be performed.
type Buffer struct {
	Width  int
	Height int

	// pixel information
	Data []byte

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

var (
	ColorWhite = RGBA{Alpha: 255, Red: 255, Green: 255, Blue: 255}
	ColorRed   = RGBA{Alpha: 255, Red: 255, Green: 0, Blue: 0}
	ColorBlue  = RGBA{Alpha: 255, Red: 0, Green: 0, Blue: 255}
	ColorGreen = RGBA{Alpha: 255, Red: 0, Green: 255, Blue: 0}

	Mat4x4Identity = Mat4x4{
		AA: 1, AB: 0, AC: 0, AD: 0,
		BA: 0, BB: 1, BC: 0, BD: 0,
		CA: 0, CB: 0, CC: 1, CD: 0,
		DA: 0, DB: 0, DC: 0, DD: 1}
)

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

type AfineVertex struct {
	X float64
	Y float64
	Z float64
	W float64
}

type Mat3x3 struct {
	AA, AB, AC,
	BA, BB, BC,
	CA, CB, CC float64
}

type Mat4x4 struct {
	AA, AB, AC, AD,
	BA, BB, BC, BD,
	CA, CB, CC, CD,
	DA, DB, DC, DD float64
}

type Camera struct {
	ModelView Mat4x4
	buffer    Buffer
}
