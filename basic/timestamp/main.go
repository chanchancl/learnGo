package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Second())
	fmt.Println(time.Now().Unix())
	fmt.Println(uint64(time.Duration(time.Now().UnixNano()) / time.Microsecond))
}
