package renderer

import (
	"math"
)

type Drawer interface {
	Draw(x, y int, color *ARGB)
}

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

func (b *Buffer) Draw(x, y int, color *ARGB) {
	index := (x*b.Width + y) * 4

	b.Data[index] = byte(color.Blue)
	b.Data[index+1] = byte(color.Green)
	b.Data[index+2] = byte(color.Red)
	b.Data[index+3] = byte(color.Alpha)
}
