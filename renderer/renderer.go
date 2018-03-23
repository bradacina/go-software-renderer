package renderer

type Buffer struct {
	Width  int
	Height int
	Data   []byte
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
