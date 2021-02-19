package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	startErrorTime := time.Duration(10) * time.Second
	nextErrorTime := time.Duration(5) * time.Second

	errorLen := 3

	for {
		timePassed := time.Now().Sub(startTime)
		EndTime := startErrorTime + time.Duration(errorLen)*nextErrorTime

		index := int((timePassed - startErrorTime) / nextErrorTime)

		fmt.Printf("timePassed %v, startErrorTime %v, EndTime %v, index %v,\n", timePassed, startErrorTime, EndTime, index)
		fmt.Printf("    NextErrorTime == 0 %v\n", nextErrorTime == 0)
		fmt.Printf("    %v\n", timePassed < startErrorTime)
		fmt.Printf("    %v\n", timePassed > EndTime)
		time.Sleep(time.Second)
	}
}
