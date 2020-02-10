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
		case a, err := <-c:
			fmt.Printf("233 %v %v\n", a, err)
		default:
			time.Sleep(time.Second)

		}
		fmt.Println("here?")
	}()

	time.Sleep(time.Second * 5)

	a := make(chan int)
	go func() { a <- 1 }()
	select {
	case <-a:
		fmt.Println("1")
	case <-c:
		fmt.Println("2")
	}
	close(c)
}
