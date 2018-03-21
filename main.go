package main

import (
	"./tga"
)

func main() {
	image := tga.NewTgaImage(1000, 1000)

	drawLine(0, 0, 200, 999, &tga.ARGB{255, 255, 255, 255}, image)
	drawLine(0, 0, 999, 200, &tga.ARGB{255, 255, 255, 255}, image)
	drawLine(0, 0, 999, 999, &tga.ARGB{255, 255, 255, 255}, image)
	image.Save("test.tga")
}
func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func drawLine(x0, y0, x1, y1 int, color *tga.ARGB, img tga.TgaImage) {
	steep := false
	if abs(x0-x1) < abs(y0-y1) {
		steep = true
		x0, y0 = y0, x0
		x1, y1 = y1, x1
	}

	if x0 > x1 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}

	dx := x1 - x0
	dy := y1 - y0
	derror := abs(dy) * 2
	err := 0
	y := y0
	for x := x0; x <= x1; x++ {
		if steep {
			img.Draw(y, x, color)
		} else {
			img.Draw(x, y, color)
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
