package main

import (
	"fmt"
	"reflect"
)

func main() {
	abc := "abc"
	def := "abc"

	{
		str1 := &abc
		str2 := &def
		fmt.Println(reflect.DeepEqual(str1, str2))
	}

	{
		var str1 *string
		str2 := &abc
		fmt.Println(reflect.DeepEqual(str1, str2))
	}

	{
		var str1 *string
		var str2 *string
		fmt.Println(reflect.DeepEqual(str1, str2))
	}

}
