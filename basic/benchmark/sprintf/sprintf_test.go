package sprintf_test

import (
	"fmt"
	"testing"
)

func Benchmark_Sprintf(b *testing.B) {
	var result string
	for i := 0; i < b.N; i++ {
		result = fmt.Sprintf("%v://%v", "http", "test.com")
	}
	b.Log(result)
}

func Benchmark_StringConnect(b *testing.B) {
	var result string
	a := "http"
	c := "test.com"
	for i := 0; i < b.N; i++ {
		result = a + "://" + c
	}
	b.Log(result)
}
