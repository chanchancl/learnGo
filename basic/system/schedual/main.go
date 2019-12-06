package main

import (
	"fmt"
	"time"
)

func main() {
	total := 0

	testNumber := 3
	start := time.Now()
	for i := 0; i < testNumber; i++ {
		end := time.Now()
		total += end.Nanosecond() - start.Nanosecond()
		//time.Sleep(time.Second)
		start = time.Now()
	}
	total /= testNumber
	fmt.Printf("%v\n", time.Duration(total))
}
