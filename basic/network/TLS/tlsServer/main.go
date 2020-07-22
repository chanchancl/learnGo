package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r != nil && r.URL != nil {
		scheme := r.URL.Scheme
		if scheme == "" {
			scheme = "http"
		}
		fmt.Fprintf(w,
			"Hi, this is an example of %s service in golang\n", scheme)
	}
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

	httpsMux := http.NewServeMux()
	httpsMux.HandleFunc("/", Handler)
	httpsServer := http.Server{
		Addr:    "localhost:8088",
		Handler: httpsMux,
		TLSConfig: &tls.Config{
			NextProtos: []string{"http/1.1"}, // this server will use http/1.1, otherwise http/2.0
		},
	}

	go func() {
		fmt.Println("Https server is running.")
		httpsServer.ListenAndServeTLS("../certs/server.pem", "../certs/server.key")
	}()

	select {}
}
