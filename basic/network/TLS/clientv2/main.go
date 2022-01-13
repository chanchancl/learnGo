package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

type wWriter struct{}

func (w *wWriter) Write(b []byte) (int, error) {
	fmt.Print(string(b))
	return len(b), nil
}

func main() {
	cert, err := tls.LoadX509KeyPair("../certs/client.crt", "../certs/clientPrivKey.pem")
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

	conf := tls.Config{
		RootCAs:            rootCertPool,
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
		KeyLogWriter:       &wWriter{},
	}

	conn, err := tls.Dial("tcp", "localhost:8080", &conf)
	if err != nil {
		log.Printf("tls.Dial error %v", err)
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
