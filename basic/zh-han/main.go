package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("Go你好呀")
	p := make([]byte, 100)
	n, err := reader.ReadAt(p, 2)
	p = p[:n]
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s, %b, %d\n", p, p, n)
}
