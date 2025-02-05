package tcp

import (
	"GO/internal/grid"
	"GO/internal/ioFile"
	"fmt"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Server struct {
	listenAddress string
	listener      net.Listener
	quit          chan struct{}
	conns         []*net.Conn
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
	defer s.listener.Close()
	fmt.Println("Listening at", s.listener.Addr().String())
	go s.acceptLoop()
	wg.Wait()
	<-s.quit
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			select {
			case <-s.quit:
				return
			default:
				fmt.Println("Cannot connect", err)
				continue
			}
		}
		s.conns = append(s.conns, &conn)
		fmt.Println("\nnew connection", conn.RemoteAddr())
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer fmt.Println(conn.RemoteAddr(), "is closed")
	defer conn.Close()
	buffer := make([]byte, 100000000)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			select {
			case <-s.quit:
				return
			default:
				fmt.Println("Cannot read", err)
				continue
			}
		}
		if buffer[0] == byte(1) {
			if buffer[1] == byte(0) {
				wg.Add(1)
				go func() {
					fmt.Println("recieved")
					defer wg.Done()
					var width int = int(buffer[2])*255 + int(buffer[3])
					imgGrid := ioFile.BytesToGrid(buffer[4:n], width)

					imgFiltered := grid.Average(imgGrid, 7)
					ioFile.Save("server.png", imgFiltered)
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
			if msg == "exit" {
				fmt.Println("Closing", conn.RemoteAddr())
				conn.Write([]byte("close"))
				// close(s.quit)
				return
			}
			if msg == "shutdown" {
				for _, conn := range s.conns {
					fmt.Println("Closing", (*conn).RemoteAddr())
					(*conn).Write([]byte("close"))
				}
				close(s.quit)
				return
			}
			fmt.Printf("[%v]: %v\n", conn.RemoteAddr(), msg)
		}
	}
}
