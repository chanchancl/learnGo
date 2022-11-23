package main

import (
	"fmt"
	"time"
)

func main() {
	last := time.Now()

	time.Sleep(time.Second)

	diff := time.Now().Sub(last)

	timeInSecond := diff.Seconds()

	fmt.Println(timeInSecond, int(timeInSecond))
}
