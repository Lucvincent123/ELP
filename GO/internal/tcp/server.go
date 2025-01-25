package tcp

import (
	"GO/internal/grid"
	"GO/internal/ioFile"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex

type Server struct {
	listenAddress string
	listener      net.Listener
	quit          chan struct{}
}

func NewServer(listenAddress string) *Server {
	return &Server{
		listenAddress: listenAddress,
		quit:          make(chan struct{}),
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.listenAddress)
	if err != nil {
		fmt.Println("Cannot open tcp:", err)
		return
	}
	s.listener = listener
	fmt.Println("Listening at", s.listener.Addr().String())
	wg.Add(1)
	go s.acceptLoop()
	wg.Wait()
}

func (s *Server) acceptLoop() {
	defer wg.Done()
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			select {
			case <-s.quit:
				return
			default:
				fmt.Println("Cannot accept connection:", err)
				continue
			}
		}
		fmt.Println("\nnew connection", conn.RemoteAddr())
		wg.Add(1)
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer wg.Done()
	defer fmt.Println(conn.RemoteAddr(), "is closed")
	defer conn.Close()
	buffer := make([]byte, 10000000000)
	for {
		n, err := conn.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Cannot read", err)
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
			continue
		}
		msg := string(buffer[:n])
		if msg == "close" {
			fmt.Println("Closing", conn.RemoteAddr())
			conn.Write([]byte("close"))
			return
		}
		if msg == "shutdown" {
			fmt.Println("Closing", conn.RemoteAddr())
			conn.Write([]byte("close"))
			s.listener.Close()
			close(s.quit)
			return
		}
		fmt.Printf("[%v]: %v\n", conn.RemoteAddr(), msg)
	}
}
