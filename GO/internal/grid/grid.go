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
					r = r + r1/257
					g = g + g1/257
					b = b + b1/257
					a = a + a1/257

				}
			}
			var r_moyen uint8 = uint8(r / 225)
			var g_moyen uint8 = uint8(g / 225)
			var b_moyen uint8 = uint8(b / 225)
			var a_moyen uint8 = uint8(a / 225)
			pix := color.NRGBA{r_moyen, g_moyen, b_moyen, a_moyen}
			y = append(y, pix)
		}
		after = append(after, y)
	}
	return
}

func Convolution(before [][]color.Color, filter [][]float64) (after [][]color.Color) {
	var niveau int = len(filter)
	for i := (niveau - 1) / 2; i < len(before)-(niveau-1)/2; i++ {
		var y []color.Color
		for j := (niveau - 1) / 2; j < len(before[0])-(niveau-1)/2; j++ {
			var r float64 = 0
			var g float64 = 0
			var b float64 = 0
			var a float64 = 0
			for k := -(niveau - 1) / 2; k < niveau/2; k++ {
				for f := -(len(filter[0]) - 1) / 2; f < len(filter[0])/2; f++ {
					var r1, g1, b1, a1 uint32 = before[i+k][j+f].RGBA()
					var rf float64 = float64(r1)
					var gf float64 = float64(g1)
					var bf float64 = float64(b1)
					var af float64 = float64(a1)
					coefficent := filter[k+(len(filter)-1)/2][f+(len(filter[0])-1)/2] //Normalement len(filter) = len(filter[0]) = niveau
					r = r + (rf/257)*(coefficent)
					g = g + (gf/257)*(coefficent)
					b = b + (bf/257)*(coefficent)
					a = a + (af/257)*(coefficent)
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

func Average(before [][]color.Color, niveau int) (after [][]color.Color) {
	before = ReflectPadding(before, niveau)
	var filter [][]float64
	filter = createBlurKernel(niveau)
	after = Convolution(before, filter)
	return
}

func Sharpener(before [][]color.Color) (after [][]color.Color) {
	before = ReflectPadding(before, 1)
	var filter [][]float64
	filter = createSharpenKernel()
	after = Convolution(before, filter)
	return
}

func ReflectPadding(before [][]color.Color, padding int) [][]color.Color { //Creer une nouvelle image avec Reflect Padding pour garder les pixels au bord
	rows := len(before)
	cols := len(before[0])

	newRows := rows + 2*padding
	newCols := cols + 2*padding

	paddedbefore := make([][]color.Color, newRows)
	for i := range paddedbefore {
		paddedbefore[i] = make([]color.Color, newCols)
	}

	for i := 0; i < newRows; i++ {
		for j := 0; j < newCols; j++ {

			reflectI := reflectIndex(i-padding, 0, rows-1)
			reflectJ := reflectIndex(j-padding, 0, cols-1)

			paddedbefore[i][j] = before[reflectI][reflectJ]
		}
	}

	return paddedbefore
}

func reflectIndex(index, min, max int) int {
	if index < min {
		return min + (min - index)
	}
	if index > max {
		return max - (index - max)
	}
	return index
}

func createBlurKernel(size int) [][]float64 { // Creer un Filtre pour moyenner l'image, il faut size impaire
	if size%2 == 0 {
		panic("Kernel size must be odd.")
	}
	kernel := make([][]float64, size)
	for i := range kernel {
		kernel[i] = make([]float64, size)
		for j := range kernel[i] {
			kernel[i][j] = 1.0 / float64(size*size)
		}
	}
	return kernel
}

func createSharpenKernel() [][]float64 { // Creer un filtre pour faire augmenter le contraste d'une image
	return [][]float64{
		{0, -1, 0},
		{-1, 5, -1},
		{0, -1, 0},
	}
}
