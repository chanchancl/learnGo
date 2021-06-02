package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now())
	go func(t time.Time) {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(t)
		fmt.Println(time.Now())
	}(time.Now())

	// arguments in go call will be evaluated immediately

	time.Sleep(time.Second)
}
