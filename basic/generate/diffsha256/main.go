package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	a := sha256.Sum256([]byte("1"))
	b := sha256.Sum256([]byte("2"))
	var count int
	for i := range a {
		c := a[i] ^ b[i]
		count += bitCount(c)
	}
	fmt.Printf("%b\n%b\n%v\n", a, b, count)
}

func bitCount(x byte) int {
	count := 0
	for x != 0 {
		if (x & 1) == 1 {
			count++
		}
		x >>= 1
	}
	return count
}
