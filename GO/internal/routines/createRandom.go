package routines

import (
	"GO/internal/grid"
	"GO/internal/ioFile"
	"image"
)

func CreateRandomRoutine(width, height int, filePath string) {
	rect := image.Rect(0, 0, width, height)    // Create a rectangle of that width and that height
	randomGrid := grid.CreateRandomImage(rect) // Create a grid of random pixel of that rectangle
	ioFile.Save(filePath, randomGrid)          // Save image from that grid
}
