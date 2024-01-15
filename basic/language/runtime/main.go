package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello")

	lc, _ := net.Listen("tcp", "127.0.0.1:9999")

	for {
		lc.Accept()
	}
}
