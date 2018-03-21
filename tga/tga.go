package tga

import (
	"encoding/binary"
	"log"
	"os"
)

type tga struct {
	width  int
	height int
	data   []byte
}

type ARGB struct {
	Alpha byte
	Red   byte
	Green byte
	Blue  byte
}

type TgaImage interface {
	Save(fileName string)
	Draw(x, y int, color *ARGB)
}

func NewTgaImage(width, height int) TgaImage {
	return &tga{width: width, height: height, data: make([]byte, width*height*4)}
}

func (tga *tga) Save(filename string) {

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	byteOrder := binary.LittleEndian

	binary.Write(f, byteOrder, byte(0))            // idLength
	binary.Write(f, byteOrder, byte(0))            // colorMapType
	binary.Write(f, byteOrder, byte(2))            // imageType
	binary.Write(f, byteOrder, [5]byte{})          // colorMapSpec
	binary.Write(f, byteOrder, uint16(0))          // xOrigin
	binary.Write(f, byteOrder, uint16(0))          // yOrigin
	binary.Write(f, byteOrder, uint16(tga.width))  // width
	binary.Write(f, byteOrder, uint16(tga.height)) // height
	binary.Write(f, byteOrder, byte(32))           // pixelDepth
	binary.Write(f, byteOrder, byte(8<<4|3<<2))    // imageDesc
	binary.Write(f, byteOrder, tga.data)
}

func (tga *tga) Draw(x, y int, color *ARGB) {
	index := (x*tga.width + y) * 4

	tga.data[index] = byte(color.Blue)
	tga.data[index+1] = byte(color.Green)
	tga.data[index+2] = byte(color.Red)
	tga.data[index+3] = byte(color.Alpha)
}
