package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/xtaci/smux"
)

var (
	group sync.WaitGroup
	hello = "Hello 2024!"
	n     = len(hello)
)

func openConn() (conn1, conn2 net.Conn) {
	addr := "localhost:10029"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err.Error())
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		conn1, err = net.Dial("tcp", addr)
		defer wg.Done()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	conn2, err = ln.Accept()
	if err != nil {
		fmt.Println(err.Error())
	}

	wg.Wait()
	return conn1, conn2
}

func server(conn net.Conn) {
	defer group.Done()

	svr, _ := smux.Server(conn, nil)

	var streams []*smux.Stream
	streams = make([]*smux.Stream, n)

	for i := 0; i < n; i++ {
		stream, _ := svr.AcceptStream()
		streams[i] = stream
	}

	for i := 0; i < n; i++ {
		var buf [512]byte
		nn, _ := streams[i].Read(buf[:])
		fmt.Printf("%v,   Recv on %d stream with %s\n", time.Now(), i, string(buf[:nn]))
	}
}

func client(conn net.Conn) {
	defer group.Done()

	cli, _ := smux.Client(conn, nil)

	var streams []*smux.Stream
	streams = make([]*smux.Stream, n)

	for i := 0; i < n; i++ {
		stream, _ := cli.OpenStream()
		streams[i] = stream
	}

	for i := 0; i < n; i++ {
		streams[i].Write([]byte{hello[i]})
		fmt.Printf("%v, Send on %d stream with : %s\n", time.Now(), i, string(hello[i]))
	}
}

func main() {
	conn1, conn2 := openConn()
	fmt.Println(conn1, conn2)

	group.Add(2)
	go server(conn1)
	go client(conn2)

	group.Wait()
	conn1.Close()
	conn2.Close()
}
