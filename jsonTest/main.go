package main

import "fmt"

type Test struct {
	a string "abc"
	b string "def"
}

func main() {
	a := Test{"1", "2"}
	fmt.Printf("%v", a)
}
