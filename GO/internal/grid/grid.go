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

func Average(before [][]color.Color, level int) (after [][]color.Color) {
	before = reflectPadding(before, level)
	var filter [][]float64 = createBlurKernel(level)
	after = convolution(before, filter)
	return
}

func Sharpener(before [][]color.Color) (after [][]color.Color) {
	before = reflectPadding(before, 1)
	var filter [][]float64 = createSharpenKernel()
	after = convolution(before, filter)
	return
}

func convolution(before [][]color.Color, filter [][]float64) (after [][]color.Color) {
	var level int = len(filter)
	for i := (level - 1) / 2; i < len(before)-(level-1)/2; i++ {
		var y []color.Color
		for j := (level - 1) / 2; j < len(before[0])-(level-1)/2; j++ {
			var r float64 = 0
			var g float64 = 0
			var b float64 = 0
			var a float64 = 0
			for k := -(level - 1) / 2; k < level/2; k++ {
				for f := -(len(filter[0]) - 1) / 2; f < len(filter[0])/2; f++ {
					var r1, g1, b1, a1 uint32 = before[i+k][j+f].RGBA()
					var rf float64 = float64(r1)
					var gf float64 = float64(g1)
					var bf float64 = float64(b1)
					var af float64 = float64(a1)
					coefficent := filter[k+(len(filter)-1)/2][f+(len(filter[0])-1)/2]
					r = r + rf*(coefficent)
					g = g + gf*(coefficent)
					b = b + bf*(coefficent)
					a = a + af*(coefficent)
				}
			}
			var r_moyen uint8 = uint8(r / 257)
			var g_moyen uint8 = uint8(g / 257)
			var b_moyen uint8 = uint8(b / 257)
			var a_moyen uint8 = uint8(a / 257)
			pix := color.NRGBA{r_moyen, g_moyen, b_moyen, a_moyen}
			y = append(y, pix)
		}
		after = append(after, y)
	}
	return
}

func reflectPadding(before [][]color.Color, padding int) [][]color.Color {
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

func createBlurKernel(size int) [][]float64 {
	kernel := make([][]float64, size)
	for i := range kernel {
		kernel[i] = make([]float64, size)
		for j := range kernel[i] {
			kernel[i][j] = 1.0 / float64(size*size)
		}
	}
	return kernel
}

func createSharpenKernel() [][]float64 {
	return [][]float64{
		{0, -1, 0},
		{-1, 3, -1},
		{0, -1, 0},
	}
}
