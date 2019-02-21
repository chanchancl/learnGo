package main

import (
	"fmt"
	"strconv"
	"time"
)

const LOOP = 1000000

func main() {
	startTime := time.Now()
	for i := 0; i < LOOP; i++ {
		fmt.Sprintf("%d", i)
	}
	fmt.Printf("fmt.Sprintf taken: %v\n", time.Since(startTime))

	startTime = time.Now()
	for i := 0; i < LOOP; i++ {
		strconv.Itoa(i)
	}
	fmt.Printf("strconv.FormatInt taken: %v\n", time.Since(startTime))
}
