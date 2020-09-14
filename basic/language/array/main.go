package main

import "fmt"

func main() {
	var array [256]int
	fmt.Println(len(array))
	fmt.Println(cap(array))

	var slice []int
	slice = array[:10]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
