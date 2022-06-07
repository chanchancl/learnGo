package main

import "fmt"

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s *S) Write(str string) {
	s.data = str
}

func f1() {
	sVals := map[int]S{1: {"A"}}

	sVals[1].Read()

	// 无法通过编译，因为 值无法调用 带指针的接收器
	// sVals[1].Write("test")

	sPtrs := map[int]*S{2: {"B"}}

	// 但指针可以调用 值接收器 和 指针接收器
	sPtrs[2].Read()
	sPtrs[2].Write("test")

	fmt.Println(sVals[1], sPtrs[2])
}

type F interface {
	f()
}

type S1 struct{}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}

func f2() {
	s1Val := S1{}
	s1Ptr := &S1{}
	s2Val := S2{}
	s2Ptr := &S2{}

	var i F
	i = s1Val
	i = s1Ptr
	// 下面这行无法编译，因为类型S2，只有指针接收器，无法赋值值对象，但 i = &s2Val 就可以
	// i = s2Val
	i = &s2Val
	i = s2Ptr

	fmt.Println(i, s2Val)

	/*
		总的来说，一个类型T，既可以实现值接收器，也可以实现指针接收器
		值接收器是指针接收器的子集，也就是说，如果类型T实现了指针接收器，它既可以接收指针，也可以接受值（通过将指针解引用）
		但反之并不成立，实现了值接收器的类型，不能接收指针

		值对象只能使用值接收器
		指针可以使用 值接收器 + 指针接收器

		如果一个类型实现了值接收器，那么  对应的接口，可以复制 值对象，也可以赋值 指针
		但若是只实现了指针接收器，那么  对应的接口只可以赋值指针
	*/
}

func main() {
	f1()
	f2()
}
