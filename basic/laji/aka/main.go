package main

import "fmt"

func f(data []byte) {
	if data == nil {
		fmt.Println("data is nil")
	}
}

func main() {
	fmt.Println("laji!")

	data := []byte(nil)
	f(data)

	a := make(chan []byte)

	go func() {
		a <- nil
	}()

	fmt.Println(<-a)
}
