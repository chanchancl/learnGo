package main

import (
	"fmt"
	"time"
)

func f(t time.Time) {
	fmt.Println(t)
}

func main() {
	fmt.Println(time.Now())
	defer f(time.Now())
	time.Sleep(2 * time.Second)

}
