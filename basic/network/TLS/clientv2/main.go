package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"log"
)

type wWriter struct{}

func (w *wWriter) Write(b []byte) (int, error) {
	return 0, errors.New("233")
}

func main() {
	cert, err := tls.LoadX509KeyPair("../certs/client.pem", "../certs/client.key")
	if err != nil {
		log.Println(err)
		return
	}
	certBytes, err := ioutil.ReadFile("../certs/client.pem")
	if err != nil {
		panic("Unable to read cert.pem")
	}
	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		panic("Failed to parse root certificate")
	}
	conf := tls.Config{
		RootCAs:            clientCertPool,
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
		KeyLogWriter:       &wWriter{},
	}

	conn, err := tls.Dial("tcp", "localhost:8080", &conf)
	if err != nil {
		log.Printf("tls.Dial error %v", err.Error())
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
