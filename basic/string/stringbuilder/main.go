package main

import (
	"fmt"
	"strings"
)

func main() {
	a := strings.Builder{}
	a.WriteString("qwe")

	// copy, Don't copy strings.Builder
	b := a
	b.WriteString("abc")
	fmt.Println(a.String())
	fmt.Println(b.String())
}
