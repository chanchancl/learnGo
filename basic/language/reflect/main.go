package main

import (
	"fmt"
	"reflect"
)

type innerStruct struct {
	a int
	b float64
}

type TestStruct struct {
	innerStruct
	namedInner innerStruct

	ch chan int
	mp map[int]int
}

func ref(in interface{}) {
	if reflect.TypeOf(in).Kind() != reflect.Ptr {
		fmt.Println("should recive a ptr struct")
		return
	}

	t := reflect.TypeOf(in).Elem()
	v := reflect.ValueOf(in).Elem()

	for idx := 0; idx < v.NumField(); idx++ {
		field := t.Field(idx)
		if field.Anonymous {
			fmt.Printf("anoymous : %v\n", field)
		} else {
			fmt.Printf("%v\n", field)
		}

		vfield := v.Field(idx)
		fmt.Println(vfield)
	}
}

func main() {
	s := &TestStruct{
		innerStruct: innerStruct{a: 10, b: 10.0},
		namedInner:  innerStruct{a: 20, b: 0.0},
		ch:          make(chan int),
		mp:          make(map[int]int),
	}
	ref(s)
}
