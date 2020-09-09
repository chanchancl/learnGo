package main

//go:noinline
func f1(a int) int {
	return a
}

//go:noinline
func main() {
	go f1(10)
}
