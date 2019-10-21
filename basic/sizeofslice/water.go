package main

import "fmt"

func main() {
	var fish []int
	for i := 0; i < 20; i++ {
		fish = append(fish, i)
		fmt.Printf("len: %d, cap: %d\n", len(fish), cap(fish))
	}
}
