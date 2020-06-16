package main

import (
	"fmt"
)

const (
	Clear = iota
)

type Test struct {
	a *int
}

func main() {
	var a *int
	if a == nil {
		fmt.Println("nil pointer")
	}
	a = new(int)

	var b *Test
	if b == nil {
		fmt.Println("b is nil pointer without init")
	}
	b = &Test{
		a,
	}
}
