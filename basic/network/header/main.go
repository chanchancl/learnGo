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

	fmt.Println(hd)
	req, _ := http.NewRequest(http.MethodGet, "test.com", nil)

	req.Header.Set("x-abc-aaa", "aa bb cc")
	req.Header.Set("x-abc-bbb", "qq ww ee")

	for key := range req.Header {
		fmt.Println(key)
		// X-Abc-Aaa
		// X-Abc-Bbb
		// So get header from http.Header, must focus the Upper/Lower case
	}

}
