package main

import (
	"encoding/json"
	"fmt"
)

type PlmnIdType struct {
	Key int `json:"key"`
	K   int `json:"k"`
}

type RoamingPlmnContainer struct {
	VisitedPlmn []PlmnIdType `json:"visited-plmn"`
	Allowed     bool         `json:"allowed"`
	A           int          `json:"numberA,omitempty"`
}

func main() {
	str := `{"visited-plmn":[{"key":1, "k":2}], "allowed":true, "numberA": 0}`
	//str := "{}"

	o := RoamingPlmnContainer{}
	err := json.Unmarshal([]byte(str), &o)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", o)
}
