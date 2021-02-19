package main

import (
	"fmt"
	"strings"
)

func main() {
	ffff := "12345"
	abc := "abc-12345678910"
	qqqqq := len(abc)
	llll := len(ffff)
	index := len("abc-") + llll
	abc = abc[:index] + strings.Repeat("0", qqqqq-index)
	fmt.Println(abc)
}
