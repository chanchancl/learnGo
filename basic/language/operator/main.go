package main

import "fmt"

func main() {

	b1 := true && false || false
	b2 := false && false || true
	//   (false) && (false) || true

	fmt.Println(b1)
	fmt.Println(b2)

	if true && false || true {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}
