package main

import "fmt"

type Animal interface {
	Color()
}

type Dog {
}

func (d Dog) Color() {
	fmt.Println("Wang!")
}

type value struct {
	a int
}

func (t value) valueFunctionAdd1() {
	t.a += 1
}

func (t *value) pointFunctionAdd1() {
	t.a += 1
}

func main() {
	a := value{1}
	b := &value{1}

	fmt.Printf("a : %+v\n", a)
	fmt.Printf("b : %+v\n", b)

	a.valueFunctionAdd1()
	b.valueFunctionAdd1()

	fmt.Println("")
	fmt.Printf("a : %+v\n", a)
	fmt.Printf("b : %+v\n", b)

	a.pointFunctionAdd1()
	b.pointFunctionAdd1()

	fmt.Println("")
	fmt.Printf("a : %+v\n", a)
	fmt.Printf("b : %+v\n", b)
}
