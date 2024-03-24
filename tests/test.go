package main

import (
	"encoding/binary"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// Open the Minecraft map item .dat file
	file, err := os.Open("map.dat")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read and decode Colors NBT tag
	var colors []int32
	err = binary.Read(file, binary.BigEndian, &colors)
	if err != nil {
		panic(err)
	}

	// Determine map dimensions (e.g., 128x128 pixels)
	width := 128
	height := 128

	// Create a new RGBA image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Iterate over Colors array and set pixel colors
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Extract ARGB components from Colors array
			colorInt := colors[y*width+x]
			alpha := uint8(colorInt >> 24)
			red := uint8(colorInt >> 16 & 0xFF)
			green := uint8(colorInt >> 8 & 0xFF)
			blue := uint8(colorInt & 0xFF)

			// Convert ARGB to RGBA color model
			rgbaColor := color.RGBA{red, green, blue, alpha}

			// Set pixel color in image
			img.Set(x, y, rgbaColor)
		}
	}

	// Save the image as PNG
	outputFile, err := os.Create("map.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Encode image as PNG and write to file
	err = png.Encode(outputFile, img)
	if err != nil {
		panic(err)
	}
}
