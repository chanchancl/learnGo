package main

import (
	"encoding/json"
	"fmt"
)

type TestStruct struct {
}

func (c *TestStruct) MarshalText() ([]byte, error) {
	return []byte("abc"), nil
}

func main() {
	res, _ := json.Marshal(&TestStruct{})
	fmt.Println(string(res))

	res, _ = json.Marshal(nil)
	fmt.Println(string(res))
}
