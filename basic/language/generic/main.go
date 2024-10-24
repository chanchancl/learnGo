package main

import "fmt"

func Index[S ~[]E, E comparable](v S, e E) int {
	for i := range v {
		if v[i] == e {
			return i
		}
	}
	return -1
}

func IndexWithoutQuote[S []E, E comparable](v S, e E) int {
	for i := range v {
		if v[i] == e {
			return i
		}
	}
	return -1
}

func EXCEPT_EQ(a, b interface{}) {
	if a != b {
		fmt.Printf("Expect %v, but got %v\n", a, b)
	}
}

type MySlice []int

func main() {
	ia := []int{1, 2, 3}

	EXCEPT_EQ(Index(ia, 5), -1)
	EXCEPT_EQ(Index(ia, 2), 1)

	ib := MySlice{1, 2, 3}

	// Compile error
	// EXCEPT_EQ(IndexWithoutQuote(ib, 10), -1)
	EXCEPT_EQ(Index(ib, 10), -1)

}
