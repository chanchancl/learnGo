package main

import (
	"fmt"
	"strconv"
)

type Entity struct {
	value map[string]interface{}
}

var entity Entity

var (
	handler = make(chan map[string]interface{})
)

func main() {

	go func() {
		for {
			mp := <-handler
			fmt.Println("Start!!!!!")
			for _, k := range mp {
				fmt.Printf("    %v\n", k)
			}
		}
	}()

	for i := 0; i < 1; i++ {
		go func() {
			for {
				entity.value = make(map[string]interface{})
				for i := 0; i < 1000; i++ {
					entity.value[strconv.Itoa(i)] = i
				}
				handler <- entity.value

				entity.value = nil
			}
		}()
	}

	select {}
}
