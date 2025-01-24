package tcp

import (
	"GO/internal/grid"
	"GO/internal/ioFile"
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"
)

type Server struct {
	listenAddress string
	listener      net.Listener
	quit          chan struct{}
	conns         []*net.Conn
	wg            sync.WaitGroup
	mutex         sync.Mutex
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
	// defer listener.Close()
	s.listener = listener
	fmt.Println("Listening at", s.listener.Addr().String())
	s.wg.Add(1)
	go s.acceptLoop()
	s.commandLoop()
	// close(s.quit)
	s.wg.Wait()
	listener.Close()
	fmt.Println("Closed")
}

func (s *Server) acceptLoop() {
	defer s.wg.Done()
	for {
		fmt.Println("acceptLoop")
		conn, err := s.listener.Accept()
		if err != nil {
			select {
			case <-s.quit:
				fmt.Println("quit")
				return
			default:
				fmt.Println("Cannot accept connection:", err)
			}

		} else {
			fmt.Println("\nnew connection", conn.RemoteAddr())
			fmt.Printf("(server%v)->", s.listenAddress)
			s.mutex.Lock()
			s.conns = append(s.conns, &conn)
			s.mutex.Unlock()
			go func() {
				s.readLoop(conn)
			}()
		}

	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer fmt.Println(conn.RemoteAddr(), "is closed")
	defer conn.Close()
	// buffer := make([]byte, 10000000000)
	fmt.Println("reading loop")
	reader := bufio.NewReader(conn)
	for {
		select {
		case <-s.quit:
			fmt.Println("quit")
			return
		default:
			buffer, err := reader.ReadBytes(byte(2))
			if err != nil && err != io.EOF {
				fmt.Println("Cannot read", err)
				return
			}
			if buffer[0] == byte(1) {
				if buffer[1] == byte(0) {
					s.wg.Add(1)
					go func() {
						defer s.wg.Done()
						var width int = int(buffer[2])*255 + int(buffer[3])
						imgGrid := ioFile.BytesToGrid(buffer[4:], width)

						imgFiltered := grid.Average(imgGrid, 5)
						s.mutex.Lock()
						ioFile.Save("server.png", imgFiltered)
						fmt.Printf("(server%v)->", s.listenAddress)
						s.mutex.Unlock()
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
			}
			msg := string(buffer)
			if msg == "close" {
				conn.Write([]byte("close"))
				// close(s.quit)
			}
		}
		// n, err := conn.Read(buffer)

	}
}

func (s *Server) commandLoop() {
	// defer s.wg.Done()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("(server%v)->", s.listenAddress)
		if scanner.Scan() { // When user enters a command
			cmd := scanner.Text() // Read command string
			if cmd == "exit" {
				fmt.Print("exit is press")
				// s.Stop()
				close(s.quit)
				fmt.Println("Closing server...")
				return
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Scanner error:", err)
		}
	}
}

func (s *Server) Stop() {
	s.mutex.Lock()
	for _, conn := range s.conns {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			fmt.Println((*conn).RemoteAddr())
			(*conn).Close()
		}()
	}
	s.mutex.Unlock()
	close(s.quit)
}
