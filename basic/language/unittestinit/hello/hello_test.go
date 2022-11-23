package hello

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("Init in test！")
}

func TestA(t *testing.T) {
	fmt.Println("Test in A")
}
