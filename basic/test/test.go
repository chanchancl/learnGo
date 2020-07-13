package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"sync"
)

func main() {
	var s sync.Mutex
	{
		s.Lock()
		defer s.Unlock()
		print("233")
	}
	fmt.Println("Try to lock")
	s.Lock()
	fmt.Println("Lock success")
	s.Unlock()
	fmt.Print("Unlock")
}

func f5() {
	str := "abcdef015.ghi234.abcdefghijk.abc"
	Mnc := str[6:9]
	Mcc := str[13:16]
	fmt.Println(Mnc, Mcc)
}

func f4() {
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
