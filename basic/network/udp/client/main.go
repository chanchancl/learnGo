package main

import (
	"bytes"
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8999")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#v\n", addr)

	conn, err := net.DialUDP("udp", nil, addr)

	n, err := conn.Write(bytes.Repeat([]byte("a"), 1600))
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Println(n)
}
