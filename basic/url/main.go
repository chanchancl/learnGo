package main

import (
	"fmt"
	"net/url"
)

func main() {

	URL := "https://www.hostname.com:1234/iam/path/"
	u, _ := url.Parse(URL)

	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.Path)

	fmt.Println(u.Host)

	fmt.Println(u.Hostname())
	fmt.Println(u.Port())

}
