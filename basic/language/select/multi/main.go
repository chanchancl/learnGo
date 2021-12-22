package main

import (
	"time"
)

func main() {
	chs := []chan int{}
	for i := 0; i < 10; i++ {
		ch := make(chan int)
		chs = append(chs, ch)
		local := i
		go func() {
			for {
				ch <- local
			}
		}()
	}
	// select 时，如果有多个case同时满足条件，则随机选择一个
	for {
		select {
		case i := <-chs[0]:
			println(i)
		case i := <-chs[1]:
			println(i)
		case i := <-chs[2]:
			println(i)
		case i := <-chs[3]:
			println(i)
		case i := <-chs[4]:
			println(i)
		case i := <-chs[5]:
			println(i)
		case i := <-chs[6]:
			println(i)
		case i := <-chs[7]:
			println(i)
		case i := <-chs[8]:
			println(i)
		case i := <-chs[9]:
			println(i)
		}
		time.Sleep(time.Millisecond * 500)
	}

}
