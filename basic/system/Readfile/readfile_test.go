package main

import (
	"net/http"
	"testing"
)

func BenchmarkSend(b *testing.B) {
	t := new(testing.T)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		send(t)
	}
}

func send(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:8080/readfile", nil)
	rsp, _ := http.DefaultClient.Do(req)
	if rsp != nil && rsp.Body != nil {
		rsp.Body.Close()
	}
}