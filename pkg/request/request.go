package request

import (
	"bufio"
	"net"
	"strings"
	"fmt"
)
type Request struct {
	Method string
	Target string
	Version string
	headers map[string]string
}

type UnparsableRequestError struct {
	s string
}

func (err *UnparsableRequestError) Error() string {
	return err.s
}

func ParseRequest(c net.Conn) (*Request, error){
	var r Request
	r.headers = make(map[string]string)
	scanner := bufio.NewScanner(c)
	scanner.Scan()
	requestLine := strings.Split(scanner.Text(), ' ')
	fmt.Printf("%s\n", requestLine)
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
