package main

import "fmt"

func main() {
	// 内置 min, max, clear

	fmt.Println(min(1, 2, 3))
	fmt.Println(max(1, 2, 3))

	mp := make(map[int]string)
	mp[1] = "hello"
	clear(mp)

	fmt.Println(mp)
}
