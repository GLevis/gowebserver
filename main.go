package main

import (
 	"github.com/GLevis/gowebserver"
 )

func main() {
	server := server.InitServer("80")
	server.Run()
}
