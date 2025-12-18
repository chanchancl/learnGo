package main

import (
	"fmt"
	"slices"
)

func main() {
	iter := func(yield func(int) bool) {
		for i := range 10 {
			if !yield(i) {
				break
			}
		}
	}

	for i := range iter {
		fmt.Println(i)
	}

	for i := range slices.Backward(slices.Collect(iter)) {
		fmt.Println(i)
	}
}
