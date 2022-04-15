package main

import (
	"encoding/hex"
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	a := "abcdefghijklmnopqrstuvwxyz012345"
	dump("a", &a)

	b := a[2:4]
	dump("b", &b)

	c := strings.Clone(b)
	dump("c", &c)
}

func dump(name string, str *string) {
	fmt.Println(strings.Repeat("-", 30))

	fmt.Printf("&%s:%p\n", name, str)

	fmt.Println(hex.Dump((*(*[0x10]byte)(unsafe.Pointer(str)))[:]))

	p := *(*int)(unsafe.Pointer(str))
	fmt.Printf("%s:%X\n", name, p)
	fmt.Println(hex.Dump((*(*[0x20]byte)(unsafe.Pointer(uintptr(p))))[:]))
}
