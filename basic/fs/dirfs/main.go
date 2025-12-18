package main

import (
	"io/fs"
	"os"
	"strings"
)

func main() {
	// DirFS don't prevet accessing files outside the directory
	sfs := os.DirFS("..")

	if sfs == nil {
		panic("os.DirFS returned nil")
	}

	tab := 0
	fs.WalkDir(sfs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == "." {
			return nil
		}
		tab = strings.Count(path, "/")

		if d.IsDir() {
			println(strings.Repeat("  ", tab)+"Dir:", path)
		} else {
			println(strings.Repeat("  ", tab)+"File:", path)
		}
		return nil
	})

}
