package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Client struct {
	conn *net.TCPConn
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Connect(serAddress string) error {
	server, err := net.ResolveTCPAddr("tcp", serAddress)
	if err != nil {
		fmt.Println("Cannot resolve address")
	}

	conn, err := net.DialTCP("tcp", nil, server)
	if err != nil {
		fmt.Println("Cannot conncect")
	}
	c.conn = conn
	go c.readLoop()
	c.writeLoop()
	return nil
}

func (c *Client) readLoop() {
	defer c.conn.CloseRead()
	buffer := make([]byte, 2048)
	for {
		n, err := c.conn.Read(buffer)
		if err != nil {
			fmt.Println("Cannot read:", err)
			continue
		}
		msg := buffer[:n]
		fmt.Println("Message:", string(msg))
	}
}

func (c *Client) writeLoop() {
	defer c.conn.CloseWrite()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("->")
		if scanner.Scan() { // When user enters a command
			message := scanner.Text() // Read command string
			if message == "exit" {
				break
			}
			_, err := c.conn.Write([]byte(message))
			if err != nil {
				fmt.Println("Cannot send message:", err)
			}
		}
	}
}
