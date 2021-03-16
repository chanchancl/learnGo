package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	fmt.Println("--------------------Case 1---------------------")
	go func() {
		defer fmt.Println("Exit select")
		fmt.Println("Enter select")
		select {
		case <-ch:
			fmt.Println("Select!")
		}
	}()

	fmt.Println("Sleep 1 second")
	time.Sleep(time.Second)
	fmt.Println("Sleep end")
	ch <- 5

	time.Sleep(time.Second)
	fmt.Println("--------------------Case 2---------------------")
	go func() {
		defer fmt.Println("Exit select")
		fmt.Println("Enter select")
		select {
		case <-ch:
			fmt.Println("Select!")
		default:
			fmt.Println("Default!")
		}
	}()

	fmt.Println("Sleep 1 second")
	time.Sleep(time.Second)
	fmt.Println("Sleep end")
	// ch <- 5
}
