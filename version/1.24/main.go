package main

type MyMap[K comparable, V any] = map[K]V

func main() {
	mymap := make(MyMap[int, string])
	mymap[1] = "Hello"
}
