package tga

import (
	"encoding/binary"
	"log"
	"os"
)

/*
	https://en.wikipedia.org/wiki/Truevision_TGA
*/

// Load loads a TGA file (true color, 24 or 32 bits per color)
func Load(filename string) (int, int, int, *[]byte) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	byteOrder := binary.LittleEndian

	var idLength byte
	var colorMapType byte
	var imageType byte
	var colorMapSpec [5]byte
	var xOrigin uint16
	var yOrigin uint16
	var width int16
	var height int16
	var pixelDepth byte
	var imageDescriptor byte

	err = binary.Read(f, byteOrder, &idLength)
	if err != nil {
		log.Fatal("When reading idLength", err)
	}

	err = binary.Read(f, byteOrder, &colorMapType)
	if err != nil {
		log.Fatal("When reading colorMapType", err)
	}

	if colorMapType != 0 {
		log.Fatal("TGA color map type not supported", colorMapType)
	}

	err = binary.Read(f, byteOrder, &imageType)
	if err != nil {
		log.Fatal("When reading imageType", err)
	}

	if imageType != 2 {
		log.Fatal("TGA image type not supported", imageType)
	}

	err = binary.Read(f, byteOrder, &colorMapSpec)
	if err != nil {
		log.Fatal("When reading colorMapSpec", err)
	}

	err = binary.Read(f, byteOrder, &xOrigin)
	if err != nil {
		log.Fatal("When reading xOrigin", err)
	}

	err = binary.Read(f, byteOrder, &yOrigin)
	if err != nil {
		log.Fatal("When reading yOrigin", err)
	}

	err = binary.Read(f, byteOrder, &width)
	if err != nil {
		log.Fatal("When reading width", err)
	}

	if width < 0 {
		log.Fatal("TGA file width is less than 0:", width)
	}
	err = binary.Read(f, byteOrder, &height)
	if err != nil {
		log.Fatal("When reading height", err)
	}

	if height < 0 {
		log.Fatal("TGA file height is less than 0:", height)
	}

	err = binary.Read(f, byteOrder, &pixelDepth)
	if err != nil {
		log.Fatal("When reading pixelDepth", err)
	}

	if pixelDepth != 32 && pixelDepth != 24 {
		log.Fatal("TGA pixel depth is not supported", pixelDepth)
	}

	err = binary.Read(f, byteOrder, &imageDescriptor)
	if err != nil {
		log.Fatal("When reading imageDescriptor", err)
	}

	// todo: handle image descriptor to see how many bytes for alpha channel

	bytesPerPixel := int(pixelDepth) / 8

	size := int(width) * int(height) * bytesPerPixel

	// color info is stored in B-G-R-A order
	data := make([]byte, size)

	err = binary.Read(f, byteOrder, &data)
	if err != nil {
		log.Fatal("When reading pixel data", err)
	}

	// add in the alpha channel
	if bytesPerPixel == 3 {
		dataOld := data
		data = make([]byte, int(width)*int(height)*4)

		idx := 0
		for i, v := range dataOld {
			data[idx] = v
			idx++
			if (i+1)%3 == 0 {
				data[idx] = 255 // alpha channel
				idx++
			}
		}
	}

	return int(width), int(height), bytesPerPixel, &data
}

// Save saves pixel info to a .tga file
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
