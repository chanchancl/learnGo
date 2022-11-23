package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:5999", nil)

	client := http.Client{}

	rsp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bt, err := io.ReadAll(rsp.Body)
	fmt.Println(string(bt))
}
