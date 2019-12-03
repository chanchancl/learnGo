package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	go func() {
		for {

			tr := &http2.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				AllowHTTP:       true,
				DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
					return net.Dial(network, addr)
				},
			}
			c := http.Client{Transport: tr}

			req, err := http.NewRequest(http.MethodPost, "http://localhost:9101", bytes.NewBuffer([]byte{}))

			if err != nil {
				fmt.Println(err)
			}

			if rsp, err := c.Do(req); err != nil {
				fmt.Println(err)
			} else {
				if rsp.StatusCode != http.StatusOK {
					fmt.Println("status is " + rsp.Status)
				} else {
					fmt.Println("Send message Successful.")
				}
			}
			time.Sleep(time.Second)
		}
	}()

	select {}
}
