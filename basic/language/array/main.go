package main

import (
	"fmt"
)

func main() {
	var array [256]int
	fmt.Println(len(array))
	fmt.Println(cap(array))

	s := array[2:]
	fmt.Printf("Address of array %p\n", &array)
	fmt.Printf("Address of s     %p\n", &s)

	var slice []int
	slice = array[:10]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	tslice := make([]int, 10)[:0]
	tslice = append(tslice, 5)
	fmt.Println(tslice)

	a := []int{1, 2, 3, 4, 5}

	b := make([]int, len(a))
	copy(b, a)
	fmt.Println(a, b)
}
