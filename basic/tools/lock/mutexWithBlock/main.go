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
	fmt.Println("Unlock")

}
