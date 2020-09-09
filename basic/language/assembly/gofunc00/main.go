package main

//go:noinline
func f1() {
	go f1()
}

//go:noinline
func main() {
	go f1()
}
