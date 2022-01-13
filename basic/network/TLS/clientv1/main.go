package main

import (
	"crypto/tls"
	"log"
	"os"
)

func main() {
	conf := tls.Config{
		InsecureSkipVerify: true,
		KeyLogWriter:       os.Stdout,
	}

	conn, err := tls.Dial("tcp", "localhost:8080", &conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("Hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}
	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}
	log.Println(string(buf[:n]))
}
