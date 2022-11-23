package main

import (
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Handler struct{}

func (c *Handler) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	body := []byte("I'm h2c server, you are on " + req.Proto)
	rsp.Header().Add("Content-Length", strconv.Itoa(len(body)))
	rsp.Header().Add("Content-Type", "application/json")
	rsp.WriteHeader(http.StatusOK)
	n, err := rsp.Write([]byte(body))
	fmt.Println(n, err)
}

func main() {
	server := http.Server{
		Addr:    ":5999",
		Handler: h2c.NewHandler(&Handler{}, &http2.Server{}),
	}

	server.ListenAndServe()
}
