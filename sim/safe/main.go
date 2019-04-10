package main

import (
	"strconv"
	"time"
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
			time.Sleep(time.Second)
			//fmt.Println("Start!!!!!")
			//for _, k := range mp {
			//	fmt.Printf("    %v\n", k)
			//}
		}
	}()
	entity.value = make(map[string]interface{})
	for i := 0; i < 5; i++ {
		go func() {
			for {

				for i := 0; i < 10; i++ {
					entity.value[strconv.Itoa(i)] = i
				}
				//handler <- entity.value

				//entity.value = nil
			}
		}()
	}

	select {}
}
