package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	//netError := net.Error{}
	urlError := url.Error{}

	var netError net.Error
	_, ok := netError.(error)

	netError = urlError
	if _, ok := urlError.(net.Error); ok {
		fmt.Println("Translate url.Error to net.Error")
	}
}
