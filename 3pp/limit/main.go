package main

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
	rl := ratelimit.New(100)

	prev := time.Now()
	total := time.Duration(0)
	N := 100
	for i := 0; i < N; i++ {
		now := rl.Take()
		diff := now.Sub(prev)
		total += diff
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
	fmt.Printf("Average : %3v\n", total/time.Duration(N-1))
}
