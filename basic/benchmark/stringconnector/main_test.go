package main

import (
	"fmt"
	"strconv"
	"testing"
)

func Benchmark_Sprintf(b *testing.B) {
	result := ""
	for i := 0; i < b.N; i++ {
		result = fmt.Sprintf("abcdef%v", i)
	}

	b.Log(result)
}

func Benchmark_DirectConnect(b *testing.B) {
	result := ""

	for i := 0; i < b.N; i++ {
		result = "abcdef" + strconv.Itoa(i)
	}
	b.Log(result)
}
