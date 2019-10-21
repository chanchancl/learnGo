package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("Go你好呀")
	p := make([]byte, 3)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %b, %d\n", p, p, n)
}
