package main

import (
	"fmt"
	"time"
)

type gate chan bool

func (g gate) enter() { g <- true }
func (g gate) leave() { <-g }

func (g gate) Len() int { return len(g) }
func (g gate) Cap() int { return cap(g) }

func (g gate) Idle() bool { return len(g) == 0 }
func (g gate) Busy() bool { return len(g) == cap(g) }

func (g gate) Fraction() float64 {
	return float64(len(g)) / float64(len(g))
}

type Data struct {
	data int
	gate
}

func NewData(n int) *Data {
	return &Data{gate: make(chan bool, n)}
}

func (d *Data) Add() {
	d.enter()
	defer d.leave()
	d.data += 1
}

func main() {
	data := NewData(1)
	for i := 0; i < 10; i++ {
		go data.Add()
	}
	time.Sleep(time.Second)
	fmt.Println(data.data)
}
