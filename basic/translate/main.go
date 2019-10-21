package main

import "fmt"

func main() {
	data := []byte{100, 100, 100}
	str := string(data)
	data2 := []byte(str)
	fmt.Println(data)
	fmt.Println(str)
	fmt.Println(data2)
}
