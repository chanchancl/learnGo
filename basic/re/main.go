package main

import (
	"fmt"
	"net"
	"regexp"
)

func matchIP(ip string) bool {
	patterns := []string{
		"^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$",
		"^((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))$",
		"^(([^:]+:){6}(([^:]+:[^:]+)|(.*\\..*)))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?)(%.+)?$",
	}

	for _, pattern := range patterns {
		match, err := regexp.MatchString(pattern, ip)
		if err == nil && match {
			return true
		}
	}

	return false
}

func main() {
	IPs := []string{
		"1.1.1.1.",
		"0.0.0.0",
		":",
		"123",
		"::",
		"1:1:1:1",
		"1:2:3:45678",
	}

	for _, ip := range IPs {
		if matchIP(ip) {

			fmt.Printf("Matach %s. ", ip)
		} else {
			fmt.Printf("Doesn't match %s. ", ip)
		}
		fmt.Printf("net.ParseIP : %v\n", net.ParseIP(ip))
	}

}
