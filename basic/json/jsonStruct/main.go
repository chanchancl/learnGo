package main

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type Ta struct {
	Text string `json:"text"`
}

type MyJsonName struct {
	A *Ta    `json:"a,omitempty"`
	B string `json:"b,omitempty"`
	C string `json:"c"`
	D uint32 `json:"d,omitempty"`
}

func main() {
	a := MyJsonName{
		A: &Ta{Text: "123"},
	}
	a.D = 0
	b, _ := json.Marshal(a)
	log.Info().Bytes("b", b).Msg("")

	a = MyJsonName{
		B: "2",
	}
	b, _ = json.Marshal(a)

	log.Info().Interface("a", a).Msg("")
	log.Info().Bytes("b", b).Msg("")
}
