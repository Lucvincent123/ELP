package main

import (
	"GO/internal/constants"
	"GO/internal/grid"
	"GO/internal/ioFile"
	"GO/internal/tcp"
	"bufio"
	"fmt"
	"image"
	random "math/rand/v2"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	scanner := bufio.NewScanner(os.Stdin)                   // Using for reading user command
	fmt.Println("You are using Filmage", constants.Version) // Log version
	for {
		fmt.Print("->")
		if scanner.Scan() { // When user enters a command
			cmd := scanner.Text() // Read command string
			if cmd == "exit" {    // Exit command
				break
			}
			cmd_parts := strings.Fields(cmd) // Divide command into parts
			if len(cmd_parts) == 0 {         // Skip empty command
				continue
			}
			switch cmd_parts[0] {
			case "create":
				switch cmd_parts[1] {
				case "random": // create random 400 100 random.png
					width, err := strconv.Atoi(cmd_parts[2]) // String to int
					if err != nil {
						fmt.Println("Invalid width")
					}
					height, err := strconv.Atoi(cmd_parts[3]) // String to int
					if err != nil {
						fmt.Println("Invalid height")
					}
					wg.Add(1)
					go func() {
						defer wg.Done()
						rect := image.Rect(0, 0, width, height)    // Create a rectangle of that width and that height
						randomGrid := grid.CreateRandomImage(rect) // Create a grid of random pixel of that rectangle
						time.Sleep(time.Duration(random.IntN(10)) * time.Second)
						mutex.Lock()
						ioFile.Save(cmd_parts[4], randomGrid) // Save image from that grid
						mutex.Unlock()
						fmt.Print("->")
					}()
				default:
					fmt.Println("Invalid create argument")
				}
			case "filter":
				switch cmd_parts[1] {
				case "average": // filter average 9 source.png destination.png
					level, err := strconv.Atoi(cmd_parts[2]) // Get user's level of average
					if err != nil {
						fmt.Println("Error:", err)
						continue
					}
					wg.Add(1)
					go func() {
						defer wg.Done()
						mutex.Lock()
						image := ioFile.Load(cmd_parts[3])
						mutex.Unlock()
						new_img := grid.Average(image, level*2+1)
						time.Sleep(time.Duration(random.IntN(10)) * time.Second)
						mutex.Lock()
						ioFile.Save(cmd_parts[4], new_img)
						mutex.Unlock()
						fmt.Print("->")
					}()
				case "sharpen": // filter sharpen source.png destination.png
					wg.Add(1)
					go func() {
						defer wg.Done()
						mutex.Lock()
						image := ioFile.Load(cmd_parts[2])
						mutex.Unlock()
						new_img := grid.Sharpener(image)
						time.Sleep(time.Duration(random.IntN(10)) * time.Second)
						mutex.Lock()
						ioFile.Save(cmd_parts[3], new_img)
						mutex.Unlock()
						fmt.Print("->")
					}()
				default:
					fmt.Println("Invalid filter command")
				}
			case "server":
				server := tcp.NewServer(cmd_parts[1])
				server.Start()
				fmt.Println("Back to main menu")
				fmt.Print("->")
			case "client":
				client := tcp.NewClient(cmd_parts[1])
				client.Connect()
			default:
				fmt.Println("Invalid command")
			}

		}
		// Check scanner error
		if err := scanner.Err(); err != nil {
			fmt.Println("Scanner error:", err)
		}
	}
	fmt.Print("Exiting...")
	wg.Wait()
	fmt.Println("\nSee you again")
}
