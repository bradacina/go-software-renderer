package renderer

type Drawer interface {
	Draw(x, y int, color *ARGB)
}

func NewBuffer(width, height int) *Buffer {
	return &Buffer{Width: width, Height: height, Data: make([]byte, width*height*4)}
}

func (b *Buffer) Draw(x, y int, color *ARGB) {
	index := (x*b.Width + y) * 4

	b.Data[index] = byte(color.Blue)
	b.Data[index+1] = byte(color.Green)
	b.Data[index+2] = byte(color.Red)
	b.Data[index+3] = byte(color.Alpha)
}
