package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	base := context.Background()

	cancleAble, cancle := context.WithCancel(base)

	CancleAction(cancleAble)

	cancle()

	time.Sleep(time.Second)
}

func CancleAction(ctx context.Context) {
	context.AfterFunc(ctx, func() {
		// cancled or timeout
		fmt.Println("cancleAble is cancled")
	})
}
