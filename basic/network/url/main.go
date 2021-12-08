package main

import (
	"fmt"
	"net/url"
)

func main() {
	//URL := "https://www.hostname.com:1234/iam/path/"
	URL := "https://user:password@2001:3CA1:010F:001A:121B:0000:0000:0010:1234/iam/path/abc?a=10&b=20&c=100&d=你好"
	u, _ := url.Parse(URL)
	user := u.User

	println(URL)
	fmt.Printf("Schema     : %v\n", u.Scheme)
	fmt.Printf("Username   : %v\n", user.Username())
	password, _ := user.Password()
	fmt.Printf("Password   : %v\n", password)
	fmt.Printf("Host       : %v\n", u.Host)
	fmt.Printf("Hostname   : %v\n", u.Hostname())
	fmt.Printf("Port       : %v\n", u.Port())
	fmt.Printf("Path       : %v\n", u.Path)
	fmt.Printf("RequestURI : %v\n", u.RequestURI())
	for i, v := range u.Query() {
		fmt.Printf("Query    %v : %v\n", i, v[0])
	}
}
