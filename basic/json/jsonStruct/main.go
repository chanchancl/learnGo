package main

import (
	"encoding/json"
	"fmt"
)

type Ta struct {
	Text string `json:"text"`
}

type MyJsonName struct {
	A *Ta    `json:"a,omitempty"`
	B string `json:"b,omitempty"`
	C string `json:"c"`
}

func main() {
	a := MyJsonName{
		A: &Ta{Text: "123"},
	}
	b, _ := json.Marshal(a)
	fmt.Println(string(b))
}
