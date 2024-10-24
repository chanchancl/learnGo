package main

import (
	"math/rand"
	"strconv"
)

func main() {
	testmap := make(map[string]int)

	for i := 0; i < 1000000; i++ {
		testmap[strconv.Itoa(rand.Int())] = 10
	}

	for key := range testmap {
		delete(testmap, key)
	}

	println(len(testmap))
}
