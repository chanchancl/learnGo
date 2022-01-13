package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("Start : %v\n", time.Now())
	go func(t time.Time) {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Input time in goroutine : %v\n", t)
		fmt.Printf("Goroutine now : %v\n", time.Now())
	}(time.Now())

	defer fmt.Printf("Time in defer : %v\n", time.Now()) // argument will passed to defer immediately

	// arguments in go call will be evaluated immediately

	time.Sleep(time.Second)
}
