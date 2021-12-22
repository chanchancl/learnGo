package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rsp http.ResponseWriter, req *http.Request) {
		rsp.WriteHeader(200)
		rsp.Write([]byte("Hello!\n"))
	})

	server := http.Server{}

	server.Addr = ":10010"
	server.Handler = mux

	server.ListenAndServe()
}
