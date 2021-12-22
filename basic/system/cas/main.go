package main

import (
	"fmt"
	"sync/atomic"
)

var (
	inited int32
)

func main() {
	f()
	release()

	f()
	f()
	release()

	f()
	release()
	release()

	release()
	release()
}

func f() {
	if swaped := atomic.CompareAndSwapInt32(&inited, 0, 1); !swaped {
		fmt.Println("Have Inited")
		return
	}

	fmt.Println("Inited!!!!!!!!!!!")
}

func release() {
	if swaped := atomic.CompareAndSwapInt32(&inited, 1, 0); !swaped {
		fmt.Println("Have released")
		return
	}
	fmt.Println("Release!!!!!!!!!!")
}
