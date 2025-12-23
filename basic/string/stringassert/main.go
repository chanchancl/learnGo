package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	FieldWithInterface   interface{} `json:"abc"`
	FieldWithArrayString []string    `json:"qwe"`
}

func test(a interface{}) {
	switch s := a.(type) {
	case int:
		fmt.Println("int ", s)
	case string:
		fmt.Println("string ", s)
	case []string:
		fmt.Println("[]string ", s)
	case []interface{}:
		result := []string{}
		for _, v := range s {
			result = append(result, v.(string))
		}
		fmt.Println("[]interface{} : ", result)
	}
}

func main() {
	str := []byte(`{"abc": ["1", "2"], "qwe": ["1", "2", "3"]}`)

	rawString := []string{"q", "w", "e"}

	a := A{}
	json.Unmarshal(str, &a)
	test(a.FieldWithInterface)
	test(a.FieldWithArrayString)
	test(rawString)
}
