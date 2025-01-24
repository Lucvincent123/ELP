package tcp

import (
	"GO/internal/grid"
	"GO/internal/ioFile"
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex

type Server struct {
	listenAddress string
	listener      net.Listener
	quit          chan bool
	conns         []*net.Conn
	quitConns     []chan bool
}

func NewServer(listenAddress string) *Server {
	return &Server{
		listenAddress: listenAddress,
		quit:          make(chan bool),
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.listenAddress)
	if err != nil {
		fmt.Println("Cannot open tcp:", err)
		return
	}
	defer listener.Close()
	s.listener = listener
	fmt.Println("Listening at", s.listener.Addr().String())
	go s.acceptLoop()
	s.commandLoop()
	fmt.Println("Closing server...")
	wg.Wait()
	close(s.quit)
	fmt.Println("Closed")
}

func (s *Server) acceptLoop() {
	for {
		select {
		case <-s.quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("acceptLoop")
			conn, err := s.listener.Accept()
			if err != nil {
				fmt.Println("Cannot accept connection:", err)
				return
			}
			quitConn := make(chan bool)
			mutex.Lock()
			s.conns = append(s.conns, &conn)
			s.quitConns = append(s.quitConns, quitConn)
			mutex.Unlock()
			fmt.Println("\nnew connection", conn.RemoteAddr())
			fmt.Printf("(server%v)->", s.listenAddress)
			fmt.Println(s.quitConns)
			fmt.Println(s.conns)
			wg.Add(1)
			go s.readLoop(conn, quitConn)
		}
	}
}

func (s *Server) readLoop(conn net.Conn, quitConn chan bool) {
	defer conn.Close()
	defer wg.Done()
	buffer := make([]byte, 10000000000)
	for {
		select {
		case <-quitConn:
			fmt.Println(conn.RemoteAddr(), "closed")
			return
		default:
			fmt.Println("reading loop")
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Cannot read:", err)
				return
			}
			if buffer[0] == byte(1) {
				if buffer[1] == byte(0) {
					wg.Add(1)
					go func() {
						defer wg.Done()
						var width int = int(buffer[2])*255 + int(buffer[3])
						imgGrid := ioFile.BytesToGrid(buffer[4:n], width)

						imgFiltered := grid.Average(imgGrid, 5)
						mutex.Lock()
						ioFile.Save("server.png", imgFiltered)
						fmt.Printf("(server%v)->", s.listenAddress)
						mutex.Unlock()
						pix, newWidth := ioFile.GridToBytes(imgFiltered)
						var data []byte
						data = append(data, byte(1))
						data = append(data, byte(0))
						data = append(data, byte(newWidth/255))
						data = append(data, byte(newWidth%255))
						data = append(data, pix...)
						time.Sleep(10 * time.Second)
						conn.Write(data)
					}()
				}
			} else {
				msg := string(buffer[:n])
				if msg == "close" {
					conn.Write([]byte("close"))
					close(quitConn)
				}
			}
		}
	}
}

func (s *Server) commandLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("(server%v)->", s.listenAddress)
		if scanner.Scan() { // When user enters a command
			cmd := scanner.Text() // Read command string
			if cmd == "exit" {
				fmt.Print("exit is press")
				s.stop()
				break
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Scanner error:", err)
		}
	}
}

func (s *Server) stop() {
	for i, conn := range s.conns {
		_, err := (*conn).Write([]byte("close"))
		if err != nil {
			fmt.Println(err)
			return
		}
		close(s.quitConns[i])
		// (*conn).Close()
		fmt.Println((*conn).RemoteAddr(), " is closed")
	}
}
