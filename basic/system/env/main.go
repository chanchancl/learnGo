package main

import (
	"fmt"
	"os"
)

func main() {
	env := os.Getenv("test")
	fmt.Println(env)
	if env == "" {
		fmt.Printf("Empty env == \"\"\n")
	}
}
