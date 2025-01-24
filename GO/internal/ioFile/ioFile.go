package ioFile

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func open(filePath string) *image.NRGBA {
	// Open file
	imgFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Cannot open file:", err)
	}
	defer imgFile.Close()
	// Decode file to *image.NRGBA
	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println("Cannot decode file:", err)
	}
	return img.(*image.NRGBA)
}

func Load(filePath string) (grid [][]color.Color) {
	// Open file to get pointer *image.NRGBA
	img := open(filePath)
	if img == nil {
		return nil
	}
	// Change to [][]color.Color
	grid = ImgToGrid(img)
	return
}

func LoadByte(filePath string) (width int, pix []byte) {
	// Open file to get pointer *image.NRGBA
	img := open(filePath)
	if img == nil {
		return
	}
	// Get width and pix
	width = img.Stride / 4
	pix = img.Pix
	return
}

func Save(filePath string, grid [][]color.Color) {
	// create file path
	imgFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Cannot create file:", err)
		return
	}
	//Change to *image.NRGBA
	img := GridToImg(grid)
	//Save image
	png.Encode(imgFile, img.SubImage(img.Rect))
	imgFile.Close()
	fmt.Printf("\nImage %v saved\n", imgFile.Name())
}

func SaveByte(filePath string, img *image.NRGBA) {
	imgFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Cannot create file:", err)
	}

	png.Encode(imgFile, img.SubImage(img.Rect))
	imgFile.Close()
	fmt.Printf("\nImage %v saved\n", imgFile.Name())
}
