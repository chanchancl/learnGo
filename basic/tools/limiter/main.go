package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	limiter := rate.NewLimiter(0.5, 3)

	for i := 0; i < 10; i++ {
		limiter.Wait(context.Background())
		fmt.Println(time.Now())
	}

}
