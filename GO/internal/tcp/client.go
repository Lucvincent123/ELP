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
	channel    chan struct{}
}

func NewClient(server string) *Client {
	return &Client{
		serAddress: server,
		channel:    make(chan struct{}),
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
	go c.readLoop()
	c.writeLoop()
}

func (c *Client) readLoop() {
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
				// var img [][]color.Color
				// img_buffer := buffer[4:n]
				// for i := 0; i < width; i++ {
				// 	var y []color.Color
				// 	for j := 0; j < height; j++ {
				// 		y = append(y, color.NRGBA{img_buffer[i*4+j*width*4], img_buffer[i*4+j*width*4+1], img_buffer[i*4+j*width*4+2], img_buffer[i*4+j*width*4+3]})
				// 	}
				// 	img = append(img, y)
				// }
				// ioFile.Save("filtered.png", img)
			}
			continue
		}
		msg := string(buffer[:n])

		// fmt.Println("Message:", msg)
		if msg == "close" {
			c.channel <- struct{}{}
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
		case <-c.channel:
			fmt.Println("Server is closed")
			close(c.channel)
			return
		default:
			if scanner.Scan() { // When user enters a command
				message := scanner.Text() // Read command string
				if message == "exit" {
					go c.forceWrite("close")
					break
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
					fmt.Println("Invalid command")
				}
				_, err := c.conn.Write([]byte(message))
				if err != nil {
					fmt.Println("Cannot send message:", err)
				}
			}
		}
	}
}

func (c *Client) forceWrite(cmd string) {
	var err error
	for {
		_, err = c.conn.Write([]byte(cmd))
		if err == nil {
			break
		}
	}
}
