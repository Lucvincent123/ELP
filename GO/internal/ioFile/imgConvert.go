package ioFile

import (
	"image"
	"image/color"
)

func ImgToGrid(img *image.NRGBA) (grid [][]color.Color) {
	size := img.Bounds().Size()
	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j))
		}
		grid = append(grid, y)
	}
	return
}

func GridToImg(grid [][]color.Color) (img *image.NRGBA) {
	xlen, ylen := len(grid), len(grid[0])
	rect := image.Rect(0, 0, xlen, ylen)
	img = image.NewNRGBA(rect)
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			img.Set(x, y, grid[x][y])
		}
	}
	return
}

func BytesToGrid(bytes []byte, imgWidth int) (grid [][]color.Color) {
	var imgHeight int = len(bytes) / 4 / imgWidth
	for i := 0; i < imgWidth; i++ {
		var y []color.Color
		for j := 0; j < imgHeight; j++ {
			y = append(y, color.NRGBA{bytes[i*4+j*imgWidth*4], bytes[i*4+j*imgWidth*4+1], bytes[i*4+j*imgWidth*4+2], bytes[i*4+j*imgWidth*4+3]})
		}
		grid = append(grid, y)
	}
	return
}

func GridToBytes(grid [][]color.Color) (bytes []byte, imgWidth int) {
	imgWidth = len(grid)
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < imgWidth; j++ {
			r, g, b, a := grid[j][i].RGBA()
			bytes = append(bytes, byte(r>>8))
			bytes = append(bytes, byte(g>>8))
			bytes = append(bytes, byte(b>>8))
			bytes = append(bytes, byte(a>>8))
		}
	}
	return
}
