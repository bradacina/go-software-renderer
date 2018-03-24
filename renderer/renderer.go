package renderer

type Buffer struct {
	Width      int
	Height     int
	Data       []byte
	DepthBuf   []float64
	halfWidth  float64
	halfHeight float64
}

type ARGB struct {
	Alpha byte
	Red   byte
	Green byte
	Blue  byte
}

type Point struct {
	X int
	Y int
}

type Vertex struct {
	X float64
	Y float64
	Z float64
}

type Vector struct {
	Vertex
}
