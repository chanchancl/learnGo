package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("Raw = %4v, Itoa = %4v\n", i, strconv.Itoa(i))
	}
}
