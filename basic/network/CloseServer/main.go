package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	s := http.Server{Addr: ":9911"}

	go func() {
		for {
			time.Sleep(5 * time.Second)
			s.Close()

		}
	}()

	for {
		err := s.ListenAndServe()
		fmt.Println(err.Error())
	}
}
