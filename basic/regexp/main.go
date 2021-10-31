package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("abc-([0-9]+)")

	fmt.Println(re.FindStringSubmatch("abc-00666a48858-444"))
}
