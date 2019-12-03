package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

func main() {
	bd := strings.Builder{}
	for i := 0; i < 100; i++ {
		bd.WriteString("A")
	}
	hex1 := bd.String()
	bd.Reset()
	for i := 0; i < 100; i++ {
		bd.WriteString("B")
	}
	hex2 := bd.String()

	begun := time.Now()
	buf := bytes.Buffer{}
	buf.WriteString(hex1)
	buf.WriteString(hex2)
	var dst1 []byte
	fmt.Sscanf(buf.String(), "%x", &dst1)
	pass1 := time.Since(begun).Nanoseconds()

	begun = time.Now()
	bufs := strings.Builder{}
	bufs.WriteString(hex1)
	bufs.WriteString(hex2)
	var dst2 []byte
	fmt.Sscanf(bufs.String(), "%x", &dst2)
	pass2 := time.Since(begun).Nanoseconds()
	a := sha256.New()
	a.Write(dst2)
	fmt.Printf("%#v %v\n%#v %v\n", dst1[:10], pass1, dst2[:10], pass2)
	fmt.Printf("%x\n", a.Sum(nil))
}
