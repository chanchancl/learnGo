package main

import "fmt"

type Callback func()

type Impl struct {
	a int
}

func (c *Impl) Do() {
	fmt.Println(c.a)
}

func main() {
	impl := Impl{10}

	// assign struct method to common function type
	var callback Callback = impl.Do

	callback()
}
