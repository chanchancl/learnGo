package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		fmt.Printf("\t%s\n", input.Text())
	}

}
