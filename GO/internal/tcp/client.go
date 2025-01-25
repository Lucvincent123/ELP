package tcp

import (
	"GO/internal/ioFile"
	"bufio"
	"fmt"
	"image"
	"net"
	"os"
	"strings"
)

type Client struct {
	serAddress string
	conn       *net.TCPConn
	quit       chan struct{}
}

func NewClient(server string) *Client {
	return &Client{
		serAddress: server,
		quit:       make(chan struct{}),
	}
}

func (c *Client) Connect() {
	server, err := net.ResolveTCPAddr("tcp", c.serAddress)
	if err != nil {
		fmt.Println("Cannot resolve address:", err)
		return
	}

	conn, err := net.DialTCP("tcp", nil, server)
	if err != nil {
		fmt.Println("Cannot connect:", err)
		return
	}
	// defer c.conn.Close()
	c.conn = conn
	wg.Add(1)
	go c.readLoop()
	c.writeLoop()
	wg.Wait()
}

func (c *Client) readLoop() {
	defer wg.Done()
	defer c.conn.Close()
	buffer := make([]byte, 10000000000)
	for {
		n, err := c.conn.Read(buffer)
		if err != nil {
			fmt.Println("Cannot read:", err)
			continue
		}
		if buffer[0] == byte(1) {
			if buffer[1] == byte(0) {
				var width int = int(buffer[2])*255 + int(buffer[3])
				var height int = (n - 4) / 4 / width
				rect := image.Rect(0, 0, width, height)
				img := &image.NRGBA{
					Pix:    buffer[4:n],
					Stride: width * 4,
					Rect:   rect,
				}
				ioFile.SaveByte("filtered.png", img)
			}
			continue
		}
		msg := string(buffer[:n])

		// fmt.Println("Message:", msg)
		if msg == "close" {
			close(c.quit)
			break
		}
	}
}

func (c *Client) writeLoop() {
	// defer c.conn.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("(client->%v)->", c.serAddress)
		select {
		case <-c.quit:
			return
		default:
			if scanner.Scan() { // When user enters a command
				message := scanner.Text() // Read command string
				if message == "exit" {
					c.conn.Write([]byte("close"))
					return
				}
				if message == "shutdown" {
					c.conn.Write([]byte("shutdown"))
					return
				}
				message_parts := strings.Fields(message)
				if len(message_parts) == 0 {
					continue
				}
				switch message_parts[0] {
				case "filter":
					width, pix := ioFile.LoadByte(message_parts[1])
					var data []byte
					data = append(data, byte(1))
					data = append(data, byte(0))
					data = append(data, byte(width/255))
					data = append(data, byte(width%255))
					data = append(data, pix...)
					c.conn.Write(data)
				default:
					_, err := c.conn.Write([]byte(message))
					if err != nil {
						fmt.Println("Cannot send message:", err)
					}
				}

			}
		}
	}
}
