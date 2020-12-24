package main

import (
 	"github.com/GLevis/gowebserver/pkg/server"
 )

func main() {
	server := server.InitServer(":8080")
	server.Run()
}
