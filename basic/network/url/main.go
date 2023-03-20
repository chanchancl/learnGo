package main

import (
	"fmt"
	"net/url"
)

func main() {
	//URL := "https://www.hostname.com:1234/iam/path/"
	URL := "https://user:password@[0011:2233:4455:6677:8899:AABB:CCDD:EEFF]:1004/iam/path/abc?a=10&b=20&c=100&d=你好"
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

	url2 := "https://[0011::EEFF]:1004/abc/def?a=20&c=abc"
	u, err := url.Parse(url2)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(u)
	fmt.Println(u.Hostname())
	fmt.Println(u.Host)
}
