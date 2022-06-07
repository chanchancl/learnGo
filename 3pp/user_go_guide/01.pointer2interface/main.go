package main

import "fmt"

type F interface {
	f()
}

type S1 struct {
	data int
}

func (s S1) f() {
	// 此处 s 是 main.f1 的拷贝，所以修改s的值只影响函数内
	fmt.Printf("\tf1 in f : %p\n", &s)
	s.data = 10
}

type S2 struct {
	data int
}

func (s *S2) f() {
	// 此处 s 与 main.f2 指向同一个对象
	fmt.Printf("\tf2 in f : %p\n", s)
	s.data = 10
}

func main() {
	f1 := S1{}
	f2 := S2{}
	fmt.Printf("f1 in stack : %p\n", &f1)
	fmt.Printf("f2 in stack : %p\n", &f2)
	f1.f()
	f2.f()

	fmt.Println(f1, f2)
}
