package main

import (
	"fmt"
	"net/netip"
)

func main() {

	{
		addr1, _ := netip.ParseAddr("0.0.1.0")

		addr2, _ := netip.ParsePrefix("0.0.0.0/24")

		if addr2.Contains(addr1) {
			fmt.Printf("%v contains %v\n", addr2.String(), addr1.String())
		} else {
			fmt.Printf("%v not contains %v\n", addr2.String(), addr1.String())
		}
	}

	{
		addr1, _ := netip.ParseAddr("192.6.8.1")

		addr2, _ := netip.ParseAddr("192.6.8.2")

		addr3, _ := netip.ParseAddr("192.6.8.3")

		fmt.Println(addr1.Compare(addr2), addr2.Compare(addr3))

		if addr1.Compare(addr2) < 0 && addr2.Compare(addr3) < 0 {
			fmt.Printf("%v is between %v and %v\n", addr2.String(), addr1.String(), addr3.String())
		} else {
			fmt.Printf("%v is not between %v and %v\n", addr2.String(), addr1.String(), addr3.String())
		}
	}

	{
		prefix, _ := netip.ParsePrefix("1122:3344::1122/32")

		addr, _ := netip.ParseAddr("1122:3344:5566::")

		if prefix.Contains(addr) {
			fmt.Printf("%v contains %v\n", prefix.String(), addr.String())
		} else {
			fmt.Printf("%v not contains %v\n", prefix.String(), addr.String())
		}

	}
}
