package main

import (
	"fmt"
	"sync"
	"testing"
)

type Data struct {
	Value int
}

func createData() *Data {
	return &Data{Value: 42}
}

var pool = sync.Pool{
	New: func() interface{} {
		return &Data{}
	},
}

func createDataFromPool() *Data {
	obj := pool.Get().(*Data)
	obj.Value = 42
	return obj
}

func TestWithoutPool(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		obj := createData() // Allocating a new object every time
		_ = obj             // Simulate usage
	}
}

func main() {
	for i := 0; i < 1000000; i++ {
		obj := createData() // Allocating a new object every time
		_ = obj             // Simulate usage
	}
	fmt.Println("Done")
}
