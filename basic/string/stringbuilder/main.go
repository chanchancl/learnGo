package main

import "strings"

func main() {
	a := strings.Builder{}
	a.WriteString("qwe")

	// copy
	b := a
	b.WriteString("abc")
}
