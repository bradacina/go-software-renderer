package renderer

import (
	"log"
	"math"
)

// NewBuffer creates a new render Buffer
func NewBuffer(width, height int, depthBuf bool) *Buffer {
	result := &Buffer{
		Width:  width,
		Height: height,
		Data:   make([]byte, width*height*4)}

	if depthBuf {
		result.DepthBuf = make([]float64, width*height)
		for i := range result.DepthBuf {
			result.DepthBuf[i] = -math.MaxFloat64
		}
	}

	return result
}

// Draw draws a pixel onto the Buffer
func (b *Buffer) Draw(col, row int, color *RGBA) {
	if col < 0 || col >= b.Width || row < 0 || row >= b.Height {
		log.Println("Trying to draw pixel outside buffer bounds")
		return
	}

	index := (row*b.Width + col) * 4

	b.Data[index] = color.Blue
	b.Data[index+1] = color.Green
	b.Data[index+2] = color.Red
	b.Data[index+3] = color.Alpha
}

func (b *Buffer) Read(col, row int, color *RGBA) {
	if col < 0 || col >= b.Width || row < 0 || row >= b.Height {
		log.Println("Trying to read from outide the buffer", col, row)
		color.Alpha = 0
		color.Red = 0
		color.Blue = 0
		color.Green = 0
	}

	index := (row*b.Width + col) * 4

	color.Blue = b.Data[index]
	color.Green = b.Data[index+1]
	color.Red = b.Data[index+2]
	color.Alpha = b.Data[index+3]
}

func (buf *Buffer) FlipVertically() {
	pixel := make([]byte, buf.Width*4)
	bufWidth := buf.Width * 4
	for row := 0; row < buf.Height; row++ {
		indexFrom := row * bufWidth
		indexTo := (buf.Height - row - 1) * bufWidth

		copy(pixel, buf.Data[indexTo:indexTo+bufWidth])
		copy(buf.Data[indexTo:indexTo+bufWidth], buf.Data[indexFrom:indexFrom+bufWidth])
		copy(buf.Data[indexFrom:indexFrom+bufWidth], pixel)
	}
}
