package main

// go test
// go test -bench=.

import (
	"fmt"
	"strconv"
	"testing"
)

func TestEqual(t *testing.T) {
	for i := 0; i < 100000; i++ {
		if fmt.Sprintf("%d", i) != strconv.Itoa(i) {
			fmt.Printf("Not Equal")
			t.FailNow()
		}
	}
}

func BenchmarkFmtSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", i)
	}
}

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Itoa(i)
	}
}
