package main

import (
	"fmt"
	lib "learnGo/lib/mylib"
)

func main() {
	fmt.Println(lib.PI)
	fmt.Println(lib.ExportFunction(9))
	//fmt.Println(lib.unexportFunction(10))
}
