package main

import (
	"sync"
	"testing"
	"time"
)

const (
	cost = 10 * time.Microsecond
)

type RW interface {
	Write()
	Read()
}

type Lock struct {
	count int
	mu    sync.Mutex
}

func (l *Lock) Read() {
	l.mu.Lock()
	time.Sleep(cost)
	_ = l.count
	l.mu.Unlock()
}

func (l *Lock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost)
	l.mu.Unlock()
}

type RWLock struct {
	count int
	mu    sync.RWMutex
}

func (r *RWLock) Read() {
	r.mu.RLock()
	time.Sleep(cost)
	_ = r.count
	r.mu.RUnlock()
}

func (r *RWLock) Write() {
	r.mu.RLock()
	r.count++
	time.Sleep(cost)
	r.mu.RUnlock()
}

func benchmark(b *testing.B, rw RW, read, write int) {

	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for k := 0; k < read*100; k++ {
			wg.Add(1)
			go func() {
				rw.Read()
				wg.Done()
			}()
		}

		for m := 0; m < write*100; m++ {
			wg.Add(1)
			go func() {
				rw.Write()
				wg.Done()
			}()
		}
		wg.Wait()
	}

}

func BenchmarkReadMore(b *testing.B) {
	benchmark(b, &Lock{}, 9, 1)
}

func BenchmarkReadMoreRW(b *testing.B) {
	benchmark(b, &RWLock{}, 9, 1)
}

func BenchmarkWriteMore(b *testing.B) {
	benchmark(b, &Lock{}, 1, 9)
}

func BenchmarkWriteMoreRW(b *testing.B) {
	benchmark(b, &RWLock{}, 1, 9)
}

func BenchmarkReadEqual(b *testing.B) {
	benchmark(b, &Lock{}, 5, 5)
}

func BenchmarkReadEqualRW(b *testing.B) {
	benchmark(b, &RWLock{}, 5, 5)
}
