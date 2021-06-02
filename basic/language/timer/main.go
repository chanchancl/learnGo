package main

import "time"

func main() {
	c := make(chan int)
	timer := time.NewTimer(time.Second)
	go func() {
		select {
		case <-timer.C:
			timer.Stop()
			c <- 1
		}
	}()
	<-c
}
