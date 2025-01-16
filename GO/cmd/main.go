package main

import (
	"GO/internal/constants"
	"GO/internal/grid"
	"GO/internal/ioFile"
	"bufio"
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("You are using Filmage", constants.Version)
	for {
		if scanner.Scan() {
			cmd := scanner.Text()
			if cmd == "exit" {
				break
			}
			cmd_parts := strings.Fields(cmd)
			switch {
			case cmd_parts[0] == "create":
				switch {
				case cmd_parts[1] == "random":
					width, err := strconv.Atoi(cmd_parts[2])
					if err != nil {
						fmt.Println("Invalid width")
					}
					height, err := strconv.Atoi(cmd_parts[3])
					if err != nil {
						fmt.Println("Invalid height")
					}
					rect := image.Rect(0, 0, width, height)
					randomImage := grid.CreateRandomImage(rect)
					ioFile.Save("random.png", randomImage)

				case cmd_parts[0] == "filter":
					image := ioFile.Load("anhnguoi.png")
					filter := grid.Average9(image)
					ioFile.Save("anhloc.png", filter)
				}
			}

		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Erreur :", err)
		}
	}
	fmt.Println("See you again")

}
