package main

import (
	"fmt"
	"net"
	"time"
)

func server() {
	ls, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		conn, _ := ls.Accept()
		fmt.Println(conn.LocalAddr().String())
		fmt.Println(conn.RemoteAddr().String())
		// Run in a goroutine, can accept multi connections
		{
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				continue
			}
			buf = buf[:n]
			fmt.Println(string(buf))
		}
	}
}

func client() {
	time.Sleep(time.Second)
	for {
		conn, err := net.Dial("tcp", "localhost:9999")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		conn.Write([]byte("123"))
		time.Sleep(time.Second)
	}
}

func main() {
	go server()
	go client()

	select {}

}
