package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := "http://localhost:5000/ping"

	count := 0
	client := http.Client{
		// Transport: &http2.Transport{
		// 	AllowHTTP: true,
		// 	DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
		// 		return net.Dial(network, addr)
		// 	},
		// },
	}

	for i := 0; i < 1; i++ {
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
				fmt.Println(string(bd))
				resp.Body.Close()
			}
		}()
	}

	select {}
}
