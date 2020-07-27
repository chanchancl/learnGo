package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abaca"
	sp := strings.Split(s, "a")
	fmt.Printf("%#v, %v\n", sp, len(sp))

	s = strings.Join(sp, "a")
	fmt.Println(s)
}