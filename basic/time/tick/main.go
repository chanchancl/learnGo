package main

import (
	"fmt"
	"time"
)

func main() {
	// ticker := time.NewTicker(time.Second)
	// for {
	// 	select {
	// 	case <-ticker.C:
	// 		fmt.Println("233")
	// 	}
	// }
	C := time.Tick(time.Second)
	for {
		select {
		case <-C:
			fmt.Println("233")
		}
	}
}
