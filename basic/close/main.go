package main

import "fmt"

func out(cl chan string) {
	fmt.Println(cl)
}
func main() {
	var cl chan string
	out(cl)
	cl = make(chan string)
	out(cl)
	close(cl)
	out(cl)
	cl = nil
	out(cl)
}
