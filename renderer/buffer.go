package renderer

import (
	"log"
	"math"
)

// Drawer is an interface that can draw a pixel
type Drawer interface {
	Draw(x, y int, color *RGBA)
}

// NewBuffer creates a new render Buffer
func NewBuffer(width, height int, depthBuf bool) *Buffer {
	result := &Buffer{
		Width:      width,
		Height:     height,
		halfWidth:  float64(width-1) / 2,
		halfHeight: float64(height-1) / 2,
		Data:       make([]byte, width*height*4)}

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

func (b *Buffer) Read(col, row int) *RGBA {
	if col < 0 || col >= b.Width || row < 0 || row >= b.Height {
		log.Println("Trying to read from outide the buffer", col, row)
		return &RGBA{}
	}

	index := (row*b.Width + col) * 4

	return &RGBA{
		Blue:  b.Data[index],
		Green: b.Data[index+1],
		Red:   b.Data[index+2],
		Alpha: b.Data[index+3]}
}
