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

func AverageGo(before [][]color.Color, level int) (after [][]color.Color) {
	before = reflectPadding(before, level)
	var filter [][]float64 = createBlurKernel(level)
	after = convolutionGo(before, filter)
	return
}

func Sharpener(before [][]color.Color) (after [][]color.Color) {
	before = reflectPadding(before, 1)
	var filter [][]float64 = createSharpenKernel()
	after = convolution(before, filter)
	return
}

func convolution(before [][]color.Color, filter [][]float64) (after [][]color.Color) {
	var widthFilter int = len(filter)
	var heigthFilter int = len(filter[0])
	var widthBefore int = len(before)
	var heigthBefore int = len(before[0])
	var widthAfter int = widthBefore - widthFilter + 1
	var heigthAfter int = heigthBefore - heigthFilter + 1
	for i := 0; i < widthAfter; i++ {
		var y []color.Color
		for j := 0; j < heigthAfter; j++ {
			var r float64 = 0
			var g float64 = 0
			var b float64 = 0
			var a float64 = 0
			for k := 0; k < widthFilter; k++ {
				for l := 0; l < heigthFilter; l++ {
					var r1, g1, b1, a1 uint32 = before[i+k][j+l].RGBA()
					var rf float64 = float64(r1)
					var gf float64 = float64(g1)
					var bf float64 = float64(b1)
					var af float64 = float64(a1)
					coefficent := filter[k][l]
					r = r + rf*(coefficent)/257
					g = g + gf*(coefficent)/257
					b = b + bf*(coefficent)/257
					a = a + af*(coefficent)/257
				}
			}
			var ri uint8 = uint8(r)
			var gi uint8 = uint8(g)
			var bi uint8 = uint8(b)
			var ai uint8 = uint8(a)
			y = append(y, color.NRGBA{ri, gi, bi, ai})
		}
		after = append(after, y)
	}
	return
}

func convolutionGo(before [][]color.Color, filter [][]float64) (after [][]color.Color) {
	var widthFilter int = len(filter)
	var heigthFilter int = len(filter[0])
	var widthBefore int = len(before)
	var heigthBefore int = len(before[0])
	var widthAfter int = widthBefore - widthFilter + 1
	var heigthAfter int = heigthBefore - heigthFilter + 1
	after = make([][]color.Color, widthAfter)
	for i := 0; i < widthAfter; i++ {
		after[i] = make([]color.Color, heigthAfter)
	}

	for i := 0; i < widthAfter; i++ {
		for j := 0; j < heigthAfter; j++ {
			go func() {
				var r float64 = 0
				var g float64 = 0
				var b float64 = 0
				var a float64 = 0
				for k := 0; k < widthFilter; k++ {
					for l := 0; l < heigthFilter; l++ {
						var r1, g1, b1, a1 uint32 = before[i+k][j+l].RGBA()
						var rf float64 = float64(r1)
						var gf float64 = float64(g1)
						var bf float64 = float64(b1)
						var af float64 = float64(a1)
						coefficent := filter[k][l]
						r = r + rf*(coefficent)
						g = g + gf*(coefficent)
						b = b + bf*(coefficent)
						a = a + af*(coefficent)
					}
				}
				var ri uint8 = uint8(r / 257)
				var gi uint8 = uint8(g / 257)
				var bi uint8 = uint8(b / 257)
				var ai uint8 = uint8(a / 257)
				after[i][j] = color.NRGBA{ri, gi, bi, ai}
			}()
		}
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
		{-1, 5, -1},
		{0, -1, 0},
	}
}
