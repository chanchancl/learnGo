package main

import (
	"fmt"

	"github.com/ahmetb/go-linq/v3"
	"github.com/kr/pretty"
)

type DropInfo struct {
	DropID      int    `bun:",pk,autoincrement" json:"id"`
	Server      string `json:"server"`
	StageID     int    `json:"stageId"`
	DropType    string `json:"dropType"`
	RangeID     int    `json:"rangeId"`
	Accumulable bool   `json:"accumulable"`
}

var drops = []*DropInfo{
	{
		DropID:   1,
		Server:   "CN",
		DropType: "Normal",
	},
	{
		DropID:   2,
		Server:   "CN",
		DropType: "Normal",
	},
	{
		DropID:   3,
		Server:   "CN",
		DropType: "External",
	},
	{
		DropID:   4,
		Server:   "CN",
		DropType: "Normal",
	},
	{
		DropID:   5,
		Server:   "CN",
		DropType: "Normal",
	},
	{
		DropID:   6,
		Server:   "CN",
		DropType: "External",
	},
	{
		DropID:   7,
		Server:   "CN",
		DropType: "Normal",
	},
	{
		DropID:   8,
		Server:   "GLOBAL",
		DropType: "Normal",
	},
	{
		DropID:   9,
		Server:   "GLOBAL",
		DropType: "Normal",
	},
	{
		DropID:   10,
		Server:   "GLOBAL",
		DropType: "Normal",
	},
}

func main() {
	query := linq.From(drops)

	s := query.WhereT(func(i *DropInfo) bool { return i.Server == "CN" }).Count()
	fmt.Println(s)

	groupResult := []linq.Group{}
	query.GroupByT(
		func(i *DropInfo) string { return i.Server }, // key
		func(i *DropInfo) *DropInfo { return i },     // value, could be anything of i, even i self
	).ToSlice(&groupResult)
	pretty.Println(groupResult)
}
