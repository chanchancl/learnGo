package main

import (
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

	ln, err := net.ListenUDP("udp", addr)

	for {
		var buf [1024]byte
		n, err := ln.Read(buf[:])
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("Read from %v : %s\n", n, buf[:n])
	}
}
