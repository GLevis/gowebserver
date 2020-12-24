package request

import (
	"bufio"
	"net"
	"strings"
)
type Request struct {
	Method string
	Target string
	Version string
	headers map[string]string
}

func ParseRequest(c net.Conn) (*Request, error){
	var r Request
	r.headers = make(map[string]string)
	scanner := bufio.NewScanner(c)
	scanner.Scan()
	requestLine := strings.Split(scanner.Text(), " ")
	r.Method = requestLine[0]
	r.Target = requestLine[1]
	r.Version = requestLine[2]
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		if len(line) < 2 {
			return nil, &UnparsableRequestError{"Error Parsing the Headers"}
		}
		r.headers[line[0]] = line[1]
	}
	return &r, nil
}
