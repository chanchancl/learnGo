package main

import (
	"fmt"
	"net/url"
)

func main() {
	//URL := "https://www.hostname.com:1234/iam/path/"
	URL := "https://2001:3CA1:010F:001A:121B:0000:0000:0010:1234/iam/path/"
	u, _ := url.Parse(URL)

	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.Path)

	fmt.Println(u.Host)

	fmt.Println(u.Hostname())
	fmt.Println(u.Port())

}
