package main

import (
 	"github.com/GLevis/gowebserver/pkg"
 )

func main() {
	server := server.InitServer(":8080")
	server.Run()
}
