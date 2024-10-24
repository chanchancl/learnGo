package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	ia := []int{42, 15, 21, 10, 4, 99}

	for i := range slices.Values(ia) {
		fmt.Println(i)
	}

	for i, v := range slices.All(ia) {
		fmt.Println(i, v)
	}

	imap := map[string]int{
		"a": 2,
		"b": 3,
		"c": 4,
	}

	for key := range maps.Keys(imap) {
		fmt.Println(key)
	}

	maps.Keys(imap)(func(v string) bool {
		fmt.Println(v)
		return true
	})
}
