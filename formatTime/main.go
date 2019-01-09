package main

import (
	"fmt"
	"time"
)

type someStruct struct {
	data      int
	startTime time.Time
}

func NewStruct() *someStruct {
	return &someStruct{startTime: time.Now()}
}

func main() {
	st := NewStruct()
	fmt.Printf("%#v\n", *st)
	//format
	// month day hour min sec year           6  1  2  3  4  5
	fmt.Printf("%s", st.startTime.Format("2006-01-02 15:04:05"))
}
