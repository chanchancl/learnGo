package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Handler struct{}

func (c *Handler) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	body := []byte("I'm http 1.1 server")
	rsp.Header().Add("Content-Length", strconv.Itoa(len(body)+10))
	n, err := rsp.Write([]byte(body))
	rsp.WriteHeader(http.StatusOK)
	fmt.Println(n, err)
}

func main() {
	server := http.Server{
		Addr:    ":5999",
		Handler: &Handler{},
	}

	server.ListenAndServe()
}
