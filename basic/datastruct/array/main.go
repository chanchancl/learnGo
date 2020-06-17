package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			defer wg.Done()
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()

		// slice
		vector1 := []string{}
		vector1[0] = "1"
		// error!!!
	}()

	// array 5
	vector2 := [5]string{}
	vector2[0] = "1"

	fmt.Println(vector2[0])

	wg.Wait()
}
