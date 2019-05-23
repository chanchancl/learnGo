package main

import (
	"fmt"
	"strconv"
	"time"
)

const LOOP = 10000000

func main() {
	fmt.Println("Benchmark.itoa Start!\n")

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
