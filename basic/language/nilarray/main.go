package main

import "fmt"

func main() {
	var a []string
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", len(a))

	a = nil
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", len(a))
}