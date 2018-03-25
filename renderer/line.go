package renderer

// LineDrawer is a interface that can DrawLines in 2D
type LineDrawer interface {
	DrawLine(x0, y0, x1, y1 int, color *RGBA)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// DrawLine will draw lines onto a Buffer
func (b *Buffer) DrawLine(x0, y0, x1, y1 int, color *RGBA) {
	steep := false

	// if the line is steep, we swap the coords
	// and we'll swap them back again right before rendering
	if abs(x0-x1) < abs(y0-y1) {
		steep = true
		x0, y0 = y0, x0
		x1, y1 = y1, x1
	}

	// sort the points, left to right
	if x0 > x1 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}

	// perform one time computations
	dx := x1 - x0
	dy := y1 - y0
	derror := abs(dy) * 2
	err := 0
	y := y0

	// go through each pixel between the 2 points
	for x := x0; x <= x1; x++ {
		if steep {
			b.Draw(y, x, color)
		} else {
			b.Draw(x, y, color)
		}

		err += derror
		if err > dx {
			if y1 > y0 {
				y++
			} else {
				y--
			}
			err -= dx * 2
		}
	}
}
