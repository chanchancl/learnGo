package main

import (
	"errors"
	"fmt"
)

func f1() {
	defer func() {
		// recover must be called in a defer func
		if err := recover(); err != nil {
			fmt.Println(err.(error).Error())
		}
	}()
	// defer recover()
	// recover()
	panic(errors.New("233"))
}

func main() {

	f1()
	fmt.Println("Recovered f1")
}
