package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// Root prevent accessing files outside the root directory such as symbolic links
	root, _ := os.OpenRoot("a")
	ReadFileOutofRoot(root.FS())

	// DirFS don't prevet accessing files outside the directory
	dirfs := os.DirFS("a")
	ReadFileOutofRoot(dirfs)
}

func ReadFileOutofRoot(ifs fs.FS) {
	data, err := fs.ReadFile(ifs, "main.go")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(data))
}
