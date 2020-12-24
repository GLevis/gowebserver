package request

import (
	"bufio"
	"net"
	"strings"
)
type Request struct {
	method string
	target string
	version string
	headers map[string]string
}

func parseRequest(c net.Conn) (*Request, error){
	var r Request
	r.headers = make(map[string]string)
	scanner := bufio.NewScanner(c)
	scanner.Scan()
	requestLine := strings.Split(scanner.Text(), " ")
	r.method = requestLine[0]
	r.target = requestLine[1]
	r.version = requestLine[2]
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		r.headers[line[0]] = line[1]
	}
	return &r, nil
}
