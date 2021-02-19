package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	v := uint8(5)

	json.Unmarshal([]byte("abcd"), &v)

	fmt.Println(v)
}
