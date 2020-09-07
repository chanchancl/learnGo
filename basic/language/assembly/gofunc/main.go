package main

//go:noinline
func f1(a, b int) (int, int) {
	return a + b, a - b
}

//go:noinline
func main() {
	go f1(10, 11)
}
