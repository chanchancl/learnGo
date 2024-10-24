package main

import (
	"fmt"
	"sync"
)

func main() {
	iarray := []string{"a", "b", "c"}

	// Print c c c in 1.21
	// Print a b c after 1.21

	wg := sync.WaitGroup{}
	done := make(chan bool)
	for i, k := range iarray {
		wg.Add(1)
		go func() {
			<-done
			fmt.Println(i, k, &i, &k)
			wg.Done()
		}()
	}

	for _ = range iarray {
		done <- true
	}
	wg.Wait()
}
