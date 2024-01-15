package main

import (
	"fmt"
	"net/netip"
	"net/url"
)

func main() {
	start, _ := netip.ParseAddr("192.168.0.1")
	end, _ := netip.ParseAddr("192.168.0.5")

	l, _ := url.Parse("http://192.168.0.3/abc/def.jpg")

	hostname := l.Hostname()

	target, _ := netip.ParseAddr(hostname)

	if IsIPBetween(start, target, end) {
		fmt.Println("Between")
	}

	tests := []struct {
		start, end netip.Addr
		target     netip.Addr
		expect     bool
	}{
		{
			start:  GetAddr("192.168.0.1"),
			end:    GetAddr("192.168.0.5"),
			target: GetAddr("192.168.0.3"),
			expect: true,
		},
		{
			start:  GetAddr("192.168.0.1"),
			end:    GetAddr("192.168.1.1"),
			target: GetAddr("192.168.0.255"),
			expect: true,
		},
		{
			start:  GetAddr("192.168.0.1"),
			end:    GetAddr("192.168.0.1"),
			target: GetAddr("192.168.0.1"),
			expect: true,
		},
		{
			start:  GetAddr("192.168.0.5"),
			end:    GetAddr("192.168.0.1"),
			target: GetAddr("192.168.0.3"),
			expect: false,
		},
		{
			start:  GetAddr("1122::0011"),
			end:    GetAddr("1122::0013"),
			target: GetAddr("1122::0012"),
			expect: true,
		},
	}

	failed := false
	for _, t := range tests {
		got := IsIPBetween(t.start, t.target, t.end)
		if got != t.expect {
			failed = true
			fmt.Printf("Error on (%v %v) and %v, got %v, but expect %v", t.start, t.end, t.target, got, t.expect)
		}
	}
	if !failed {
		fmt.Println("All testcases pass")
	}

}

func GetAddr(ip string) netip.Addr {
	p, _ := netip.ParseAddr(ip)
	return p
}

func IsIPBetween(ip1, ip2, ip3 netip.Addr) bool {
	return ip1.Compare(ip2) <= 0 && ip2.Compare(ip3) <= 0
}
