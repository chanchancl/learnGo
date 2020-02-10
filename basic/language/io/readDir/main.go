package main

import (
	"fmt"
	"io/ioutil"
	"os"
	pathpkg "path"
	"strings"
)

func main() {
	read(os.Args[1], 0)
}

func read(path string, indent int) {
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Read %s failed, %s\n", path, err.Error())
		return
	}
	for _, info := range infos {
		if info.IsDir() {
			fmt.Printf("%s%s\\\n", strings.Repeat(" ", indent), info.Name())
			read(pathpkg.Join(path, info.Name()), indent+2)
		} else {
			fmt.Printf("%s%s\n", strings.Repeat(" ", indent), info.Name())
		}
	}
}
