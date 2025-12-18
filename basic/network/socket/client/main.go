package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	i := 0
	conn, err := net.Dial("tcp", "127.0.0.1:9998")
	for {
		fmt.Printf("%v, %v, %v\n", conn, err, i)
		if err != nil {
			i = 0
			conn, err = net.Dial("tcp", "127.0.0.1:9998")
			fmt.Println("Failed to connect remote")
			time.Sleep(time.Second)
			continue
		}

		data := strconv.Itoa(i)

		// if the connection drop, the err is not nil
		_, err = conn.Write([]byte(data))
		fmt.Printf("Success to send %v to %v\n", data, conn.RemoteAddr().String())
		i++
		// fmt.Printf("%+v %s\n", conn, conn.LocalAddr().String())
		time.Sleep(time.Second)
	}
}
