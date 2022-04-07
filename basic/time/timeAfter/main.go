package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var timer *time.Timer
var lock sync.Mutex

func SetTimer() {
	lock.Lock()
	defer lock.Unlock()
	if timer == nil {
		// Do some work
		fmt.Println("Set timer successful")
		timer = time.AfterFunc(time.Second, func() {
			fmt.Println("Start to do work")
			time.Sleep(time.Millisecond * 500)
			fmt.Println("End to do work")
			lock.Lock()
			defer lock.Unlock()
			fmt.Println("Timer expired and work done")
			if timer == nil {
				fmt.Println("Timer is nil when work done")
			}
			timer = nil
		})
	}
}

func CancelTimer() {
	lock.Lock()
	defer lock.Unlock()
	if timer != nil {
		fmt.Println("Success to cancel timer")
		timer.Stop()
		timer = nil
	}
}

func main() {
	fmt.Println("Test 1 : Just SetTimer")
	fmt.Println()
	SetTimer()
	time.Sleep(time.Second * 2)
	fmt.Println(strings.Repeat("*", 80))

	fmt.Println("Test 2 : Multi SetTimer")
	fmt.Println()
	for i := 0; i < 100000; i++ {
		go SetTimer()
	}
	time.Sleep(time.Second * 2)
	fmt.Println(strings.Repeat("*", 80))

	fmt.Println("Test 3 : Multi CancelTimer")
	fmt.Println()
	SetTimer()
	for i := 0; i < 100000; i++ {
		go CancelTimer()
	}
	time.Sleep(time.Second * 2)
	fmt.Println(strings.Repeat("*", 80))

	fmt.Println("Test 4 : Cancel timer when it expired but doing work")
	fmt.Println()
	SetTimer()
	time.Sleep(time.Second)
	CancelTimer()
	time.Sleep(time.Second)

	fmt.Println(strings.Repeat("*", 80))

	fmt.Println("Test 5 : Cancel timer just a little early than expired, to compare with test 4")
	fmt.Println()
	SetTimer()
	time.Sleep(time.Millisecond * 990)
	CancelTimer()
	time.Sleep(time.Second)

	fmt.Println(strings.Repeat("*", 80))
}
