package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "0")
		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
		}
		if r.Method == http.MethodPost {
			w.WriteHeader(http.StatusCreated)
		}
		if r.Method == http.MethodPut {
			time.Sleep(time.Duration(20) * time.Second)
		}

		w.WriteHeader(http.StatusNotImplemented)
	})
	h2s := &http2.Server{}
	s := &http.Server{
		Addr:    ":9001",
		Handler: h2c.NewHandler(handler, h2s)}

	go s.ListenAndServe()

	// tr := &http2.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// 	AllowHTTP:       true,
	// 	DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
	// 		return net.Dial(network, addr)
	// 	},
	// }
	client := &http.Client{
		//Transport: tr,
		Timeout: time.Duration(time.Second * 2),
	}

	// You can change port 9001 to 9002
	// You will get connection refused instead of timeout
	req, _ := http.NewRequest("PUT", "http://localhost:9001", nil)

	rsp, err := client.Do(req)

	fmt.Printf("%v\n", rsp)
	fmt.Printf("%v\n", err)

	if e, ok := err.(net.Error); ok {
		if e.Timeout() {
			fmt.Println("This is timeout !")
		} else {
			fmt.Println("This is no response !")
		}
	}
}
