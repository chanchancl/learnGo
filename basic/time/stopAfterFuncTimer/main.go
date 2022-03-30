package main

import (
	"fmt"
	"time"
)

func main() {

	time.AfterFunc(time.Second, func() {
		fmt.Println("Time!!!!!!!!!")
	})

	// fmt.Println(timer.Stop())

	select {}
}
