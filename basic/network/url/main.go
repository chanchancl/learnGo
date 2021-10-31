package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	//URL := "https://www.hostname.com:1234/iam/path/"
	URL := "https://user:password@2001:3CA1:010F:001A:121B:0000:0000:0010:1234/iam/path/abc?a=10&b=20&c=100&d=你好"
	u, _ := url.Parse(URL)

	fmt.Println(u.Scheme)
	user := u.User
	fmt.Println(user.Username())
	fmt.Println(user.Password())
	fmt.Println(u.Path)

	fmt.Println(u.Host)

	fmt.Println(u.Hostname())
	fmt.Println(u.Port())
	fmt.Println(u.RequestURI())
	fmt.Println(u.String()[:strings.Index(u.String(), "?")])
}
