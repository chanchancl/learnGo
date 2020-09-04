package main

import "fmt"

func withoutArg() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	f1 := withoutArg()
	a := f1()
	b := f1()
	c := f1()

	fmt.Println(a, b, c)
}
