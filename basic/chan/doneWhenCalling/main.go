package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Printf("Go end %v\n", time.Now())
		c <- 5
	}()

	for {
		select {
		case <-c:
			fmt.Printf("recieved c %v\n", time.Now())
			return
		default:
			fmt.Printf("In default %v\n", time.Now())
			time.Sleep(5 * time.Second)
			fmt.Printf("Out default %v\n", time.Now())
		}
	}

	fmt.Printf("Main end %v\n", time.Now())
}
