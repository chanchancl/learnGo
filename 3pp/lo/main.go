package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/samber/lo"
	// lop "github.com/samber/lo/parallel"
)

func Intersect() {
	fmt.Println(strings.Repeat("*", 80))
	fmt.Println("Intersect")
	names := lo.Uniq([]string{"a", "b", "a"})
	fmt.Println(names)

	fmt.Println(lo.Contains([]int{1, 2, 3}, 5))

	fmt.Println(lo.ContainsBy([]int{1, 2, 3}, func(x int) bool {
		return x%2 == 0
	}))

	fmt.Println(lo.Every([]int{1, 2, 3}, []int{1, 2, 5}))

	fmt.Println(lo.Some([]int{1, 2, 3}, []int{5, 6, 7, 1}))

	fmt.Println(lo.None([]int{1, 2, 3}, []int{7, 8, 9}))

	fmt.Println(lo.Intersect([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 6, 7}))

	fmt.Println(lo.Difference([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 6, 7}))

	fmt.Println(lo.Union([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 6, 7}))

	fmt.Println(lo.Without([]int{1, 2, 3, 4, 5}, 3, 4, 5, 6, 7))

	fmt.Println(lo.WithoutEmpty([]int{0, 1, 2, 3, 4, 0, 5}))
}

func Channel() {
	fmt.Println(strings.Repeat("*", 80))
	fmt.Println("Channel")

	root := make(chan int)

	childs := lo.ChannelDispatcher(root, 5, 0, lo.DispatchingStrategyRoundRobin)

	go func() {
		// Dispatch 12 data into 5 different channel
		for i := range 12 {
			root <- i
		}
	}()

	wg := sync.WaitGroup{}
	for i, ch := range childs {
		wg.Add(1)
		go func() {
			// after go 1.22, no need to copy i and ch into lambda
			defer wg.Done()
			for {
				v, ok := <-ch
				if !ok {
					fmt.Printf("Channel %v exit\n", i)
					return
				}
				fmt.Printf("channel %v received %v\n", i, v)
			}
		}()
	}

	time.AfterFunc(time.Millisecond*200, func() {
		// must close root here, other wise, all goroutines are asleep
		close(root)
	})
	wg.Wait()

	fmt.Println(strings.Repeat("*", 40))

	root = make(chan int, 20)

	go func() {
		for i := range 10 {
			root <- i
		}
		close(root)
	}()

	buf, length, time, ok := lo.Buffer(root, 3)
	fmt.Println(buf, length, time, ok)
	buf, length, time, ok = lo.Buffer(root, 10)
	fmt.Println(buf, length, time, ok)

	fmt.Println(strings.Repeat("*", 40))
	capa := 10
	chs := make([]chan int, 0, capa)
	for range capa {
		chs = append(chs, make(chan int))
	}

	rdchs := make([]<-chan int, 0, capa)
	for _, ch := range chs {
		rdchs = append(rdchs, ch)
	}

	out := lo.FanIn(0, rdchs...)

	wg = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := range out {
			fmt.Println(i)
		}
		wg.Done()
	}()

	for i, ch := range chs {
		ch <- 10 - i
		close(ch)
	}

	wg.Wait()
}

func Map() {
	fmt.Println(strings.Repeat("*", 80))
	fmt.Println("Map")

	mp := make(map[string]int)
	mp["a"] = 1
	mp["b"] = 2
	mp["c"] = 3

	fmt.Println(lo.Keys(mp))
	fmt.Println(lo.Values(mp))

	fmt.Println(lo.ValueOr(mp, "d", 4))

	fmt.Println(lo.ToPairs(mp))

	fmt.Println(lo.ToPairs(lo.Invert(mp)))

	fmt.Println(lo.MapKeys(mp, func(_ int, k string) string {
		return k + k + k
	}))

	fmt.Println(lo.MapValues(mp, func(v int, _ string) int {
		return v * 2
	}))
}

func Math() {
	fmt.Println(strings.Repeat("*", 80))
	fmt.Println("Math")

	// python : range(5)
	fmt.Println(lo.Range(5))
	// python : range(10, 15)
	fmt.Println(lo.RangeFrom(10, 5))

	fmt.Println(lo.Sum(lo.Range(101)))
}

func main() {
	Intersect()
	Channel()
	Map()
	Math()

}
