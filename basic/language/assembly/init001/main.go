package main

import "fmt"

func f1() {
	var sum int
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("%v\n", sum)
}

func f2() {
	v := []int{}
	for i := 0; i < 5; i++ {
		v = append(v, i)
	}
	fmt.Printf("%v\n", v)
}

//go:noinline
func main() {
	f1()
	f2()
}
