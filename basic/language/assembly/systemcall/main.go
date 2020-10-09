package main

import "fmt"

var (
	c = make(chan int)
)

func f() {
	fmt.Println("233")
	c <- 0
}

func main() {
	go f()
	<-c
}
