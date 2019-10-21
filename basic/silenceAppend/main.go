package main

import "fmt"

func testSlice() {
	a := make([]int, 0, 10)
	a = append(a, 10, 11)
	fmt.Printf("%v, %v\n", a[0], a[1])
}

func main() {
	var v []int
	for i := 0; i < 20; i++ {
		v = append(v, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(v), v)
	}

	testSlice()
}
