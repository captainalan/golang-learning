package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

// From https://golang.org/pkg/image/#Image

type Image struct{
}

func (img Image) ColorModel() color.Model {
	// Hard-coded value for now
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	// Hard-coded value for now
	return image.Rectangle{image.Point{0, 0}, image.Point{100, 100}}
}

func (img Image) At(x, y int) color.Color {
	var v uint8 = uint8((x*y) % 255) // Different formulas will give different images
	return color.RGBA{v, v, 255, 255}
}


func main() {
	m := Image{}
	pic.ShowImage(m)
}
