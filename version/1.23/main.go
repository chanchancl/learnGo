package main

import (
	"fmt"
	"iter"
	"maps"
	"slices"
)

func slicesIterator() {
	// range 支持迭代器模式 (iterator)
	ia := []int{42, 15, 21, 10, 4, 99}

	for i := range slices.Values(ia) {
		fmt.Println(i)
	}

	for i, v := range slices.All(ia) {
		fmt.Println(i, v)
	}
}

func mapIterator() {
	imap := map[string]int{
		"a": 2,
		"b": 3,
		"c": 4,
	}

	for key := range maps.Keys(imap) {
		fmt.Println(key)
	}

	maps.Keys(imap)(func(v string) bool {
		fmt.Println(v)
		return true
	})
}

func myIterator() iter.Seq[int] {
	// 遍历这个迭代器，你会依次获得 0 - 4
	return func(yield func(int) bool) {
		for i := range 5 {
			if !yield(i) {
				return
			}
		}
	}
}

func iterator() {
	for it := range myIterator() {
		fmt.Println(it)
	}
}

func main() {
	slicesIterator()
	mapIterator()
	iterator()
}
