package main

import (
	"fmt"
	"math/rand"
)

func size(s int) int {
	return (s + 7) &^ 7
}

func main() {
	for s := 1; s < 100000; s <<= 1 {
		fmt.Printf("%5v , %5v\n", s, size(s))
	}
	fmt.Println("/////////////////////")
	for i := 0; i < 10; i++ {
		s := rand.Int() % 2049
		fmt.Printf("%4v , 0x%X -> %4v , 0x%X\n", s, s, size(s), size(s))
	}
	// 8 Byte allign, up to
}
