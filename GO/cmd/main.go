package main

import (
	"GO/internal/constants"
	"GO/internal/grid"
	"GO/internal/ioFile"
	"GO/internal/routines"
	"GO/internal/tcp"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // Using for reading user command
	var wg sync.WaitGroup
	fmt.Println("You are using Filmage", constants.Version) // Log version
	for {
		fmt.Print("->")
		if scanner.Scan() { // When user enters a command
			cmd := scanner.Text() // Read command string
			if cmd == "exit" {
				break
			}
			cmd_parts := strings.Fields(cmd) // Divide command into parts
			switch {
			case cmd_parts[0] == "create":
				switch {
				case cmd_parts[1] == "random": // create random 400 100 random.png
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
						routines.CreateRandomRoutine(width, height, cmd_parts[4])
					}()
				}
			case cmd_parts[0] == "filter":
				niveau, err := strconv.Atoi(cmd_parts[1])
				if err != nil {
					fmt.Println("Error:", err)
				}
				image := ioFile.Load("anhanime.png")
				filter := grid.Average(image, niveau)
				ioFile.Save("anhloc1.png", filter)
			case cmd_parts[0] == "server":
				server := tcp.NewServer(cmd_parts[1])
				log.Fatal(server.Start())
			case cmd_parts[0] == "client":
				client := tcp.NewClient()
				log.Fatal(client.Connect(cmd_parts[1]))
			default:
				fmt.Println("Invalid command")
			}

		}
		err := scanner.Err()
		if err != nil {
			fmt.Println("Scanner error:", err)
		} // Check scanner error
	}
	wg.Wait()
	fmt.Println("See you again")

}
