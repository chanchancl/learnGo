package main

import "net/http"

func main() {
	response := f1()

	print(response.Header.Get("test"))

}

func f1() *http.Response {
	resp, _ := http.Get("http://www.baidu.com")
	defer tag(resp.Header)
	return resp
}

func tag(header http.Header) {
	header.Set("test", "233")
}
