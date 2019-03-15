package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/hkwi/h2c"
)

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v\n", r.URL)
	fmt.Println("Received!")
}
func main() {

	go func() {
		mux := http.NewServeMux()
		mux.Handle("/", handler{})

		srv := &h2c.Server{
			Handler: mux,
		}

		fmt.Println("Starting Server to receive notification.")
		if err := http.ListenAndServe(":"+strconv.Itoa(9101), srv); err != nil {
			fmt.Println("Starting Server Failed")
		}
	}()

	select {}
}
