package server

import (
	"fmt"
	"log"
	"net"

	"github.com/GLevis/gowebserver/pkg/request"
)

type Server struct {
	port string
}

func InitServer(port string) Server {
	server := Server{
		port: port,
	}
	return server
}

func (s *Server) Run() error {
	fmt.Printf("Starting server on port %s\n", s.port)

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
			r, err := request.ParseRequest(c)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf(r.Method + "\n")
			fmt.Printf(r.Target + "\n")
			fmt.Printf(r.Version + "\n")
			// Shut down the connection.
			c.Close()
		}(conn)
	}

	return nil
}
