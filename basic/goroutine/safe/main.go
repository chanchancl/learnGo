package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	c := make(chan string)

	go func() {
		for {
			str := <-c
			_ = str
			time.Sleep(time.Millisecond)
		}
	}()

	for i := 0; i < 5; i++ {
		c <- strconv.Itoa(i)
		fmt.Printf("Writen %v Done.\n", i)
	}
}
