package main

import (
	"fmt"
	"strings"
)

func remoteEndpoint(endpoint string) string {
	split := "http://"
	if strings.Contains(endpoint, "https://") {
		split = "https://"
	}
	str := strings.Split(endpoint[len(split):len(endpoint)], "/")
	for k, v := range str {
		if k == 0 {
			return v
		}
	}
	return "-"
}

func remoteHostnameAndPort(endpoint string) (string, string) {
	endpoint = remoteEndpoint(endpoint)

	if strings.Contains(endpoint, ":") {
		str := strings.Split(endpoint, ":")
		if len(str) > 1 {
			return str[0], str[1]
		}
	}
	return endpoint, "80"
}

func main() {
	url := []string{
		"http://123.com:23333/uri/aaaa",
		"https://test/uri/abc",
		"http://123.com:1222",
		"http://123.com",
		"http://123.com/abc",
		"http://test:8080/aaa/bbb/ccc",
	}
	for _, v := range url {
		endpoint := remoteEndpoint(v)
		hostname, port := remoteHostnameAndPort(v)
		fmt.Printf("url : %-30s, endpoint : %-20s, hostname %-10s, port : %-10s\n", v, endpoint, hostname, port)
	}

}
