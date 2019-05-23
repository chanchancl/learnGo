package main

import "fmt"

func main() {
	mp := make(map[string]int)
	mp["123"] = 2

	var p map[string]int

	p = mp

	fmt.Printf("%v\n", p["123"])
	mp["123"] = 3

	fmt.Printf("%v\n", p["123"])
}
