package main

import (
	"fmt"
	"sync"
)

func main() {
	case1()
}

type sa struct {
	InternalMutex sync.Mutex
}

func case1() {
	a := sa{}
	// a := sa{InternalMutex: new(sync.Mutex)}
	b := a

	fmt.Println(a.InternalMutex)
	fmt.Println(b.InternalMutex)

	a.InternalMutex.Lock()
	b.InternalMutex.Lock()
}
