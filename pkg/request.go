package request

import (
	"bufio"
	"net"
)
type Request struct {
	method string
	target string
	version string
	headers map[string]string
}

func parseRequest(c net.Conn) (*Request, error){
	var r Request
	r.headers := make(map[string]string)
	scanner := bufio.NewScanner(c)
	scanner.Scan()
	requestLine := scanner.Text().Split(" ")
	r.method := requestLine[0]
	r.target := requestLine[1]
	r.version := requestLine[2]
	for scanner.Scan() {
		line := scanner.Text().Split(":")
		r.headers[line[0]] := line[1]
	}
	return &r, nil
}
