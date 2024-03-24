package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"os"

	"github.com/Tnze/go-mc/nbt"
)

// https://github.com/mircokroon/minecraft-maps-to-images/blob/master/src/main/java/BasicColor.java
// https://minecraft.wiki/w/Map_item_format

func main() {
	// Step 1: Read the file
	data, err := os.ReadFile("map_0.dat")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Step 2: Decompress the gzip-compressed data
	gr, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		log.Fatalf("Failed to create gzip reader: %v", err)
	}
	defer gr.Close()

	decompressedData, err := io.ReadAll(gr)
	if err != nil {
		log.Fatalf("Failed to decompress data: %v", err)
	}

	// Step 3: Decode the NBT data
	var mapData MapData
	if err := nbt.Unmarshal(decompressedData, &mapData); err != nil {
		log.Fatalf("Failed to decode NBT data: %v", err)
	}

	fmt.Println(mapData)

	colorMap := map[int][]int{
		0:  {0, 0, 0, 0},
		1:  {0, 0, 0, 0},
		2:  {0, 0, 0, 0},
		3:  {0, 0, 0, 0},
		4:  {89, 125, 39, 255},
		5:  {109, 153, 48, 255},
		6:  {127, 178, 56, 255},
		7:  {67, 94, 29, 255},
		8:  {174, 164, 115, 255},
		9:  {213, 201, 140, 255},
		10: {247, 233, 163, 255},
		11: {130, 123, 86, 255},
		12: {140, 140, 140, 255},
		13: {171, 171, 171, 255},
		14: {199, 199, 199, 255},
		15: {105, 105, 105, 255},
		16: {180, 0, 0, 255},
		17: {220, 0, 0, 255},
		18: {255, 0, 0, 255},
		19: {135, 0, 0, 255},
		20: {112, 112, 180, 255},
	}

	// Define the dimensions of the image based on the size of the Colors array
	// For simplicity, let's assume the image is square and its size is the square root of the length of the Colors array
	imageSize := int(math.Sqrt(float64(len(mapData.Data.Colors))))
	img := image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))

	width := int(math.Sqrt(float64(len(mapData.Data.Colors)))) // Assuming the map is square
	height := width                                            // Assuming the map is square

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Calculate the index in the Colors array
			index := x + y*width
			colorIndex := mapData.Data.Colors[index]

			// Look up the RGB values in the colorMap
			rgb, ok := colorMap[int(colorIndex)]
			if !ok {
				log.Printf("Color index %d not found in colorMap", colorIndex)
				continue
			}

			// Set the pixel color in the image
			img.Set(x, y, color.RGBA{R: uint8(rgb[0]), G: uint8(rgb[1]), B: uint8(rgb[2]), A: 255})
		}
	}

	// Save the image to a file
	file, err := os.Create("output.png")
	if err != nil {
		log.Fatalf("Failed to create image file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatalf("Failed to encode image: %v", err)
	}

	fmt.Println("Image created successfully.")
}
