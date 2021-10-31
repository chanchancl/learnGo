package main

import (
	"fmt"
	"sort"
)

type Item struct {
	Score       int
	SameGroupID bool
}

type SortBySameGroupIDAndScore []*Item

func (a SortBySameGroupIDAndScore) Len() int      { return len(a) }
func (a SortBySameGroupIDAndScore) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBySameGroupIDAndScore) Less(i, j int) bool {
	if a[i].SameGroupID && a[j].SameGroupID {
		return a[i].Score < a[j].Score
	}
	if a[i].SameGroupID {
		return true
	}
	if a[j].SameGroupID {
		return false
	}
	return a[i].Score < a[j].Score
}

func main() {
	items := []*Item{
		{100, false},
		{50, false},
		{100, true},
		{0, false},
		{101, true},
		{99, true},
	}

	sort.Sort(SortBySameGroupIDAndScore(items))
	for _, item := range items {
		fmt.Println(item)
	}
}
