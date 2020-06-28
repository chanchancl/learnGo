package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func main() {
	sum := sha256.New()
	sum.Write([]byte("really?"))
	fmt.Printf("%x", sum.Sum(nil))
}

func f3() {
	a := make([]byte, 10)
	b := []byte(string("abc"))
	copy(a, b)
	fmt.Println(a, b)
}

func f2() {
	reader := strings.NewReader("abcdefg")
	b := make([]byte, 10)
	fmt.Println(len(b), cap(b))
	b = b[:10]
	fmt.Println(len(b), cap(b))
	n, err := reader.Read(b)
	fmt.Println(n, b, err)
	b = b[:n]
	fmt.Println(n, b, err)
}

func f1() {
	var keyBuf []byte
	kaut := "0001020304050607080910ffaa00"
	fmt.Sscanf(kaut, "%x", &keyBuf)
	fmt.Printf("%#v\n", keyBuf)
}
