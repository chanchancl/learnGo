package main

import (
	"fmt"
	"net/http"
)

func main() {
	hd := http.Header{}

	hd.Add("wow", "value1")
	hd.Add("wow", "value2")

	hd.Set("moe", "value1")
	hd.Set("moe", "value2")

	fmt.Println(hd.)
}
