package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	source := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range source {
			fmt.Println(i)
		}
		fmt.Println("source chan is closed, 溜了")
	}()

	for i := 0; i < 3; i++ {
		fmt.Printf("准备发送%v\n", i)
		source <- i
	}

	time.Sleep(1)
	source <- 4

	close(source)
	wg.Wait()

}
