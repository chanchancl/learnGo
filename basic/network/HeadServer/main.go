package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		time.Sleep(time.Second)
		diff := time.Since(start)
		fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
		// tagLatency(w.Header(), diff)
		w.Header().Set("x-test-header", diff.String())
		fmt.Println(diff)
		// for k, v := range r.Header {
		// 	fmt.Fprintf(w, "Header [%q] = %q\n", k, v)
		// }
		// fmt.Fprintf(w, "Host = %q\n", r.Host)
		// fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
		// if err := r.ParseForm(); err != nil {
		// 	log.Print(err)
		// }
		// for k, v := range r.Form {
		// 	fmt.Fprintf(w, "Form [%q] = %q\n", k, v)
		// }
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func tagLatency(header http.Header, diff time.Duration) {
	header.Set("Test", strconv.Itoa(int(diff.Milliseconds())))
}
