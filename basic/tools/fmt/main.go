package main

import "fmt"

type std struct {
	a, b int
	c    string
}

func main() {
	fmt.Printf("%#v\n", 1.0)
	var f float64
	f = 356848.0
	fmt.Printf("%#f\n", f)

	fmt.Printf("%#v\n", std{})
}
