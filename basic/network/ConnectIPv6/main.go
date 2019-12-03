package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.DefaultClient
	req, _ := http.NewRequest("GET", "http://ipv6.baidu.com", nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer resp.Body.Close()

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading  %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d : %s", resp.StatusCode, resp.Status)
}
