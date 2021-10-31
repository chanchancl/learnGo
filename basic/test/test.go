package main

import (
	"crypto/sha256"
	"fmt"
	"runtime"
	"strings"
	"sync"
)

func f9() {
	a := []int{0, 1, 2, 3, 4, 5}

	removeIndex := []int{0, 2, 5}

	for i := range removeIndex {

		b := make([]int, 6)
		copy(b, a)

		fmt.Println("||||||||||||||||||||||||||||||||||||||")
		index := removeIndex[i]
		fmt.Println(index)
		fmt.Println(b[:index])
		fmt.Println(b[index+1:])
		fmt.Println(append(b[:index], b[index+1:]...))
	}
}

func f() (int, int) {
	return 10, 10
}

func ff() (int, int) {
	return 20, 20
}

func f10() {
	initValue := runtime.GOMAXPROCS(2)

	fmt.Println(initValue)
	fmt.Println(runtime.NumCPU())
}

var i = 0

func f11f() *int {
	defer func() { i += 1 }()
	if i < 5 {
		return nil
	}
	return &i
}

func f11() {
	a := f11f()
	for a == nil {
		fmt.Println(i)
		a = f11f()
	}
}

func main() {
	f11()
}

func getpc(n int) {
	pc, file, line, ok := runtime.Caller(0)
	fmt.Printf("%x\n", pc)
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(ok)
	fmt.Println("****************************")
	if n > 0 {
		getpc(n - 1)
	}
}

func removeTokenURLQuery(s string) string {
	// Post "https://hostname:port/path/token?query": error
	start := strings.Index(s, "token")
	end := strings.LastIndex(s, "\":")
	if start == -1 || end == -1 {
		return s
	}
	str := s[:start+len("token")] + s[end:]
	return strings.ReplaceAll(str, "\"", "")
}

func f8() {
	s := "Post \"https://abcdefgtoken?asdbdsa315fe3sa23df\": context deadline exceeded (Client.Timeout exceeded while awaiting headers)"
	fmt.Println(removeTokenURLQuery(s))
}

func f7() {
	str := "abc def ghi"
	str = strings.Replace(str, "abc", "ABC", 1)
	str = strings.Replace(str, "def", "DEF", 1)
	str = strings.Replace(str, "ghi", "GHI", 1)
	fmt.Println(str)

	sp := strings.Split(str, " ")
	fmt.Printf("%#v\n", sp)
}

func f6() {
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
