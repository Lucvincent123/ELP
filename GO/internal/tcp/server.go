package tcp

import (
	"fmt"
	"net"
)

type Server struct {
	listenAddress string
	listener      net.Listener
	quitChannel   chan struct{}
}

func NewServer(listenAddress string) *Server {
	return &Server{
		listenAddress: listenAddress,
		quitChannel:   make(chan struct{}),
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.listenAddress)
	if err != nil {
		return err
	}
	defer listener.Close()
	s.listener = listener
	go s.acceptLoop()
	<-s.quitChannel
	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("Cannot accept connection:", err)
			continue
		}
		fmt.Println("new connection", conn.RemoteAddr())
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Cannot read:", err)
			continue
		}
		msg := buffer[:n]
		fmt.Println("Message:", string(msg))
	}
}
