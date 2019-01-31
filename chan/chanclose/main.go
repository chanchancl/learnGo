package main

import "fmt"

func main() {
	datachan := make(chan int, 5)
	syncchan1 := make(chan struct{}, 1)
	syncchan2 := make(chan struct{}, 2)

	go func() {
		<-syncchan1
		for {
			if elem, ok := <-datachan; ok {
				fmt.Printf("Received: %d [receiver]\n", elem)
			} else {
				break
			}
		}
		syncchan2 <- struct{}{}
		fmt.Println("Done. [receiver]")
	}()
	go func() {
		for i := 0; i < 5; i++ {
			datachan <- i
			fmt.Printf("Sent: %d [sender]\n", i)
		}
		close(datachan)
		syncchan1 <- struct{}{}
		fmt.Println("Done. [sender]")
		syncchan2 <- struct{}{}
	}()
	<-syncchan2
	<-syncchan2
}
