package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 1)
	a := make(chan int)
	b := make(chan int)
	exit := false
	go func() {
		for {
			if exit {
				b <- 5
				break
			}
			select {
			case <-ticker.C:
				ticker = time.NewTicker(time.Second * 1)
				fmt.Println("233")
				//break
			case <-a:
				fmt.Println("exit")
				exit = true
				break
			}
		}
	}()

	time.Sleep(time.Second*3 + 100000)
	a <- 5
	<-b
}
