package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"testing"
)

const (
	Size = sha256.Size
)

func Benchmark_CompareByteArray(b *testing.B) {
	lbts := [Size]byte{}
	rbts := [Size]byte{}

	rand.Read(lbts[:])
	copy(rbts[:], lbts[:])
	b.Log(lbts, rbts)

	b.Run("CompareWithEqualOperator", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = lbts == rbts
		}
	})

	b.Run("CompareWithBytesCompare", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = bytes.Compare(lbts[:], rbts[:])
		}
	})
}
