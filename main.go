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
		0:  {0, 0, 0},
		1:  {127, 178, 56},
		2:  {247, 233, 163},
		3:  {199, 199, 199},
		4:  {255, 0, 0},
		5:  {160, 160, 255},
		6:  {167, 167, 167},
		7:  {0, 124, 0},
		8:  {255, 255, 255},
		9:  {164, 168, 184},
		10: {151, 109, 77},
		11: {112, 112, 112},
		12: {64, 64, 255},
		13: {143, 119, 72},
		14: {255, 252, 245},
		15: {216, 127, 51},
		16: {178, 76, 216},
		17: {102, 153, 216},
		18: {229, 229, 51},
		19: {127, 204, 25},
		20: {242, 127, 165},
		21: {76, 76, 76},
		22: {153, 153, 153},
		23: {76, 127, 153},
		24: {127, 63, 178},
		25: {51, 76, 178},
		26: {102, 76, 51},
		27: {102, 127, 51},
		28: {153, 51, 51},
		29: {25, 25, 25},
		30: {250, 238, 77},
		31: {92, 219, 213},
		32: {74, 128, 255},
		33: {0, 217, 58},
		34: {129, 86, 49},
		35: {112, 2, 0},
		36: {209, 177, 161},
		37: {159, 82, 36},
		38: {149, 87, 108},
		39: {112, 108, 138},
		40: {186, 133, 36},
		41: {103, 117, 53},
		42: {160, 77, 78},
		43: {57, 41, 35},
		44: {135, 107, 98},
		45: {87, 92, 92},
		46: {122, 73, 88},
		47: {76, 62, 92},
		48: {76, 50, 35},
		49: {76, 82, 42},
		50: {142, 60, 46},
		51: {37, 22, 16},
		52: {189, 48, 49},
		53: {148, 63, 97},
		54: {92, 25, 29},
		55: {22, 126, 134},
		56: {58, 142, 140},
		57: {86, 44, 62},
		58: {20, 180, 133},
		59: {100, 100, 100},
		60: {216, 175, 147},
		61: {127, 167, 150},
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
