package ioFile

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func Load(filePath string) (grid [][]color.Color) {
	imgFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Cannot open file:", err)
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println("Cannot decode file:", err)
	}

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

func Save(filePath string, grid [][]color.Color) {
	xlen, ylen := len(grid), len(grid[0])
	rect := image.Rect(0, 0, xlen, ylen)
	img := image.NewNRGBA(rect)
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			img.Set(x, y, grid[x][y])
		}
	}

	imgFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Cannot create file:", err)
	}
	defer imgFile.Close()

	png.Encode(imgFile, img.SubImage(img.Rect))
}
