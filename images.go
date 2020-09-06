package main

import(
	"golang.org/x/tour/pic"
	"image"
	"image/color"
);

type MyImage struct{}

func (myImage MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (myImage MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (myImage MyImage) At(x, y int) color.Color {
	return color.RGBA{20, 20, 20, 20}
}

func testImage() {
	m := MyImage{}
	pic.ShowImage(m)
}
