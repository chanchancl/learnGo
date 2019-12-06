package main

import "fmt"

func main() {
	var a int
	var b string
	var c float32
	var d rune
	var e chan int
	var f map[int]int
	var g bool
	var h []int

	fmt.Printf("int : %v\n", a)
	fmt.Printf("string : %+v\n", b)
	fmt.Printf("float32 : %v\n", c)
	fmt.Printf("rune : %v\n", d)
	fmt.Printf("chan : %v\n", e)
	fmt.Printf("map : %v\n", f)
	fmt.Printf("bool : %v\n", g)
	fmt.Printf("slice : %v\n", h)

	i := make(chan int)
	j := make(chan int, 1)
	k := make(chan int, 5)
	fmt.Printf("make(chan int) : %v\n", i)
	fmt.Printf("make(chan int, 1) : %v\n", j)
	fmt.Printf("make(chan int, 5) : %v\n", k)

}
