package main

import (
	"fmt"
	"net/netip"
)

// Generic function with type parameter
func Equal[T comparable](a, b T) bool {
	return a == b
}

func f1() {
	fmt.Println(Equal(1+2, 2+1))
	fmt.Println(Equal(1, 2))
	fmt.Println(Equal("a", "b"))
	fmt.Println(Equal(1.0, 2.0))
}

// Generic function with interface constraint
type Stringer interface {
	String() string
}

func Print[T Stringer](s T) {
	println(s.String())
}

type MyString string
type MyInt int

func (s MyString) String() string { return string(s) }
func (i MyInt) String() string    { return fmt.Sprintf("%d", i) }

func f2() {
	Print(MyString("hello"))
	// Print("world") // string does not satisfy Stringer (missing method String)
	Print(MyInt(42))

	// addr have implemented Stringer interface
	addr, _ := netip.ParseAddrPort("192.168.0.1:8080")
	Print(addr)
}

// Generic function with int like type constraint
func Sum[T ~int](nums []T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}

// witout ~
func SumInt[T int](nums []T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}

func f3() {
	fmt.Println(Sum([]int{1, 2, 3}))
	fmt.Println(Sum([]MyInt{1, 2, 3}))

	fmt.Println(SumInt([]int{1, 2, 3}))
	// fmt.Println(SumInt([]MyInt{1, 2, 3})) // cannot use []MyInt as type []int in argument to SumInt
}

type Number interface {
	int | float64
}

func Add[T Number](a, b T) T {
	return a + b
}

func f4() {
	fmt.Println(Add(1, 2))
	fmt.Println(Add(1.0, 2.0))
	// fmt.Println(Add(1, 2.0)) // invalid operation: 1 + 2.0 (mismatched types int and float64)
}

// Accept a slice with element type E
func FirstElement[S ~[]E, E any](s S) E {
	if len(s) == 0 {
		var zero E
		return zero
	}
	return s[0]
}

func f5() {
	fmt.Println(FirstElement([]int{1, 2, 3}))
	fmt.Println(FirstElement([]string{"a", "b", "c"}))
	fmt.Println(FirstElement([]MyInt{4, 5, 6}))
}

type Pair[T any, U any] struct {
	First  T
	Second U
}

func f6() {
	p := Pair[int, string]{42, "hello"}
	fmt.Println(p)
}

func main() {
	funcs := []func(){f1, f2, f3, f4, f5, f6}

	for i, f := range funcs {
		fmt.Printf("In f%v\n", i+1)
		f()
		fmt.Println("******************************")
	}
}
