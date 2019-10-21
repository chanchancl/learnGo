package main

import "runtime"

func main() {
	go println("Go! Goroutine!")
	runtime.Gosched()
}
