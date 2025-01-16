package grid

import (
	"crypto/rand"
	"image"
	"image/color"
	"time"
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
	time.Sleep(20 * time.Second)
	return
}

// func Average9(before [][]color.Color) (after [][]color.Color) {
// 	// create matrice 3x3
// 	// convolui
// }

// func Convolution(before, h [][]color.Color) (after [][]color.Color) {

// }

func Average9(before [][]color.Color) (after [][]color.Color) {

	for i := 7; i < len(before)-7; i++ {
		var y []color.Color
		for j := 7; j < len(before[0])-7; j++ { //parcourir tous les pixels de l'image
			var r uint32 = 0
			var g uint32 = 0
			var b uint32 = 0
			var a uint32 = 0
			for k := -7; k < 8; k++ { //Calculer la somme de r g b a de tous les 9 pixels voisins
				for f := -7; f < 8; f++ {
					r1, g1, b1, a1 := before[i+k][j+f].RGBA()
					r = r + r1/257/2
					g = g + g1/257/2
					b = b + b1/257*0
					a = a + a1/257

				}
			}
			var r_moyen uint8 = uint8(r)
			var g_moyen uint8 = uint8(g)
			var b_moyen uint8 = uint8(b)
			var a_moyen uint8 = uint8(a)
			pix := color.NRGBA{r_moyen, g_moyen, b_moyen, a_moyen}
			y = append(y, pix)
		}
		after = append(after, y)
	}
	return
}
