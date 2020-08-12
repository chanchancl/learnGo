package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	now := time.Now()
	if r != nil && r.URL != nil {
		if r.URL.Query().Get("readfile") != "" {
			buf, _ := ioutil.ReadFile("tmp.data")
			_ = buf
		}
	}
	p := time.Now().Sub(now)
	fmt.Fprintf(w, "%v\n", p.Seconds())
}

func main() {
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/", Handler)
	httpServer := http.Server{
		Addr:    ":8080",
		Handler: httpMux,
	}

	go func() {
		fmt.Println("Http server is running.")
		httpServer.ListenAndServe()
	}()

	select {}
}