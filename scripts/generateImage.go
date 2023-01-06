package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// Create a new image with a white background
	img := image.NewRGBA(image.Rect(0, 0, 300, 300))
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			img.Set(x, y, color.RGBA{100, 100, 100, 255})
		}
	}

	// Save the image as a PNG file
	f, _ := os.Create("image.png")
	defer f.Close()
	png.Encode(f, img)
}
