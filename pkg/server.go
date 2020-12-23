package server

import (
	"fmt"
	"net"
)

type Server struct {
	port string
}

func InitServer(port) {
	server := Server{
		port: port
	}
}

func (s *Server) Run() error {
	fmt.Printf("Starting server on port :%s\n", s.port)

	l, err := net.Listen("tcp", s.port)
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Process request

			// Shut down the connection.
			c.Close()
		}(conn)
	}

	return nil
}
