package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	strs := []string{
		"b",
		"sant",
		"ptr",
		"int",
		"apple",
	}

	sort.Sort(StringSlice(strs))

	for i := range strs {
		fmt.Println(strs[i])
	}
}
