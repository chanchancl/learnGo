package main

import (
	"fmt"
	"time"
)

func main() {
	total := 0

	testNumber := 1000000
	for i := 0; i < testNumber; i++ {
		start := time.Now()
		// gp := sync.WaitGroup{}
		// gp.Add(2)
		// go func() {
		// 	gp.Done()
		// }()
		// go func() {
		// 	gp.Done()
		// }()
		// gp.Wait()
		end := time.Now()

		total += end.Nanosecond() - start.Nanosecond()
	}
	total /= testNumber
	fmt.Printf("Total time : %v\n", time.Duration(total))
}
