package main

import "fmt"

func main() {
	a := "233"
	switch a {
	case "233":
		fmt.Println("233")
	case "123":
		fmt.Println("123")
	default:
		fmt.Println("111")
	}
}
