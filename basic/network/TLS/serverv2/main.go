package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
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

	certBytes, err := ioutil.ReadFile("../certs/caRootCert.crt")
	if err != nil {
		panic("Unable to read caRootCert.crt")
	}
	certBlock, _ := pem.Decode(certBytes)
	rootCert, _ := x509.ParseCertificate(certBlock.Bytes)

	rootCertPool := x509.NewCertPool()
	rootCertPool.AddCert(rootCert)

	config := tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    rootCertPool,
		KeyLogWriter: os.Stdout,
	}

	fmt.Println("Listening :8080")
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
	log.Printf("Connect from %s\n", conn.RemoteAddr())

	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("Meet an error %v", err)
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
