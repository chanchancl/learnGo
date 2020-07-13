package main

import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
)

func main() {
	{
		mutex.Lock()
		defer mutex.Unlock()
		fmt.Println("Lock")
	}
	mutex.Lock()
	fmt.Println("Unlock")
	// This prove that block don't trigger the defer function
	// mutex is locked during the function
	// So don't use this
}
