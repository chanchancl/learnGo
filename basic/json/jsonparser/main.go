package main

import (
	"fmt"

	"github.com/buger/jsonparser"
)

var jsonBytes = []byte(`
{
	"name": "abc",
	"array": [1,2,3],
	"information": {
	  "type": "food",
	  "num": 10.5
	}
}
`)

func main() {

	tp, _ := jsonparser.GetString(jsonBytes, "information", "type")

	num, _ := jsonparser.GetFloat(jsonBytes, "information", "num")

	fmt.Println(tp, num)
}
