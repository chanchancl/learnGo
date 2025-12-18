package main

import (
	"fmt"
	"sync"
	"time"
)

func rangeInt() {
	// 1.支持直接循环整数
	// for i := 0; i < 5; i++ {}
	for i := range 5 {
		fmt.Println(i, &i)
	}
}

func indepentLoopVariable() {
	// 2.在循环中，循环变量变为每个迭代独立创建，而不再是同一个地址
	iarray := []string{"a", "b", "c"}

	// Print c c c in 1.21
	// Print a b c after 1.21

	wg := sync.WaitGroup{}
	next := []chan struct{}{}
	for range len(iarray) {
		next = append(next, make(chan struct{}))
	}
	for i, k := range iarray {
		wg.Add(1)

		// 闭包捕获 i, k
		// 过去 i, k 的地址在循环中是相等，间接所有闭包捕获的值也都相同
		go func() {
			<-next[i]
			fmt.Println(i, k, &i, &k)
			wg.Done()
		}()

		// 新语法
		// wg.Go(func() {
		// 	fmt.Println(i, k, &i, &k)
		// })
	}

	for _, n := range next {
		n <- struct{}{}
		time.Sleep(1 * time.Millisecond)
	}

	wg.Wait()
}

func main() {
	rangeInt()

	indepentLoopVariable()
}
