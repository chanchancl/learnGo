package main

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type tjson struct {
	A int
	B int
	C string
}

func main() {
	t := tjson{1, 5, "testestest"}
	bt, _ := json.Marshal(t)
	logrus.WithField("field1", 2).Infof("test for struct print %s", string(bt))
}
