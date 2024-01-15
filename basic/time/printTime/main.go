package main

import (
	"fmt"
	"time"
)

func main() {
	timeDuration := 10 * time.Second

	fmt.Println(timeDuration / time.Second)
	fmt.Println(timeDuration.String())
	fmt.Println(timeDuration.Seconds())
}
