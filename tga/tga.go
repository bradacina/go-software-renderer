package tga

import (
	"encoding/binary"
	"log"
	"os"
)

func Save(width, height int, data []byte, filename string) {

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	byteOrder := binary.LittleEndian

	binary.Write(f, byteOrder, byte(0))         // idLength
	binary.Write(f, byteOrder, byte(0))         // colorMapType
	binary.Write(f, byteOrder, byte(2))         // imageType
	binary.Write(f, byteOrder, [5]byte{})       // colorMapSpec
	binary.Write(f, byteOrder, uint16(0))       // xOrigin
	binary.Write(f, byteOrder, uint16(0))       // yOrigin
	binary.Write(f, byteOrder, uint16(width))   // width
	binary.Write(f, byteOrder, uint16(height))  // height
	binary.Write(f, byteOrder, byte(32))        // pixelDepth
	binary.Write(f, byteOrder, byte(8<<4|3<<2)) // imageDesc
	binary.Write(f, byteOrder, data)
}
