package main

import (
	"fmt"
	"time"
)

func main() {
	closeChan := make(chan struct{})
	finished := make(chan struct{})

	start := time.Now()
	go func() {
		defer func() { finished <- struct{}{} }()
		select {
		case <-time.Tick(5 * time.Second):
			fmt.Println("Tick is timeout")
		case _, ok := <-closeChan:
			// The closed chan will read a message (zero-value, false) immediately
			if !ok {
				fmt.Println("Not ok")
			}
			fmt.Println("Closechan is closed")
		}
	}()

	time.AfterFunc(time.Second, func() {
		close(closeChan)
	})

	<-finished
	eplased := time.Now().Sub(start)
	fmt.Printf("%v\n", eplased)
}
