package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second)

	timestamp, ok := <-timer.C
	fmt.Printf("Expired timer, timestamp : %v, ok : %v\n", timestamp, ok)

	timer = time.NewTimer(time.Second * 10)
	stoped := make(chan int)
	// This will cause deadlock
	// timer.Stop()
	// timestamp, ok = <-timer.C
	// fmt.Printf("Stoped timer, timestamp : %v, ok : %v\n", timestamp, ok)

	go func() {
		time.Sleep(time.Second)
		fmt.Println(timer.Stop())
		stoped <- 0
	}()

	fmt.Println(timer.Stop())
	select {
	case timestamp, ok = <-timer.C:
		fmt.Printf("Expired, timestamp : %v, ok : %v\n", timestamp, ok)
	case <-stoped:
		fmt.Printf("The timer is stoped, timestamp : %v\n", time.Now())
	}
	fmt.Println(timer.Stop())
}
