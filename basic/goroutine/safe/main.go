package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go func() {
		for {
			str := <-c
			_ = str
			time.Sleep(time.Second * 5)
		}
	}()

	for i := 0; i < 100; i++ {
		go func(i int) {
			c <- string(i)
			fmt.Printf("Writen %v Done.\n", i)
		}(i)
	}

	select {}
}
