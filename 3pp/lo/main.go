package main

import (
	"fmt"

	"github.com/samber/lo"
	// lop "github.com/samber/lo/parallel"
)

func main() {
	names := lo.Uniq([]string{"a", "b", "a"})
	fmt.Println(names)

}
