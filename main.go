package main

import (
	"./tga"
)

func main() {
	image := tga.NewTgaImage(100, 100)

	for i := 0; i < 100; i++ {
		image.Draw(i, i, tga.ARGB{byte(i), 255, 0, 0})
	}

	image.Save("test.tga")
}
