package main

import "fmt"

type Mock struct {
	a map[int]int
}

func main() {
	mock := &Mock{}

	// use directly
	// mock.a[1] = 1
	// fmt.Printf("%v", mock)

	mock.a = make(map[int]int)
	mock.a[1] = 1
	mock.a[2] = 2
	fmt.Printf("%v\n", mock)

}
