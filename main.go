package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
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

// https://minecraft.wiki/w/Map_item_format

func main() {

	var filename string
	fmt.Print("Enter the name of the map file: (without .dat) ")
	fmt.Scan(&filename)

	data, err := os.ReadFile(filename + ".dat")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	gr, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		log.Fatalf("Failed to create gzip reader, make sure it is a valid map file: %v", err)
	}
	defer gr.Close()

	decompressedData, err := io.ReadAll(gr)
	if err != nil {
		log.Fatalf("Failed to decompress data: %v", err)
	}

	var mapData MapData
	if err := nbt.Unmarshal(decompressedData, &mapData); err != nil {
		log.Fatalf("Failed to decode NBT data: %v", err)
	}

	colorJson, err := os.ReadFile("colormap.json")
	if err != nil {
		log.Fatalf("Failed to read color map file: %v", err)
	}

	var colorMap ColorMap
	if err := json.Unmarshal(colorJson, &colorMap); err != nil {
		log.Fatalf("Failed to unmarshal color map JSON: %v", err)
	}

	imageSize := int(math.Sqrt(float64(len(mapData.Data.Colors))))
	img := image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))

	width := int(math.Sqrt(float64(len(mapData.Data.Colors))))
	height := width

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			index := x + y*width
			colorIndex := mapData.Data.Colors[index]

			rgb, ok := colorMap.Colormap_1_12[int(colorIndex)]
			if !ok {
				log.Printf("Color index %d not found in colorMap", colorIndex)
				continue
			}
			img.Set(x, y, color.RGBA{R: uint8(rgb[0]), G: uint8(rgb[1]), B: uint8(rgb[2]), A: 255})
		}
	}

	file, err := os.Create(filename + ".png")
	if err != nil {
		log.Fatalf("Failed to create image file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatalf("Failed to encode image: %v", err)
	}

	fmt.Printf("Image created successfully.\nFilename: %v\nDimensions: %vx%v\n", filename, width, height)
}
