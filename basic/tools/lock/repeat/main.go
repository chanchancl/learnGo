package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	fmt.Println("Acquire the mutex. (main)")
	mutex.Lock()
	fmt.Println("The lock is acquired. (main)")
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("\tAcquire the mutex. (g%d)\n", i)
			mutex.Lock()
			defer func() {
				fmt.Printf("\tRelease the mutex. (g%d)\n", i)
				mutex.Unlock()
			}()
			fmt.Printf("\tThe lock is acquired. (g%d)\n", i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Release the lock. (main)")
	mutex.Unlock()
	fmt.Println("The lock is released. (main)")
	time.Sleep(time.Second)
}
