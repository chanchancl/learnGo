package main

import (
	"sync"
	"testing"
)

type Small struct {
	a [10]int
}

type Medium struct {
	a [1000]int
}

type Large struct {
	a [100000]int
}

var smallPool = sync.Pool{
	New: func() interface{} { return new(Small) },
}

var mediumPool = sync.Pool{
	New: func() interface{} { return new(Medium) },
}

var largePool = sync.Pool{
	New: func() interface{} { return new(Large) },
}


var globalSink *Small
var globalSinkMedium *Medium
var globalSinkLarge *Large

func Benchmark_Small(b *testing.B) {
	b.Run("WithoutPool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := &Small{}
			s.a[0] = 42
			globalSink = s
		}
	})
	b.Run("WithPool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := smallPool.Get().(*Small)
			s.a[0] = 42
			smallPool.Put(s)
			globalSink = s
		}
	})
}

func Benchmark_Medium(b *testing.B) {
	b.Run("WithoutPool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := &Medium{}
			m.a[0] = 42
			globalSinkMedium = m
		}
	})
	b.Run("WithPool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := mediumPool.Get().(*Medium)
			m.a[0] = 42
			mediumPool.Put(m)
			globalSinkMedium = m
		}
	})
}


func Benchmark_Large(b *testing.B) {
	b.Run("WithoutPool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			l := &Large{}
			l.a[0] = 42
			globalSinkLarge = l
		}
	})
	b.Run("WithPool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			l := largePool.Get().(*Large)
			l.a[0] = 42
			largePool.Put(l)
			globalSinkLarge = l
		}
	})
}
