package grid

import (
	"crypto/rand"
	"image"
	"image/color"
)

func CreateRandomImage(rect image.Rectangle) (grid [][]color.Color) {
	xlen, ylen := rect.Dx(), rect.Dy()
	for i := 0; i < xlen; i++ {
		var y []color.Color
		for j := 0; j < ylen; j++ {
			var rgba = make([]uint8, 4)
			rand.Read(rgba)
			var pix = color.NRGBA{rgba[0], rgba[1], rgba[2], rgba[3]}
			y = append(y, pix)
		}
		grid = append(grid, y)
	}
	return
}
