package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/net/http2"
)

func main() {
	url := "http://localhost:9001"

	count := 0
	client := http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}

	for i := 0; i < 100; i++ {
		go func() {
			for {
				count++
				req, _ := http.NewRequest(http.MethodGet, url, nil)
				resp, err := client.Do(req)
				if err != nil {
					fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
					continue
				}
				bd, _ := ioutil.ReadAll(resp.Body)
				it, _ := strconv.Atoi(string(bd))
				fmt.Println(it)
				resp.Body.Close()
			}
		}()
	}

	select {}
}
