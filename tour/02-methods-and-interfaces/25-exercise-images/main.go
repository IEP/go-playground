package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image structure
type Image struct {
	w, h int
}

// Bounds return image bounds
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

// ColorModel return image color model
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// At return pixel color at certain coordinate
func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x ^ y), uint8(x ^ y), 255, 255}
}

func main() {
	m := Image{500, 500}
	pic.ShowImage(m)
}
