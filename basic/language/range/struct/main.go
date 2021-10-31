package main

import "fmt"

type test struct {
	a int
}

func main() {
	tests := []*test{
		{1}, {2}, {3},
	}

	tt2 := []*test{}
	for _, t := range tests {
		fmt.Printf("%p\n", t)
		tt2 = append(tt2, t)
	}

	for _, t := range tt2 {
		fmt.Printf("%p\n", t)
		fmt.Println(t)
	}

}
