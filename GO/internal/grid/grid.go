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
func Average9(before [][]color.Color) (after [][]color.Color) {

	for i := 1; i < len(before)-1; i++ {
		var y []color.Color
		for j := 1; j < len(before[0])-1; j++ { //parcourir tous les pixels de l'image
			var r uint32 = 0
			var g uint32 = 0
			var b uint32 = 0
			var a uint32 = 0
			for k := -1; k < 2; k++ { //Calculer la somme de r g b a de tous les 9 pixels voisins
				for f := -1; f < 2; f++ {
					r1, g1, b1, a1 := before[i+k][j+f].RGBA()
					r = r + r1
					g = g + g1
					b = b + b1
					a = a + a1

				}
			}
			var r_moyen uint8 = uint8(r / 9)
			var g_moyen uint8 = uint8(g / 9)
			var b_moyen uint8 = uint8(b / 9)
			var a_moyen uint8 = uint8(a / 9)
			pix := color.NRGBA{r_moyen, g_moyen, b_moyen, a_moyen}
			y = append(y, pix)
		}
		after = append(after, y)
	}
	return
}
