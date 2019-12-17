package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	close(intChan)
	syncChan := make(chan struct{}, 1)
	go func() {
	Loop:
		for {
			select {
			case e, ok := <-intChan:
				if !ok {
					fmt.Println("End.")
					break Loop
				}
				fmt.Printf("Received: %v\n", e)
			}
		}
		syncChan <- struct{}{}
	}()
	<-syncChan

	c := make(chan int)
	go func() {
		select {
		case <-c:
			fmt.Println("233")
		default:
			time.Sleep(time.Second)

		}
		fmt.Println("here?")
	}()

	close(c)
	time.Sleep(time.Second * 5)
}
