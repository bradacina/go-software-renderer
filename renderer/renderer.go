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
