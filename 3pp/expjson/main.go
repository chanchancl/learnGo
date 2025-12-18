package main

import (
	jsonv1 "encoding/json"
	"fmt"
	"time"

	jsonv2 "github.com/go-json-experiment/json"
	jsonv1on2 "github.com/go-json-experiment/json/v1"
)

type NormalStruct struct {
	A int `json:"a,omitempty"`
	B int `json:"b"`
}

func Marshal(t any) {
	v1, _ := jsonv1.Marshal(t)
	v2, _ := jsonv2.Marshal(t)
	v3, _ := jsonv1on2.Marshal(t)

	fmt.Println("v1   :", string(v1))
	fmt.Println("v2   :", string(v2))
	fmt.Println("v1on2:", string(v3))
	fmt.Println("")
}

type StructWithOmitzero struct {
	A int `json:"a,omitzero"`
	T struct {
		A int `json:"a"`
		B int `json:"b"`
	} `json:"t,omitzero"`
	MP map[string]int `json:"mp,omitzero"`
}

type StructWithTime struct {
	TimeOmitEmpty time.Time `json:"time_omitempty,omitempty"`
	TimeOmitZero  time.Time `json:"time_omitzero,omitzero"`

	DurationOmitEmpty time.Duration `json:"duration_omitempty,omitempty"`
	DurationOmitZero  time.Duration `json:"duration_omitzero,omitzero"`
}

func main() {

	timeNow := time.Now()
	Marshal(timeNow)

	duration := time.Second * 20
	Marshal(duration)

	t := NormalStruct{}
	Marshal(t)

	t2 := StructWithOmitzero{}
	Marshal(t2)

	t3 := StructWithTime{}
	Marshal(t3)

	// use omitzero only after go v1.24
}
