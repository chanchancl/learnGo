package main

import (
	"fmt"

	"github.com/buger/jsonparser"
)

type Test struct {
	a string `abc`
	b string `def`
}

func main() {
	a := Test{"1", "2"}
	fmt.Printf("%v\n", a)

	jsonstr := "{\"abc\":true}"
	result, _ := jsonparser.GetBoolean([]byte(jsonstr), "abc")
	fmt.Println(result)
}
