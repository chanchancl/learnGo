package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	root, _ := os.OpenRoot("a")
	ReadFileOutofRoot(root.FS())
}

func ReadFileOutofRoot(ifs fs.FS) {
	data, err := fs.ReadFile(ifs, "main.go")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(data))
}
