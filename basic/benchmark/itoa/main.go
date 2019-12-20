package main

import (
	"fmt"
	"strconv"
	"time"
)

const LOOP = 100000

func main() {
	fmt.Println("Benchmark.itoa Start!")

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

	for i := 0; i < LOOP; i++ {
		if fmt.Sprintf("%d", i) != strconv.Itoa(i) {
			fmt.Printf("Not Equal")
			return
		}
	}

	fmt.Printf("They are Equal.")
}
