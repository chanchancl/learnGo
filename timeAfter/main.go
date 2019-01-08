package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().After(time.Time{}))

	// 看到了比较奇怪的代码，所以验证一下
	fmt.Println(time.Now().Sub(time.Time{}).String())

	// 结论是，代码真的很奇怪，没有任何作用
}
