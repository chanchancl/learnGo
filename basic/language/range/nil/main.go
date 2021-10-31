package main

import "fmt"

func main() {
	slice := []string{}

	for i := range slice {
		fmt.Println(slice[i])
	}
}
