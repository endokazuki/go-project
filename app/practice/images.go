package practice

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	x0    int
	y0    int
	x1    int
	y1    int
	red   int
	green int
	blue  int
	alpha int
}

//インターフェイスを使用
func (i Image) Bounds() image.Rectangle {
	return image.Rect(i.x0, i.y0, i.x1, i.y1)
}

//インターフェイスを無視
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

//インターフェイスを併用
func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), uint8(i.green), uint8(i.alpha)}
}

func ImageAction() {
	rectangle := Image{0, 0, 100, 100, 10, 10, 255, 255}
	pic.ShowImage(rectangle)
}

// func main() {
// 	m := Image{}
// 	pic.ShowImage(m)
// }
