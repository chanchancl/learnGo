package main

import (
	"fmt"
	"sync"
)

// 不要把 sync.Mutex，嵌入到结构体中

// bad
type SMap struct {
	sync.Mutex // 嵌入
	data       map[string]string
}

// good
type SMapgood struct {
	mu   sync.Mutex
	data map[string]string
}

func NewSMap() *SMap {
	return &SMap{
		data: make(map[string]string),
	}
}
func NewSMapgood() *SMapgood {
	return &SMapgood{
		data: make(map[string]string),
	}
}

func (c *SMap) Get(key string) string {
	c.Lock()
	defer c.Unlock()
	return c.data[key]
}

func (c *SMapgood) Get(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.data[key]
}

func main() {
	i := 0

	// bad
	mu := new(sync.Mutex)
	mu.Lock()
	i += 1
	mu.Unlock()

	// good
	muu := sync.Mutex{}
	muu.Lock()
	i += 1
	muu.Unlock()

	// bad
	map1 := NewSMap()
	// sync.Mutex 的 Lock 方法被暴露了
	map1.Lock()
	i += 1
	map1.Unlock()

	fmt.Println(i)
}
