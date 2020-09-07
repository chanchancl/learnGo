package main

func f1(a, b int) (int, int) {
	return a + b, a - b
}

//go:noinline
func main() {
	f1(1, 2)

	f1(2, 3)
}
