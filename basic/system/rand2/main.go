package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int64())
	}

	seed := [32]byte{}

	cc := rand.NewChaCha8(seed)

	fmt.Println("")
	for i := 0; i < 10; i++ {
		fmt.Println(cc.Uint64())
	}
}
