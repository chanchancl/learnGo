package main

import (
	"encoding/json"
	"fmt"
)

type body struct {
	A string `json:"valueA"`
	B string `json:"valueB"`
}

func main() {
	v := uint8(5)

	json.Unmarshal([]byte("abcd"), &v)

	fmt.Println(v)

	// json type
	// null, string, float, bool, []array, {}dict
	b := []*body{
		{"1", "2"},
		{"4", "5"},
	}
	result, _ := json.Marshal(b)

	fmt.Println(string(result))
}
