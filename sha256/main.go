package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := md5.Sum([]byte("123456"))
	c2 := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n%x\n", c1, c2)
}
