package main

import "fmt"

const (
	a = -1
	b = -2
	c = iota
	d
	e = -3
	f
	g = "abc"
	h
	i
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
}
