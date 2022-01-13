package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("File 1 is %v\n", FileExist("1"))
	fmt.Printf("File 2 is %v\n", FileExist("2"))
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}
