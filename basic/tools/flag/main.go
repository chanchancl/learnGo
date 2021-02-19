package main

import (
	"flag"
	"fmt"
)

var (
	a = flag.Bool("abc", true, "help text")
	b = flag.Int("def", 100, "help text")
)

func main() {
	flag.Parse()
	fmt.Println(*a)
	fmt.Println(*b)
}
