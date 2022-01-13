package main

import (
	"bufio"
	"crypto/tls"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	cert, err := tls.LoadX509KeyPair("../certs/server.crt", "../certs/serverPrivKey.pem")
	if err != nil {
		log.Println(err)
		return
	}

	config := tls.Config{Certificates: []tls.Certificate{cert}, KeyLogWriter: os.Stdout}
	ln, err := tls.Listen("tcp", ":8080", &config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	log.Printf("Connect to %s\n", conn.RemoteAddr())

	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Received : %s", msg)
		_, err = conn.Write([]byte("Hello\n"))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
